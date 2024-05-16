# grafana-loki

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | loki | 6.5.2 |
| https://grafana.github.io/helm-charts | promtail | 6.15.5 |

## Values

<table height="400px" >
	<thead>
		<th>Key</th>
		<th>Type</th>
		<th>Default</th>
		<th>Description</th>
	</thead>
	<tbody>
		<tr>
			<td id="loki--backend--replicas"><a href="./values.yaml#L138">loki.backend.replicas</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
0
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--backend--resources--limits--cpu"><a href="./values.yaml#L142">loki.backend.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--backend--resources--limits--memory"><a href="./values.yaml#L143">loki.backend.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--backend--resources--requests--cpu"><a href="./values.yaml#L145">loki.backend.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--backend--resources--requests--memory"><a href="./values.yaml#L146">loki.backend.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--deploymentMode"><a href="./values.yaml#L6">loki.deploymentMode</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"SingleBinary"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--enabled"><a href="./values.yaml#L2">loki.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--gateway--enabled"><a href="./values.yaml#L118">loki.gateway.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--gateway--resources--limits--cpu"><a href="./values.yaml#L122">loki.gateway.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--gateway--resources--limits--memory"><a href="./values.yaml#L123">loki.gateway.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--gateway--resources--requests--cpu"><a href="./values.yaml#L125">loki.gateway.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--gateway--resources--requests--memory"><a href="./values.yaml#L126">loki.gateway.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--auth_enabled"><a href="./values.yaml#L11">loki.loki.auth_enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--commonConfig--replication_factor"><a href="./values.yaml#L14">loki.loki.commonConfig.replication_factor</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--limits_config--volume_enabled"><a href="./values.yaml#L17">loki.loki.limits_config.volume_enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--revisionHistoryLimit"><a href="./values.yaml#L9">loki.loki.revisionHistoryLimit</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
3
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--schemaConfig--configs[0]--from"><a href="./values.yaml#L24">loki.loki.schemaConfig.configs[0].from</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"2024-04-01"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--schemaConfig--configs[0]--index--period"><a href="./values.yaml#L30">loki.loki.schemaConfig.configs[0].index.period</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"24h"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--schemaConfig--configs[0]--index--prefix"><a href="./values.yaml#L29">loki.loki.schemaConfig.configs[0].index.prefix</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"index_"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--schemaConfig--configs[0]--object_store"><a href="./values.yaml#L26">loki.loki.schemaConfig.configs[0].object_store</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"filesystem"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--schemaConfig--configs[0]--schema"><a href="./values.yaml#L27">loki.loki.schemaConfig.configs[0].schema</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"v13"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--schemaConfig--configs[0]--store"><a href="./values.yaml#L25">loki.loki.schemaConfig.configs[0].store</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"tsdb"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--storage--type"><a href="./values.yaml#L20">loki.loki.storage.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"filesystem"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--structuredConfig--memberlist--bind_addr"><a href="./values.yaml#L37">loki.loki.structuredConfig.memberlist.bind_addr</a></td>
			<td>
