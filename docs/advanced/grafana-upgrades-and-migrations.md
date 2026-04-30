# Grafana upgrades and migrations (CloudNativePg)

Upgrading Grafana across major versions with CloudNativePg requires coordinating application upgrades with database backup/restore strategy. This guide assumes PostgreSQL is deployed via **CloudNativePg operator** on Kubernetes.

## Version progression: 12.3.x → 12.4.x → 13.0.x

This guide covers upgrading from Grafana 12.3.x through 12.4.x to 13.0.1, the current stable release, while maintaining database integrity through CloudNativePg.

### Pre-upgrade checklist with CloudNativePg

**1. Verify CloudNativePg cluster health**

```bash
# Check cluster status
kubectl get cluster grafana-db -n postgres -o wide

# Verify replication is healthy
kubectl get pod -n postgres -l cnpg.io/cluster=grafana-db

# Check database connections
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d grafana -c "SELECT datname, count(*) FROM pg_stat_activity GROUP BY datname;"
```

**2. Create pre-upgrade backup via CloudNativePg**

```bash
# Verify scheduled backup is running
kubectl get scheduledbackup -n postgres

# Trigger immediate backup
kubectl create -f - <<EOF
apiVersion: postgresql.cnpg.io/v1
kind: Backup
metadata:
  name: grafana-pre-v13-upgrade
  namespace: postgres
spec:
  cluster:
    name: grafana-db
EOF

# Wait for backup to complete
kubectl wait --for=condition=Ready backup grafana-pre-v13-upgrade \
  -n postgres --timeout=600s

# Verify backup
kubectl get backup grafana-pre-v13-upgrade -n postgres -o wide
```

**3. Snapshot Helm values**

```bash
# Save current Helm deployment state
helm get values grafana -n grafana > grafana-12.3-values.yaml
helm get values grafana -n grafana --all > grafana-12.3-all-values.yaml
```

**4. Document datasource encryption keys**

If using encrypted datasources, ensure you have:

```bash
# CloudNativePg secret used by Grafana
kubectl get secret grafana-db-credentials -n postgres -o yaml > db-credentials-backup.yaml

# Any secrets mounted in Grafana pod
kubectl get secret -n grafana -l app=grafana -o yaml > grafana-secrets-backup.yaml
```

## Breaking changes: 12.3.x → 12.4.x

**Plugin environment variables**
- Plugins no longer receive host environment variables by default
- If plugins need specific env vars, explicitly pass them in plugin configuration
- Affects custom plugins and third-party datasource plugins

**Database schema stability**
- CloudNativePg handles all migrations automatically
- No manual SQL scripts needed
- Migrations are tested on CNP side before applying

## Notable features in 12.4.x

- **Alerting enrichment** — enterprise feature for context-aware alerting
- **RBAC enhancements** — permissions for reading/writing enrichments
- **Dashboard UI improvements** — faster rendering with V2 dashboard system
- **API stability** — security scope checks on dashboard endpoints

## Major changes: 12.4.x → 13.0.1

### What's new in 13.0.1

- **Dashboard timezone preservation** — user preferences now retained when converting V1 → V2 dashboards
- **Provisioning improvements** — dashboard validation errors included in pull request comments (GitOps workflows)
- **Unified storage** — migrations skip if already applied (safe with CNP replication)
- **API client enhancements** — lazy hooks, automatic PATCH headers

### Database implications with CloudNativePg

- All schema migrations applied automatically by Grafana startup
- CloudNativePg replication continues during migrations (no downtime)
- Point-in-time recovery available if migration rollback needed

## Helm upgrade strategy (using grafana-community charts + CloudNativePg)

### Step 1: Prepare values

```yaml
# values-12.3.yaml (current state)
---
grafana:
  image:
    tag: "12.3.0"
  adminPassword: "changeme"
  datasources:
    datasources.yaml:
      apiVersion: 1
      datasources:
      - name: PostgreSQL
        type: postgres
        database: grafana
        user: grafana
        password: ${POSTGRES_PASSWORD}  # From CloudNativePg secret
        host: grafana-db-rw.postgres.svc.cluster.local  # CNP R/W endpoint

# values-13.0.yaml (target)
---
grafana:
  image:
    tag: "13.0.1"
  adminPassword: "changeme"
  datasources:
    datasources.yaml:
      apiVersion: 1
      datasources:
      - name: PostgreSQL
        type: postgres
        database: grafana
        user: grafana
        password: ${POSTGRES_PASSWORD}  # Same CloudNativePg secret
        host: grafana-db-rw.postgres.svc.cluster.local  # Same R/W endpoint
```

### Step 2: Dry run the upgrade

```bash
# Helm dry-run with new version
helm upgrade --dry-run grafana grafana/grafana \
  --namespace grafana \
  -f values-13.0.yaml \
  --version 13.0.0 \
  2>&1 | tee dry-run.log

# Check for:
# - API deprecation warnings
# - Resource requirement changes
# - Permission model shifts
```

### Step 3: Ensure CloudNativePg is ready

```bash
# Verify cluster is in healthy state
kubectl get cluster grafana-db -n postgres -o jsonpath='{.status.phaseDetails}'

# Check no other upgrades in progress
kubectl get pod -n postgres -o wide | grep -i upgrade

# Ensure standby replicas are synchronized
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d postgres -c "SELECT pid, state, sync_state FROM pg_stat_replication;"
```

Expected output: `sync` state for all replicas, `streaming` state.

### Step 4: Execute upgrade with zero downtime

