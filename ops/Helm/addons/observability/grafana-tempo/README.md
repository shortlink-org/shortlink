# grafana-tempo

![Version: 0.2.2](https://img.shields.io/badge/Version-0.2.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.30.0 || >= v1.30.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | tempo | 1.14.0 |

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
			<td id="tempo--enabled"><a href="./values.yaml#L2">tempo.enabled</a></td>
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
			<td id="tempo--persistence--enabled"><a href="./values.yaml#L72">tempo.persistence.enabled</a></td>
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
			<td id="tempo--persistence--size"><a href="./values.yaml#L74">tempo.persistence.size</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--persistence--storageClassName"><a href="./values.yaml#L73">tempo.persistence.storageClassName</a></td>
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
			<td id="tempo--serviceMonitor--enabled"><a href="./values.yaml#L66">tempo.serviceMonitor.enabled</a></td>
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
			<td id="tempo--serviceMonitor--labels--release"><a href="./values.yaml#L69">tempo.serviceMonitor.labels.release</a></td>
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
			<td id="tempo--tempo--ingester--max_block_bytes"><a href="./values.yaml#L30">tempo.tempo.ingester.max_block_bytes</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
500000000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--ingester--max_block_duration"><a href="./values.yaml#L31">tempo.tempo.ingester.max_block_duration</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"30m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--ingester--trace_idle_period"><a href="./values.yaml#L29">tempo.tempo.ingester.trace_idle_period</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10s"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--metricsGenerator--enabled"><a href="./values.yaml#L25">tempo.tempo.metricsGenerator.enabled</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"enable"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--metricsGenerator--remoteWriteUrl"><a href="./values.yaml#L26">tempo.tempo.metricsGenerator.remoteWriteUrl</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://prometheus-prometheus.prometheus-operator:9090/api/v1/write"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--querier--max_concurrent_queries"><a href="./values.yaml#L34">tempo.tempo.querier.max_concurrent_queries</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
100
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--querier--search--prefer_self"><a href="./values.yaml#L36">tempo.tempo.querier.search.prefer_self</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
50
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--query_frontend--max_outstanding_per_tenant"><a href="./values.yaml#L41">tempo.tempo.query_frontend.max_outstanding_per_tenant</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
2000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--query_frontend--search--concurrent_jobs"><a href="./values.yaml#L43">tempo.tempo.query_frontend.search.concurrent_jobs</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
2000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--query_frontend--search--target_bytes_per_job"><a href="./values.yaml#L44">tempo.tempo.query_frontend.search.target_bytes_per_job</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
400000000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--reportingEnabled"><a href="./values.yaml#L8">tempo.tempo.reportingEnabled</a></td>
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
			<td id="tempo--tempo--resources--limits--cpu"><a href="./values.yaml#L18">tempo.tempo.resources.limits.cpu</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
2
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--resources--limits--memory"><a href="./values.yaml#L19">tempo.tempo.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"3Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--resources--requests--cpu"><a href="./values.yaml#L21">tempo.tempo.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"500m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--resources--requests--memory"><a href="./values.yaml#L22">tempo.tempo.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"256Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--retention"><a href="./values.yaml#L38">tempo.tempo.retention</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"12h"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--backend"><a href="./values.yaml#L48">tempo.tempo.storage.trace.backend</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--block--bloom_filter_false_positive"><a href="./values.yaml#L51">tempo.tempo.storage.trace.block.bloom_filter_false_positive</a></td>
			<td>
float
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
0.05
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--block--parquet_dedicated_columns[0]--name"><a href="./values.yaml#L55">tempo.tempo.storage.trace.block.parquet_dedicated_columns[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"rpc.service"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--block--parquet_dedicated_columns[0]--scope"><a href="./values.yaml#L54">tempo.tempo.storage.trace.block.parquet_dedicated_columns[0].scope</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"span"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--block--parquet_dedicated_columns[0]--type"><a href="./values.yaml#L56">tempo.tempo.storage.trace.block.parquet_dedicated_columns[0].type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"string"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--block--version"><a href="./values.yaml#L50">tempo.tempo.storage.trace.block.version</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"vParquet4"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--local--path"><a href="./values.yaml#L58">tempo.tempo.storage.trace.local.path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/var/tempo/traces"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--pool--max_workers"><a href="./values.yaml#L62">tempo.tempo.storage.trace.pool.max_workers</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
400
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--pool--queue_depth"><a href="./values.yaml#L63">tempo.tempo.storage.trace.pool.queue_depth</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
30000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempo--storage--trace--wal--path"><a href="./values.yaml#L60">tempo.tempo.storage.trace.wal.path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/var/tempo/wal"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tempo--tempoQuery--enabled"><a href="./values.yaml#L5">tempo.tempoQuery.enabled</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
