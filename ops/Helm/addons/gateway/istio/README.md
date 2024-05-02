# istio

![Version: 0.3.1](https://img.shields.io/badge/Version-0.3.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://istio-release.storage.googleapis.com/charts | base | 1.21.2 |
| https://istio-release.storage.googleapis.com/charts | istiod | 1.21.2 |

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
			<td id="base--enabled"><a href="./values.yaml#L2">base.enabled</a></td>
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
			<td id="istiod--defaults--global--istiod--enableAnalysis"><a href="./values.yaml#L12">istiod.defaults.global.istiod.enableAnalysis</a></td>
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
			<td id="istiod--defaults--global--logAsJson"><a href="./values.yaml#L9">istiod.defaults.global.logAsJson</a></td>
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
			<td id="istiod--defaults--global--tracer--zipkin--address"><a href="./values.yaml#L16">istiod.defaults.global.tracer.zipkin.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana:9411"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--istio_cni--enabled"><a href="./values.yaml#L27">istiod.defaults.istio_cni.enabled</a></td>
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
			<td id="istiod--defaults--meshConfig--enablePrometheusMerge"><a href="./values.yaml#L33">istiod.defaults.meshConfig.enablePrometheusMerge</a></td>
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
			<td id="istiod--defaults--meshConfig--enableTracing"><a href="./values.yaml#L32">istiod.defaults.meshConfig.enableTracing</a></td>
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
			<td id="istiod--defaults--meshConfig--extensionProviders[0]--envoyOtelAls--port"><a href="./values.yaml#L38">istiod.defaults.meshConfig.extensionProviders[0].envoyOtelAls.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--meshConfig--extensionProviders[0]--envoyOtelAls--service"><a href="./values.yaml#L37">istiod.defaults.meshConfig.extensionProviders[0].envoyOtelAls.service</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana.svc.cluster.local"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--meshConfig--extensionProviders[0]--name"><a href="./values.yaml#L35">istiod.defaults.meshConfig.extensionProviders[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"otel"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--meshConfig--extensionProviders[1]--name"><a href="./values.yaml#L39">istiod.defaults.meshConfig.extensionProviders[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"otel-tracing"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--meshConfig--extensionProviders[1]--opentelemetry--port"><a href="./values.yaml#L42">istiod.defaults.meshConfig.extensionProviders[1].opentelemetry.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--meshConfig--extensionProviders[1]--opentelemetry--service"><a href="./values.yaml#L41">istiod.defaults.meshConfig.extensionProviders[1].opentelemetry.service</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"otel-collector-collector.grafana.svc.cluster.local"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--pilot--env--ENABLE_NATIVE_SIDECARS"><a href="./values.yaml#L23">istiod.defaults.pilot.env.ENABLE_NATIVE_SIDECARS</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--pilot--env--ISTIO_ENABLE_CONTROLLER_QUEUE_METRICS"><a href="./values.yaml#L24">istiod.defaults.pilot.env.ISTIO_ENABLE_CONTROLLER_QUEUE_METRICS</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="istiod--defaults--pilot--traceSampling"><a href="./values.yaml#L20">istiod.defaults.pilot.traceSampling</a></td>
			<td>
float
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
			<td id="istiod--defaults--telemetry--v2--metadataExchange--wasmEnabled"><a href="./values.yaml#L47">istiod.defaults.telemetry.v2.metadataExchange.wasmEnabled</a></td>
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
			<td id="istiod--defaults--telemetry--v2--prometheus--enabled"><a href="./values.yaml#L49">istiod.defaults.telemetry.v2.prometheus.enabled</a></td>
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
			<td id="istiod--defaults--telemetry--v2--prometheus--wasmEnabled"><a href="./values.yaml#L50">istiod.defaults.telemetry.v2.prometheus.wasmEnabled</a></td>
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
			<td id="istiod--enabled"><a href="./values.yaml#L5">istiod.enabled</a></td>
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
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
