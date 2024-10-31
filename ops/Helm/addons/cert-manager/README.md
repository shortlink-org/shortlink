# cert-manager

![Version: 0.4.1](https://img.shields.io/badge/Version-0.4.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.1](https://img.shields.io/badge/AppVersion-1.16.1-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.30.0 || >= v1.30.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://charts.jetstack.io | certmanager(cert-manager) | v1.16.1 |
| https://charts.jetstack.io | cert-manager-approver-policy | v0.16.0 |
| https://charts.jetstack.io | spiffe(cert-manager-csi-driver-spiffe) | v0.8.1 |

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
			<td id="cert-manager-approver-policy--app--metrics--service--servicemonitor--enabled"><a href="./values.yaml#L76">cert-manager-approver-policy.app.metrics.service.servicemonitor.enabled</a></td>
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
			<td id="cert-manager-approver-policy--enabled"><a href="./values.yaml#L70">cert-manager-approver-policy.enabled</a></td>
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
			<td id="certmanager--cainjector--resources--limits--cpu"><a href="./values.yaml#L49">certmanager.cainjector.resources.limits.cpu</a></td>
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
			<td id="certmanager--cainjector--resources--limits--memory"><a href="./values.yaml#L50">certmanager.cainjector.resources.limits.memory</a></td>
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
			<td id="certmanager--cainjector--resources--requests--cpu"><a href="./values.yaml#L46">certmanager.cainjector.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--cainjector--resources--requests--memory"><a href="./values.yaml#L47">certmanager.cainjector.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"32Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--crds--enabled"><a href="./values.yaml#L27">certmanager.crds.enabled</a></td>
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
			<td id="certmanager--crds--keep"><a href="./values.yaml#L28">certmanager.crds.keep</a></td>
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
			<td id="certmanager--enabled"><a href="./values.yaml#L14">certmanager.enabled</a></td>
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
			<td id="certmanager--extraArgs[0]"><a href="./values.yaml#L31">certmanager.extraArgs[0]</a></td>
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
			<td id="certmanager--featureGates"><a href="./values.yaml#L24">certmanager.featureGates</a></td>
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
			<td id="certmanager--prometheus--enabled"><a href="./values.yaml#L62">certmanager.prometheus.enabled</a></td>
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
			<td id="certmanager--prometheus--servicemonitor--enabled"><a href="./values.yaml#L65">certmanager.prometheus.servicemonitor.enabled</a></td>
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
			<td id="certmanager--prometheus--servicemonitor--labels--release"><a href="./values.yaml#L67">certmanager.prometheus.servicemonitor.labels.release</a></td>
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
			<td id="certmanager--resources--limits--cpu"><a href="./values.yaml#L21">certmanager.resources.limits.cpu</a></td>
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
			<td id="certmanager--resources--limits--memory"><a href="./values.yaml#L22">certmanager.resources.limits.memory</a></td>
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
			<td id="certmanager--resources--requests--cpu"><a href="./values.yaml#L18">certmanager.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--resources--requests--memory"><a href="./values.yaml#L19">certmanager.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"32Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--startupapicheck--resources--limits--cpu"><a href="./values.yaml#L58">certmanager.startupapicheck.resources.limits.cpu</a></td>
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
			<td id="certmanager--startupapicheck--resources--limits--memory"><a href="./values.yaml#L59">certmanager.startupapicheck.resources.limits.memory</a></td>
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
			<td id="certmanager--startupapicheck--resources--requests--cpu"><a href="./values.yaml#L55">certmanager.startupapicheck.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--startupapicheck--resources--requests--memory"><a href="./values.yaml#L56">certmanager.startupapicheck.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"32Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--webhook--resources--limits--cpu"><a href="./values.yaml#L40">certmanager.webhook.resources.limits.cpu</a></td>
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
			<td id="certmanager--webhook--resources--limits--memory"><a href="./values.yaml#L41">certmanager.webhook.resources.limits.memory</a></td>
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
			<td id="certmanager--webhook--resources--requests--cpu"><a href="./values.yaml#L37">certmanager.webhook.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"10m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="certmanager--webhook--resources--requests--memory"><a href="./values.yaml#L38">certmanager.webhook.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"32Mi"
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
			<td id="spiffe--app--approver--metrics--service--servicemonitor--enabled"><a href="./values.yaml#L95">spiffe.app.approver.metrics.service.servicemonitor.enabled</a></td>
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
			<td id="spiffe--app--driver--resources--limits--cpu"><a href="./values.yaml#L88">spiffe.app.driver.resources.limits.cpu</a></td>
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
			<td id="spiffe--app--driver--resources--limits--memory"><a href="./values.yaml#L89">spiffe.app.driver.resources.limits.memory</a></td>
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
			<td id="spiffe--app--driver--resources--requests--cpu"><a href="./values.yaml#L85">spiffe.app.driver.resources.requests.cpu</a></td>
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
			<td id="spiffe--app--driver--resources--requests--memory"><a href="./values.yaml#L86">spiffe.app.driver.resources.requests.memory</a></td>
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
			<td id="spiffe--enabled"><a href="./values.yaml#L79">spiffe.enabled</a></td>
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
			<td id="type"><a href="./values.yaml#L11">type</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
