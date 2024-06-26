# argo-rollouts

![Version: 0.3.18](https://img.shields.io/badge/Version-0.3.18-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.9.0](https://img.shields.io/badge/AppVersion-2.9.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-rollouts | 2.37.1 |

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
			<td id="argo-rollouts--controller--replicas"><a href="./values.yaml#L7">argo-rollouts.controller.replicas</a></td>
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
			<td id="argo-rollouts--controller--resources--limits--cpu"><a href="./values.yaml#L16">argo-rollouts.controller.resources.limits.cpu</a></td>
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
			<td id="argo-rollouts--controller--resources--limits--memory"><a href="./values.yaml#L17">argo-rollouts.controller.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"2Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--controller--resources--requests--cpu"><a href="./values.yaml#L19">argo-rollouts.controller.resources.requests.cpu</a></td>
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
			<td id="argo-rollouts--controller--resources--requests--memory"><a href="./values.yaml#L20">argo-rollouts.controller.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"50Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--controller--trafficRouterPlugins--trafficRouterPlugins"><a href="./values.yaml#L10">argo-rollouts.controller.trafficRouterPlugins.trafficRouterPlugins</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"- name: \"argoproj-labs/gatewayAPI\"\n  location: \"https://github.com/argoproj-labs/rollouts-plugin-trafficrouter-gatewayapi/releases/download/v0.3.0/gateway-api-plugin-linux-amd64\""
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--enabled"><a href="./values.yaml#L23">argo-rollouts.dashboard.enabled</a></td>
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
			<td id="argo-rollouts--dashboard--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L40">argo-rollouts.dashboard.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"cert-manager-production"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/backend-protocol""><a href="./values.yaml#L41">argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HTTP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L43">argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="argo-rollouts--dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L42">argo-rollouts.dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="argo-rollouts--dashboard--ingress--enabled"><a href="./values.yaml#L35">argo-rollouts.dashboard.ingress.enabled</a></td>
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
			<td id="argo-rollouts--dashboard--ingress--hosts[0]"><a href="./values.yaml#L46">argo-rollouts.dashboard.ingress.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--ingress--ingressClassName"><a href="./values.yaml#L37">argo-rollouts.dashboard.ingress.ingressClassName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"nginx"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--ingress--paths[0]"><a href="./values.yaml#L49">argo-rollouts.dashboard.ingress.paths[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/rollouts"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L54">argo-rollouts.dashboard.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--ingress--tls[0]--secretName"><a href="./values.yaml#L52">argo-rollouts.dashboard.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--dashboard--readonly"><a href="./values.yaml#L24">argo-rollouts.dashboard.readonly</a></td>
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
			<td id="argo-rollouts--dashboard--resources--limits--cpu"><a href="./values.yaml#L28">argo-rollouts.dashboard.resources.limits.cpu</a></td>
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
			<td id="argo-rollouts--dashboard--resources--limits--memory"><a href="./values.yaml#L29">argo-rollouts.dashboard.resources.limits.memory</a></td>
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
			<td id="argo-rollouts--dashboard--resources--requests--cpu"><a href="./values.yaml#L31">argo-rollouts.dashboard.resources.requests.cpu</a></td>
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
			<td id="argo-rollouts--dashboard--resources--requests--memory"><a href="./values.yaml#L32">argo-rollouts.dashboard.resources.requests.memory</a></td>
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
			<td id="argo-rollouts--enabled"><a href="./values.yaml#L2">argo-rollouts.enabled</a></td>
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
			<td id="argo-rollouts--fullnameOverride"><a href="./values.yaml#L4">argo-rollouts.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo-rollouts"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-rollouts--metrics--enabled"><a href="./values.yaml#L57">argo-rollouts.metrics.enabled</a></td>
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
			<td id="argo-rollouts--metrics--serviceMonitor--enabled"><a href="./values.yaml#L59">argo-rollouts.metrics.serviceMonitor.enabled</a></td>
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
