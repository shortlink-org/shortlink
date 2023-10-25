# cilium

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://helm.cilium.io/ | cilium | 1.14.3 |

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
			<td id="cilium--annotateK8sNode"><a href="./values.yaml#L9">cilium.annotateK8sNode</a></td>
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
			<td id="cilium--autoDirectNodeRoutes"><a href="./values.yaml#L8">cilium.autoDirectNodeRoutes</a></td>
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
			<td id="cilium--containerRuntime--integration"><a href="./values.yaml#L12">cilium.containerRuntime.integration</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"auto"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--enableCiliumEndpointSlice"><a href="./values.yaml#L14">cilium.enableCiliumEndpointSlice</a></td>
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
			<td id="cilium--enabled"><a href="./values.yaml#L2">cilium.enabled</a></td>
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
			<td id="cilium--hubble--listenAddress"><a href="./values.yaml#L17">cilium.hubble.listenAddress</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
":4244"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enableOpenMetrics"><a href="./values.yaml#L30">cilium.hubble.metrics.enableOpenMetrics</a></td>
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
			<td id="cilium--hubble--metrics--enabled[0]"><a href="./values.yaml#L21">cilium.hubble.metrics.enabled[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"dns"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[1]"><a href="./values.yaml#L22">cilium.hubble.metrics.enabled[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"tcp"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[2]"><a href="./values.yaml#L23">cilium.hubble.metrics.enabled[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"flow"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[3]"><a href="./values.yaml#L24">cilium.hubble.metrics.enabled[3]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"dns"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[4]"><a href="./values.yaml#L25">cilium.hubble.metrics.enabled[4]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"drop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[5]"><a href="./values.yaml#L26">cilium.hubble.metrics.enabled[5]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"icmp"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[6]"><a href="./values.yaml#L27">cilium.hubble.metrics.enabled[6]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--enabled[7]"><a href="./values.yaml#L28">cilium.hubble.metrics.enabled[7]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"port-distribution"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--hubble--metrics--serviceMonitor--enabled"><a href="./values.yaml#L32">cilium.hubble.metrics.serviceMonitor.enabled</a></td>
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
			<td id="cilium--hubble--relay--enabled"><a href="./values.yaml#L35">cilium.hubble.relay.enabled</a></td>
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
			<td id="cilium--hubble--ui--enabled"><a href="./values.yaml#L38">cilium.hubble.ui.enabled</a></td>
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
			<td id="cilium--k8sServiceHost"><a href="./values.yaml#L5">cilium.k8sServiceHost</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"127.0.0.1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--k8sServicePort"><a href="./values.yaml#L6">cilium.k8sServicePort</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
6443
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cilium--kubeProxyReplacement"><a href="./values.yaml#L4">cilium.kubeProxyReplacement</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"strict"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
