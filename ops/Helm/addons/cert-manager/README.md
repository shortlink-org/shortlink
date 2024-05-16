# cert-manager

![Version: 0.3.3](https://img.shields.io/badge/Version-0.3.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://charts.jetstack.io | certmanager(cert-manager) | v1.14.5 |
| https://charts.jetstack.io | cert-manager-approver-policy | v0.14.1 |
| https://charts.jetstack.io | spiffe(cert-manager-csi-driver-spiffe) | v0.6.0 |

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
			<td id="annotations"><a href="./values.yaml#L7">annotations</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{}
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="cert-manager-approver-policy--app--metrics--service--servicemonitor--enabled"><a href="./values.yaml#L39">cert-manager-approver-policy.app.metrics.service.servicemonitor.enabled</a></td>
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
			<td id="cert-manager-approver-policy--enabled"><a href="./values.yaml#L33">cert-manager-approver-policy.enabled</a></td>
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
			<td id="certmanager--enabled"><a href="./values.yaml#L10">certmanager.enabled</a></td>
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
			<td id="certmanager--extraArgs[0]"><a href="./values.yaml#L21">certmanager.extraArgs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"--logging-format=json"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--featureGates"><a href="./values.yaml#L12">certmanager.featureGates</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ExperimentalGatewayAPISupport=true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--installCRDs"><a href="./values.yaml#L18">certmanager.installCRDs</a></td>
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
			<td id="certmanager--prometheus--enabled"><a href="./values.yaml#L25">certmanager.prometheus.enabled</a></td>
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
			<td id="certmanager--prometheus--servicemonitor--enabled"><a href="./values.yaml#L28">certmanager.prometheus.servicemonitor.enabled</a></td>
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
			<td id="certmanager--prometheus--servicemonitor--labels--release"><a href="./values.yaml#L30">certmanager.prometheus.servicemonitor.labels.release</a></td>
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
			<td id="certmanager--type"><a href="./values.yaml#L16">certmanager.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"cloudflare"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="email"><a href="./values.yaml#L5">email</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"mymail@gmail.com"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="spiffe--app--approver--metrics--service--servicemonitor--enabled"><a href="./values.yaml#L58">spiffe.app.approver.metrics.service.servicemonitor.enabled</a></td>
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
			<td id="spiffe--app--driver--resources--limits--cpu"><a href="./values.yaml#L51">spiffe.app.driver.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"100m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="spiffe--app--driver--resources--limits--memory"><a href="./values.yaml#L52">spiffe.app.driver.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"128Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="spiffe--app--driver--resources--requests--cpu"><a href="./values.yaml#L48">spiffe.app.driver.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"15m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="spiffe--app--driver--resources--requests--memory"><a href="./values.yaml#L49">spiffe.app.driver.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"5Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="spiffe--enabled"><a href="./values.yaml#L42">spiffe.enabled</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
