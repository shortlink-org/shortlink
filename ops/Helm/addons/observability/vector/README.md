# oneuptime

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://helm.vector.dev | vector | 0.34.0 |

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
			<td id="vector--customConfig--api--address"><a href="./values.yaml#L25">vector.customConfig.api.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"127.0.0.1:8686"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--api--enabled"><a href="./values.yaml#L24">vector.customConfig.api.enabled</a></td>
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
			<td id="vector--customConfig--api--playground"><a href="./values.yaml#L26">vector.customConfig.api.playground</a></td>
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
			<td id="vector--customConfig--data_dir"><a href="./values.yaml#L22">vector.customConfig.data_dir</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/vector-data-dir"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--compression"><a href="./values.yaml#L58">vector.customConfig.sinks.loki.compression</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"snappy"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--encoding--codec"><a href="./values.yaml#L57">vector.customConfig.sinks.loki.encoding.codec</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"native_json"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--endpoint"><a href="./values.yaml#L61">vector.customConfig.sinks.loki.endpoint</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-loki.grafana:3100"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--inputs[0]"><a href="./values.yaml#L55">vector.customConfig.sinks.loki.inputs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kubernetes_logs"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--labels--*"><a href="./values.yaml#L65">vector.customConfig.sinks.loki.labels.*</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"{{ print \"{{ metadata }}\" }}"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--labels--pod_labels_*"><a href="./values.yaml#L63">vector.customConfig.sinks.loki.labels.pod_labels_*</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"{{ print \"{{ kubernetes.pod_labels }}\" }}"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--loki--remove_label_fields"><a href="./values.yaml#L59">vector.customConfig.sinks.loki.remove_label_fields</a></td>
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
			<td id="vector--customConfig--sinks--loki--remove_timestamp"><a href="./values.yaml#L60">vector.customConfig.sinks.loki.remove_timestamp</a></td>
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
			<td id="vector--customConfig--sinks--loki--type"><a href="./values.yaml#L54">vector.customConfig.sinks.loki.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"loki"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--prom_exporter--address"><a href="./values.yaml#L52">vector.customConfig.sinks.prom_exporter.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.0.0.0:9090"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--prom_exporter--inputs[0]"><a href="./values.yaml#L51">vector.customConfig.sinks.prom_exporter.inputs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"host_metrics"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--prom_exporter--inputs[1]"><a href="./values.yaml#L51">vector.customConfig.sinks.prom_exporter.inputs[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"internal_metrics"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sinks--prom_exporter--type"><a href="./values.yaml#L50">vector.customConfig.sinks.prom_exporter.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"prometheus_exporter"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sources--host_metrics--filesystem--devices--excludes[0]"><a href="./values.yaml#L34">vector.customConfig.sources.host_metrics.filesystem.devices.excludes[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"binfmt_misc"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sources--host_metrics--filesystem--filesystems--excludes[0]"><a href="./values.yaml#L36">vector.customConfig.sources.host_metrics.filesystem.filesystems.excludes[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"binfmt_misc"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sources--host_metrics--filesystem--mountpoints--excludes[0]"><a href="./values.yaml#L38">vector.customConfig.sources.host_metrics.filesystem.mountpoints.excludes[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"*/proc/sys/fs/binfmt_misc"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sources--host_metrics--type"><a href="./values.yaml#L39">vector.customConfig.sources.host_metrics.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"host_metrics"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sources--internal_metrics--type"><a href="./values.yaml#L41">vector.customConfig.sources.internal_metrics.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"internal_metrics"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--customConfig--sources--kubernetes_logs--type"><a href="./values.yaml#L29">vector.customConfig.sources.kubernetes_logs.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kubernetes_logs"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--enabled"><a href="./values.yaml#L6">vector.enabled</a></td>
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
			<td id="vector--podMonitor--enabled"><a href="./values.yaml#L19">vector.podMonitor.enabled</a></td>
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
			<td id="vector--resources--limits--cpu"><a href="./values.yaml#L15">vector.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"200m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--resources--limits--memory"><a href="./values.yaml#L16">vector.resources.limits.memory</a></td>
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
			<td id="vector--resources--requests--cpu"><a href="./values.yaml#L12">vector.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"200m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="vector--resources--requests--memory"><a href="./values.yaml#L13">vector.resources.requests.memory</a></td>
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
			<td id="vector--role"><a href="./values.yaml#L8">vector.role</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Agent"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
