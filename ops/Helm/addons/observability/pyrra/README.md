# pyrra

![Version: 0.2.3](https://img.shields.io/badge/Version-0.2.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://rlex.github.io/helm-charts/ | pyrra | 0.13.0 |

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
			<td id="pyrra--enabled"><a href="./values.yaml#L6">pyrra.enabled</a></td>
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
			<td id="pyrra--extraKubernetesArgs[0]"><a href="./values.yaml#L17">pyrra.extraKubernetesArgs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"--config-map-mode=true"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--fullnameOverride"><a href="./values.yaml#L8">pyrra.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"pyrra"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--genericRules--enabled"><a href="./values.yaml#L53">pyrra.genericRules.enabled</a></td>
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
			<td id="pyrra--image--tag"><a href="./values.yaml#L11">pyrra.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"v0.7.2"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L23">pyrra.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="pyrra--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L25">pyrra.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="pyrra--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L24">pyrra.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="pyrra--ingress--className"><a href="./values.yaml#L21">pyrra.ingress.className</a></td>
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
			<td id="pyrra--ingress--enabled"><a href="./values.yaml#L20">pyrra.ingress.enabled</a></td>
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
			<td id="pyrra--ingress--hosts[0]--host"><a href="./values.yaml#L27">pyrra.ingress.hosts[0].host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"status.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--ingress--hosts[0]--paths[0]--path"><a href="./values.yaml#L29">pyrra.ingress.hosts[0].paths[0].path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--ingress--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L30">pyrra.ingress.hosts[0].paths[0].pathType</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Prefix"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L34">pyrra.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"status.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--ingress--tls[0]--secretName"><a href="./values.yaml#L32">pyrra.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"status-page-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--prometheusExternalUrl"><a href="./values.yaml#L14">pyrra.prometheusExternalUrl</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/prometheus"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--prometheusUrl"><a href="./values.yaml#L13">pyrra.prometheusUrl</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://prometheus-prometheus.prometheus-operator:9090"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--resources--limits--cpu"><a href="./values.yaml#L38">pyrra.resources.limits.cpu</a></td>
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
			<td id="pyrra--resources--limits--memory"><a href="./values.yaml#L39">pyrra.resources.limits.memory</a></td>
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
			<td id="pyrra--resources--requests--cpu"><a href="./values.yaml#L41">pyrra.resources.requests.cpu</a></td>
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
			<td id="pyrra--resources--requests--memory"><a href="./values.yaml#L42">pyrra.resources.requests.memory</a></td>
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
			<td id="pyrra--serviceMonitor--enabled"><a href="./values.yaml#L45">pyrra.serviceMonitor.enabled</a></td>
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
			<td id="pyrra--serviceMonitor--jobLabel"><a href="./values.yaml#L47">pyrra.serviceMonitor.jobLabel</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"pyrra"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyrra--serviceMonitor--labels--release"><a href="./values.yaml#L50">pyrra.serviceMonitor.labels.release</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
