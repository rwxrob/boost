# Grafana upgrades and migrations

Upgrading Grafana across major versions requires careful planning, especially when managing dashboards, data sources, and Helm deployment configurations.

## Version progression: 12.3.x → 12.4.x → 13.0.x

This guide covers upgrading from Grafana 12.3.x through 12.4.x to 13.0.1, the current stable release.

### Key decision points before upgrading

**Helm or direct deployment?**

Official Grafana Helm charts are now maintained by the **grafana-community** org (not Bitnami). This is significant—if you're using Bitnami charts, you should migrate to the official community charts during this upgrade.

- **grafana-community/helm-charts**: Official, actively maintained
- **Bitnami charts**: Deprecated for new deployments (use community charts instead)

**Backup strategy**

Before upgrading:

```bash
# Dump your database if using external DB
mysqldump -u grafana -p grafana > grafana-backup.sql
# or
pg_dump -U grafana grafana > grafana-backup.sql

# Export dashboards and data sources via API
# or use Helm values snapshot
helm get values release-name > helm-values-backup.yaml
```

## Breaking changes: 12.3.x → 12.4.x

**Plugin environment variables**
- Plugins no longer receive host environment variables by default
- If plugins need specific env vars, explicitly pass them in plugin configuration
- Affects custom plugins and third-party datasource plugins

**Deprecations to plan for**
- Some legacy authentication methods flagged for removal in v13
- Service account tokens—migrate to new RBAC model if not already done

## Notable features in 12.4.x

- **Alerting enrichment** — enterprise feature for context-aware alerting
- **RBAC enhancements** — permissions for reading/writing enrichments
- **Dashboard UI improvements** — faster rendering with V2 dashboard system
- **API stability** — security scope checks on dashboard endpoints

## Major changes: 12.4.x → 13.0.1

### What's new in 13.0.1

- **Dashboard timezone preservation** — user preferences now retained when converting V1 → V2 dashboards
- **Provisioning improvements** — dashboard validation errors included in pull request comments (GitOps workflows)
- **Unified storage** — migrations skip if already applied (safer idempotent upgrades)
- **API client enhancements** — lazy hooks, automatic PATCH headers

### Plugin and deployment implications

- Host environment variable isolation may require plugin reconfigurations
- Dashboard rendering pipeline changes—legacy custom panels may need updates
- Database migrations are now safer and more robust

## Helm upgrade strategy (using grafana-community charts)

### Step 1: Prepare values

```yaml
# values-12.3.yaml (current)
---
grafana:
  image:
    tag: "12.3.0"
  adminPassword: "changeme"
  datasources: {}  # Reference your current datasources

# values-13.0.yaml (target)
---
grafana:
  image:
    tag: "13.0.1"
  adminPassword: "changeme"
  datasources: {}  # Same datasources
```

### Step 2: Dry run the upgrade

```bash
# Helm dry-run with new version
helm upgrade --dry-run grafana grafana/grafana \
  --namespace grafana \
  -f values-13.0.yaml \
  --version 13.0.0
```

Check for:
- API deprecation warnings
- Resource requirement changes
- Permission model shifts

### Step 3: Execute upgrade

```bash
# Upgrade in-place
helm upgrade grafana grafana/grafana \
  --namespace grafana \
  -f values-13.0.yaml \
  --version 13.0.0 \
  --wait

# Verify pod health
kubectl rollout status deployment/grafana -n grafana
```

### Step 4: Post-upgrade validation

```bash
# Port-forward and test
kubectl port-forward svc/grafana 3000:80 -n grafana

# Check logs for errors
kubectl logs -f deployment/grafana -n grafana

# API health check
curl -s http://localhost:3000/api/health | jq
```

## Database migration considerations

v13.0.1 introduces safer database migrations:

- **Unified storage migration** — checks if already applied, skips if complete
- **Idempotent operations** — safe to re-apply without corruption
- **Backwards-compatible** — you can roll back if needed (though not recommended long-term)

Monitor migration logs:

```bash
# Check if migrations ran
kubectl logs deployment/grafana -n grafana | grep -i migration

# Watch for completion
kubectl logs deployment/grafana -n grafana | tail -20
```

## Datastore-specific considerations

### PostgreSQL

- v13 works with PostgreSQL 12+
- No schema changes requiring manual intervention
- Connection pooling (PgBouncer) works as before

### MySQL/MariaDB

- v13 works with MySQL 5.7+, MariaDB 10.2+
- No breaking schema changes
- Same connection requirements as v12

## Dashboard and datasource migration

Existing dashboards auto-convert from V1 to V2 format:

- **Automatic on first load** — no action needed
- **Timezone preferences preserved** — user settings maintained
- **Legacy panels** — still work but consider updating to new panel types

Check for incompatible plugins:

```bash
# In Grafana UI: Configuration > Plugins
# Verify each plugin is compatible with v13
# Check plugin repo for 13.0 compatibility
```

## Rollback procedure

If upgrade fails:

```bash
# Helm rollback (if using Helm)
helm rollback grafana 0 -n grafana

# Or redeploy previous version
helm upgrade grafana grafana/grafana \
  --namespace grafana \
  -f values-12.4.yaml \
  --version 12.4.3

# Restore from database backup if needed
mysql grafana < grafana-backup.sql
```

## Troubleshooting common upgrade issues

**Datasources not appearing after upgrade**
- Check datasource permissions in new RBAC model
- Verify secrets are mounted correctly
- Restart datasource discovery: `kubectl rollout restart deployment/grafana`

**Plugins broken after upgrade**
- Plugins require recompilation for v13 if they use internal APIs
- Check plugin GitHub repo for v13 release
- Fallback: disable plugin if no v13 build available

**Dashboard rendering slow**
- V1→V2 conversion happens on first view (takes ~30 seconds for complex dashboards)
- Subsequent loads are faster
- Monitor dashboard load time in browser DevTools

**Authentication failures**
- Service account tokens may need regeneration
- RBAC permissions changed—verify user roles assigned
- Check auth provider compatibility (LDAP, SAML, OAuth)

## Production deployment checklist

- [ ] Database backup created and verified
- [ ] Helm values dry-run completed without warnings
- [ ] Datasources and plugins tested in staging
- [ ] Alerting rules exported for reference
- [ ] RBAC permissions documented and tested
- [ ] Rollback procedure tested
- [ ] Maintenance window scheduled (5–15 min typical)
- [ ] Post-upgrade health checks automated
- [ ] Team notified of maintenance window

## Resources

- [Official Grafana Helm charts](https://github.com/grafana-community/helm-charts)
- [Grafana v13 download](https://grafana.com/grafana/download/13.0.1)
- [What's new in Grafana v13](https://grafana.com/docs/grafana/latest/whatsnew/)
- [Grafana breaking changes](https://grafana.com/docs/grafana/latest/breaking-changes/)
