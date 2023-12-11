# kubescape

![Version: 1.16.2](https://img.shields.io/badge/Version-1.16.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kubescape.github.io/helm-charts/ | kubescape(kubescape-operator) | 1.16.5 |

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
			<td id="kubescape--capabilities--continuousScan"><a href="./values.yaml#L14">kubescape.capabilities.continuousScan</a></td>
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
			<td id="kubescape--capabilities--networkGenerator"><a href="./values.yaml#L17">kubescape.capabilities.networkGenerator</a></td>
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
			<td id="kubescape--capabilities--networkPolicyService"><a href="./values.yaml#L16">kubescape.capabilities.networkPolicyService</a></td>
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
			<td id="kubescape--capabilities--runtimeObservability"><a href="./values.yaml#L15">kubescape.capabilities.runtimeObservability</a></td>
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
			<td id="kubescape--capabilities--seccompGenerator"><a href="./values.yaml#L18">kubescape.capabilities.seccompGenerator</a></td>
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
			<td id="kubescape--capabilities--vexGeneration"><a href="./values.yaml#L19">kubescape.capabilities.vexGeneration</a></td>
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
			<td id="kubescape--clusterName"><a href="./values.yaml#L21">kubescape.clusterName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kubescape--configurations--otelUrl"><a href="./values.yaml#L31">kubescape.configurations.otelUrl</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-tempo.grafana:4317"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kubescape--enabled"><a href="./values.yaml#L6">kubescape.enabled</a></td>
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
			<td id="kubescape--global--networkPolicy--createEgressRules"><a href="./values.yaml#L11">kubescape.global.networkPolicy.createEgressRules</a></td>
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
			<td id="kubescape--global--networkPolicy--enabled"><a href="./values.yaml#L10">kubescape.global.networkPolicy.enabled</a></td>
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
			<td id="kubescape--kubescape--labels--release"><a href="./values.yaml#L25">kubescape.kubescape.labels.release</a></td>
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
			<td id="kubescape--kubescape--serviceMonitor--enabled"><a href="./values.yaml#L28">kubescape.kubescape.serviceMonitor.enabled</a></td>
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
			<td id="kubescape--operator--resources--requests--cpu"><a href="./values.yaml#L39">kubescape.operator.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"70m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kubescape--operator--resources--requests--memory"><a href="./values.yaml#L40">kubescape.operator.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"120Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kubescape--persistence--storageClass"><a href="./values.yaml#L34">kubescape.persistence.storageClass</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