list
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
[]
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--loki--structuredConfig--query_range--align_queries_with_step"><a href="./values.yaml#L34">loki.loki.structuredConfig.query_range.align_queries_with_step</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--lokiCanary--enabled"><a href="./values.yaml#L115">loki.lokiCanary.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--dashboards--annotations--grafana_dashboard_folder"><a href="./values.yaml#L91">loki.monitoring.dashboards.annotations.grafana_dashboard_folder</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Loki"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--dashboards--enabled"><a href="./values.yaml#L89">loki.monitoring.dashboards.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--name"><a href="./values.yaml#L105">loki.monitoring.rules.additionalGroups[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"additional-loki-rules"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--rules[0]--expr"><a href="./values.yaml#L108">loki.monitoring.rules.additionalGroups[0].rules[0].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sum(rate(loki_request_duration_seconds_bucket[1m])) by (le, job)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--rules[0]--record"><a href="./values.yaml#L107">loki.monitoring.rules.additionalGroups[0].rules[0].record</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"job:loki_request_duration_seconds_bucket:sum_rate"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--rules[1]--expr"><a href="./values.yaml#L110">loki.monitoring.rules.additionalGroups[0].rules[1].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sum(rate(loki_request_duration_seconds_bucket[1m])) by (le, job, route)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--rules[1]--record"><a href="./values.yaml#L109">loki.monitoring.rules.additionalGroups[0].rules[1].record</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"job_route:loki_request_duration_seconds_bucket:sum_rate"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--rules[2]--expr"><a href="./values.yaml#L112">loki.monitoring.rules.additionalGroups[0].rules[2].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sum(rate(container_cpu_usage_seconds_total[1m])) by (node, namespace, pod, container)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--additionalGroups[0]--rules[2]--record"><a href="./values.yaml#L111">loki.monitoring.rules.additionalGroups[0].rules[2].record</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"node_namespace_pod_container:container_cpu_usage_seconds_total:sum_rate"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--rules--enabled"><a href="./values.yaml#L103">loki.monitoring.rules.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--selfMonitoring--grafanaAgent--installOperator"><a href="./values.yaml#L100">loki.monitoring.selfMonitoring.grafanaAgent.installOperator</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--serviceMonitor--enabled"><a href="./values.yaml#L94">loki.monitoring.serviceMonitor.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--monitoring--serviceMonitor--labels--release"><a href="./values.yaml#L96">loki.monitoring.serviceMonitor.labels.release</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus-operator"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--nameOverride"><a href="./values.yaml#L4">loki.nameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-loki"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--read--persistence--storageClass"><a href="./values.yaml#L60">loki.read.persistence.storageClass</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local-path"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--read--replicas"><a href="./values.yaml#L57">loki.read.replicas</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
0
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--read--resources--limits--cpu"><a href="./values.yaml#L64">loki.read.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--read--resources--limits--memory"><a href="./values.yaml#L65">loki.read.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--read--resources--requests--cpu"><a href="./values.yaml#L67">loki.read.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--read--resources--requests--memory"><a href="./values.yaml#L68">loki.read.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--sidecar--resources--limits--cpu"><a href="./values.yaml#L131">loki.sidecar.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--sidecar--resources--limits--memory"><a href="./values.yaml#L132">loki.sidecar.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--sidecar--resources--requests--cpu"><a href="./values.yaml#L134">loki.sidecar.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--sidecar--resources--requests--memory"><a href="./values.yaml#L135">loki.sidecar.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--extraArgs[0]"><a href="./values.yaml#L74">loki.singleBinary.extraArgs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"--pattern-ingester.enabled=true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--persistence--storageClass"><a href="./values.yaml#L77">loki.singleBinary.persistence.storageClass</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local-path"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--replicas"><a href="./values.yaml#L71">loki.singleBinary.replicas</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--resources--limits--cpu"><a href="./values.yaml#L81">loki.singleBinary.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--resources--limits--memory"><a href="./values.yaml#L82">loki.singleBinary.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--resources--requests--cpu"><a href="./values.yaml#L84">loki.singleBinary.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--singleBinary--resources--requests--memory"><a href="./values.yaml#L85">loki.singleBinary.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--test--enabled"><a href="./values.yaml#L40">loki.test.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
false
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--tracing--enabled"><a href="./values.yaml#L149">loki.tracing.enabled</a></td>
			<td>
bool
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
true
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--tracing--jaegerAgentHost"><a href="./values.yaml#L150">loki.tracing.jaegerAgentHost</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana:6831"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--write--persistence--storageClass"><a href="./values.yaml#L46">loki.write.persistence.storageClass</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local-path"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--write--replicas"><a href="./values.yaml#L43">loki.write.replicas</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
0
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--write--resources--limits--cpu"><a href="./values.yaml#L50">loki.write.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--write--resources--limits--memory"><a href="./values.yaml#L51">loki.write.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--write--resources--requests--cpu"><a href="./values.yaml#L53">loki.write.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="loki--write--resources--requests--memory"><a href="./values.yaml#L54">loki.write.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
