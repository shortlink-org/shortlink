# 31. Postgres tips [Cookbook]

Date: 2024-01-22

## Status

Accepted

## Cookbook

### Recipe 1: One PID to Lock Them All: Finding the Source of the Lock in Postgres

> [!TIP]
> Tutorial: https://www.crunchydata.com/blog/one-pid-to-lock-them-all-finding-the-source-of-the-lock-in-postgres

```sql
WITH sos AS (
	SELECT array_cat(array_agg(pid),
           array_agg((pg_blocking_pids(pid))[array_length(pg_blocking_pids(pid),1)])) pids
	FROM pg_locks
	WHERE NOT granted
)
SELECT a.pid, a.usename, a.datname, a.state,
	   a.wait_event_type || ': ' || a.wait_event AS wait_event,
       current_timestamp-a.state_change time_in_state,
       current_timestamp-a.xact_start time_in_xact,
       l.relation::regclass relname,
       l.locktype, l.mode, l.page, l.tuple,
       pg_blocking_pids(l.pid) blocking_pids,
       (pg_blocking_pids(l.pid))[array_length(pg_blocking_pids(l.pid),1)] last_session,
       coalesce((pg_blocking_pids(l.pid))[1]||'.'||coalesce(case when locktype='transactionid' then 1 else array_length(pg_blocking_pids(l.pid),1)+1 end,0),a.pid||'.0') lock_depth,
       a.query
FROM pg_stat_activity a
     JOIN sos s on (a.pid = any(s.pids))
     LEFT OUTER JOIN pg_locks l on (a.pid = l.pid and not l.granted)
ORDER BY lock_depth;
```
