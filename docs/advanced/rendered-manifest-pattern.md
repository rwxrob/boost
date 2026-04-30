# The rendered manifest pattern

Never run `helm install` without first seeing exactly what it will put into your cluster. This is a foundational Kubernetes operations habit, and it applies far beyond Helm.

## The problem with blind installs

Helm charts are templates. When you run `helm install`, Helm renders those templates with your values, generates a set of Kubernetes resource manifests, and applies them to your cluster. The problem is that most people never look at what was rendered. They trust the chart documentation, the CHANGELOG, and the README to tell them what the chart does.

That trust is misplaced. Chart documentation is frequently incomplete, sometimes wrong, and often out of date. A Helm chart upgrade might add a new ClusterRole with broad permissions, register a mutating webhook, create CRDs you did not ask for, inject a sidecar container into your pods, change resource limits, or modify network policies — all without a single mention in the release notes. The only way to know what a chart actually does is to look at what it produces.

## Always render first

The fix is straightforward. Before you install or upgrade anything, render the chart to plain Kubernetes YAML and read it:

```
helm template my-release some-chart/ -f values.yaml > rendered.yaml
```

That file contains every resource that Helm would create — every Deployment, Service, ConfigMap, RBAC rule, CRD, and webhook. Read it. Understand what is going into your cluster before it gets there.

## Commit everything to git

This pattern becomes useful when you treat the rendered output as source-controlled configuration. For every Helm-based application you manage, maintain a git repo (or a directory within a mono-repo) that contains:

- The vendored chart tarball (the exact version you are deploying)
- Your `values.yaml` with any customizations
- The rendered manifests (the output of `helm template`)
- A deploy script or Makefile that applies the manifests

Commit all of it. Tag each deployment. This gives you a versioned, auditable record of what is running in your cluster at any point in time.

## The diff is the real changelog

When a new version of a chart is released, the upgrade workflow looks like this:

1. Vendor the new chart tarball
2. Render it with your existing values: `helm template my-release new-chart/ -f values.yaml > rendered-new.yaml`
3. Diff the new rendered output against the previous committed version: `diff rendered.yaml rendered-new.yaml`
4. Read the diff. Understand every change. Then decide whether to apply it.

That diff is the real changelog. It shows you exactly what changed in the resources that will be applied to your cluster — not what the maintainer says changed, not what the CHANGELOG claims, but what actually changed in the rendered Kubernetes manifests.

This has caught real problems in production environments: charts that silently added cluster-wide RBAC permissions during a patch release, charts that registered validating webhooks without documenting them, charts that changed default resource limits in ways that would starve other workloads, charts that added new CRDs that conflicted with existing ones. None of these were documented. All of them were visible in the diff.

## The K8SAPP conventions

This pattern was formalized years ago as K8SAPP (Kubernetes Static Application Package) — a set of conventions for managing Kubernetes applications that treats rendered manifests as the source of truth and git as the management layer.

The core K8SAPP contract asks that every application repo follows a consistent structure:

- **Fetch** — acquire the chart or manifests from the upstream source
- **Review** — render, read the output, review RBAC, CRDs, policies, and dependencies
- **Configure** — adapt the values to match your environment
- **Install** — apply the rendered manifests for the first time
- **Upgrade** — render the new version, diff, review, then apply
- **Uninstall** — remove all traces from the cluster
- **Check** — verify the latest version, check for updates, validate dependencies

Each application gets its own repo (or directory) with a README documenting these procedures, the vendored chart, the values file, and the rendered output. This gives every team member a consistent interface for every application regardless of how complex the underlying chart is.

The key insight is that Helm discarded git — the standard version control system used by every other part of the software lifecycle — and replaced it with tarballs. K8SAPP puts everything Helm produces back into git where it can be forked, branched, diffed, blamed, and bisected like application code.

## How AI fits in

Diffing rendered manifests is well suited to AI-assisted review. An AI agent can compare two rendered manifest files and explain what changed in plain language: "This upgrade adds a new ClusterRole that grants list and watch permissions on all secrets cluster-wide. The previous version only had access to secrets in the application namespace."

That kind of summary takes a human several minutes of careful reading. An AI agent produces it in seconds. The rendered manifest pattern gives the agent the right input — two plain YAML files to compare — and the output is immediately actionable.

The workflow becomes: render, commit, upgrade, render again, hand the diff to an AI agent, read the summary, decide whether to proceed. The rendered manifests are the documentation. The AI reads them for you and flags changes worth reviewing.

## Beyond Helm

This pattern is not specific to Helm. Any system that generates configuration from templates benefits from the same approach: render the output, commit it, diff across versions. Kustomize overlays, Terraform plans, Ansible playbooks that generate config files, CI/CD pipelines that produce deployment manifests — all of them become easier to audit when you can see and version what they produce.

The principle is always the same: never trust the tool to tell you what it did. Look at what it produced. Commit it. Diff it next time. The rendered output is the truth. Everything else is commentary.

## The bottom line

Render first. Read what you are about to deploy. Commit the rendered output to git. When you upgrade, diff the new rendering against the old one. Let AI help you review the diff. This catches the problems that documentation does not tell you about.