```bash
# Scale to multiple replicas (if not already)
# This ensures traffic continues during rolling restart
kubectl scale deployment grafana --replicas=3 -n grafana

# Upgrade Grafana
helm upgrade grafana grafana/grafana \
  --namespace grafana \
  -f values-13.0.yaml \
  --version 13.0.0 \
  --wait \
  --timeout 5m

# Monitor rollout
kubectl rollout status deployment/grafana -n grafana

# Scale back if desired
kubectl scale deployment grafana --replicas=1 -n grafana
```

### Step 5: Post-upgrade validation

```bash
# Check Grafana pod health
kubectl get pod -n grafana -l app=grafana

# Verify Grafana is connected to CloudNativePg
kubectl logs -f deployment/grafana -n grafana | grep -i "database"

# Check for migration logs
kubectl logs deployment/grafana -n grafana | grep -i "migration"

# API health check via port-forward
kubectl port-forward svc/grafana 3000:80 -n grafana &
sleep 2
curl -s http://localhost:3000/api/health | jq
killall kubectl

# Verify CloudNativePg cluster still healthy
kubectl get cluster grafana-db -n postgres -o jsonpath='{.status.phase}'
```

## CloudNativePg-specific procedures

### Database migration monitoring

```bash
# Watch database activity during Grafana upgrade
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d grafana -c \
  "SELECT pid, usename, application_name, query, state FROM pg_stat_activity WHERE state != 'idle';" \
  --watch=5

# Monitor replication lag
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d postgres -c \
  "SELECT slot_name, slot_type, restart_lsn FROM pg_replication_slots;" \
  --watch=5
```

### Point-in-time recovery (if needed)

If Grafana 13.0.1 introduces unexpected issues, recover from backup:

```bash
# List available backups
kubectl get backup -n postgres | grep grafana-db

# Restore from specific backup
kubectl create -f - <<EOF
apiVersion: postgresql.cnpg.io/v1
kind: Cluster
metadata:
  name: grafana-db-restored
  namespace: postgres
spec:
  instances: 2
  bootstrap:
    recovery:
      backup:
        name: grafana-pre-v13-upgrade
EOF

# Wait for recovery
kubectl wait --for=condition=Ready cluster grafana-db-restored \
  -n postgres --timeout=600s

# Update Grafana connection string to restored cluster
# (Use grafana-db-restored-rw.postgres.svc.cluster.local)
```

### Verifying data consistency post-upgrade

```bash
# Check CloudNativePg data integrity
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d grafana -c \
  "SELECT schemaname, tablename FROM pg_tables WHERE schemaname = 'public' ORDER BY tablename;" | head -20

# Verify dashboard count unchanged
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d grafana -c \
  "SELECT COUNT(*) FROM dashboard;"
```

## Troubleshooting with CloudNativePg backend

**Grafana can't connect to database after upgrade**

```bash
# Verify CloudNativePg credentials haven't changed
kubectl get secret grafana-db-credentials -n postgres -o yaml

# Check Grafana pod has correct secret mounted
kubectl describe pod deployment/grafana -n grafana | grep -A5 "Mounts:"

# Test connection manually from pod
kubectl exec -it deployment/grafana -n grafana -- \
  psql "postgresql://grafana:PASSWORD@grafana-db-rw.postgres.svc.cluster.local:5432/grafana" \
  -c "SELECT 1"
```

**CloudNativePg replica lag during upgrade**

```bash
# Check current WAL position
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d postgres -c "SELECT pg_current_wal_lsn();"

# Force WAL flush if needed
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d postgres -c "SELECT pg_wal_lsn_diff(pg_current_wal_lsn(), '0/0');"

# If lag is persistent, trigger manual checkpoint
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d postgres -c "CHECKPOINT;"
```

**Database connections stuck during migration**

```bash
# List active connections
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d grafana -c \
  "SELECT pid, usename, application_name, query FROM pg_stat_activity WHERE state = 'active';"

# Terminate blocking connections (use with caution)
kubectl exec -it grafana-db-1 -n postgres -- \
  psql -U postgres -d grafana -c \
  "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE application_name = 'old_grafana_version';"
```

## Rollback procedure with CloudNativePg

If upgrade fails critically:

```bash
# Option 1: Helm rollback (fastest)
helm rollback grafana 0 -n grafana

# Option 2: Restore from CloudNativePg backup
# (See "Point-in-time recovery" section above)

# Verify Grafana reconnects to original database
kubectl logs deployment/grafana -n grafana | grep -i "connected"

# Confirm CloudNativePg cluster unaffected
kubectl get cluster grafana-db -n postgres -o jsonpath='{.status.phase}'
```

## Production deployment checklist

- [ ] CloudNativePg cluster in healthy state (all replicas synced)
- [ ] Pre-upgrade backup created and verified
- [ ] Helm values migrated to grafana-community format
- [ ] Dry-run completed without warnings
- [ ] Helm values snapshots taken
- [ ] CloudNativePg credentials documented
- [ ] Datasources and plugins tested in staging
- [ ] Alerting rules exported for reference
- [ ] RBAC permissions documented and tested
- [ ] Multiple Grafana replicas for rolling update
- [ ] Rollback procedure tested (in staging cluster)
- [ ] Team notified of maintenance window
- [ ] Monitoring dashboards prepared to watch upgrade


## Resources

- [Official Grafana Helm charts](https://github.com/grafana-community/helm-charts)
- [CloudNativePg operator documentation](https://cloudnative-pg.io/)
- [Grafana v13 download](https://grafana.com/grafana/download/13.0.1)
- [What's new in Grafana v13](https://grafana.com/docs/grafana/latest/whatsnew/)
- [Grafana breaking changes](https://grafana.com/docs/grafana/latest/breaking-changes/)
