# tracetest

![Version: 0.2.3](https://img.shields.io/badge/Version-0.2.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kubeshop.github.io/helm-charts | tracetest | 0.3.23 |

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
			<td id="tracetest--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L47">tracetest.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="tracetest--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L49">tracetest.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="tracetest--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L48">tracetest.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="tracetest--ingress--className"><a href="./values.yaml#L45">tracetest.ingress.className</a></td>
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
			<td id="tracetest--ingress--enabled"><a href="./values.yaml#L44">tracetest.ingress.enabled</a></td>
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
			<td id="tracetest--ingress--hosts[0]--host"><a href="./values.yaml#L51">tracetest.ingress.hosts[0].host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"tracetest.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--ingress--hosts[0]--paths[0]--path"><a href="./values.yaml#L53">tracetest.ingress.hosts[0].paths[0].path</a></td>
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
			<td id="tracetest--ingress--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L54">tracetest.ingress.hosts[0].paths[0].pathType</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ImplementationSpecific"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L58">tracetest.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"tracetest.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--ingress--tls[0]--secretName"><a href="./values.yaml#L56">tracetest.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"tracetest-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--postgresql--enabled"><a href="./values.yaml#L7">tracetest.postgresql.enabled</a></td>
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
			<td id="tracetest--postgresql--global--storageClass"><a href="./values.yaml#L10">tracetest.postgresql.global.storageClass</a></td>
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
			<td id="tracetest--postgresql--metrics--enabled"><a href="./values.yaml#L19">tracetest.postgresql.metrics.enabled</a></td>
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
			<td id="tracetest--postgresql--metrics--prometheusRule--enabled"><a href="./values.yaml#L27">tracetest.postgresql.metrics.prometheusRule.enabled</a></td>
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
			<td id="tracetest--postgresql--metrics--prometheusRule--labels--release"><a href="./values.yaml#L29">tracetest.postgresql.metrics.prometheusRule.labels.release</a></td>
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
			<td id="tracetest--postgresql--metrics--serviceMonitor--enabled"><a href="./values.yaml#L22">tracetest.postgresql.metrics.serviceMonitor.enabled</a></td>
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
			<td id="tracetest--postgresql--metrics--serviceMonitor--labels--release"><a href="./values.yaml#L24">tracetest.postgresql.metrics.serviceMonitor.labels.release</a></td>
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
			<td id="tracetest--postgresql--primary--persistence--size"><a href="./values.yaml#L34">tracetest.postgresql.primary.persistence.size</a></td>
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
			<td id="tracetest--postgresql--primary--persistence--storageClass"><a href="./values.yaml#L33">tracetest.postgresql.primary.persistence.storageClass</a></td>
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
			<td id="tracetest--postgresql--serviceAccount--create"><a href="./values.yaml#L16">tracetest.postgresql.serviceAccount.create</a></td>
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
			<td id="tracetest--postgresql--volumePermissions--enabled"><a href="./values.yaml#L13">tracetest.postgresql.volumePermissions.enabled</a></td>
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
			<td id="tracetest--provisioning"><a href="./values.yaml#L68">tracetest.provisioning</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"type: DataStore\nspec:\n  name: Grafana Tempo\n  type: tempo\n  default: true\n  tempo:\n    type: http\n    http:\n      url: http://grafana-tempo.grafana:3100\n      tls:\n        insecure: true\n---\ntype: Test\nspec:\n  id: dSzgkfKIR\n  name: \"API: get links\"\n  trigger:\n    type: http\n    httpRequest:\n      method: GET\n      url: https://shortlink.best/api/links\n      headers:\n        - key: Content-Type\n          value: application/json\n---\ntype: TestSuite\nspec:\n  id: 2e3YoYKSR\n  name: shortlink-link\n  description: Link boundary\n  steps:\n    - dSzgkfKIR"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--resources--limits--cpu"><a href="./values.yaml#L62">tracetest.resources.limits.cpu</a></td>
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
			<td id="tracetest--resources--limits--memory"><a href="./values.yaml#L63">tracetest.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"512Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--resources--requests--cpu"><a href="./values.yaml#L65">tracetest.resources.requests.cpu</a></td>
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
			<td id="tracetest--resources--requests--memory"><a href="./values.yaml#L66">tracetest.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"126Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="tracetest--telemetry--exporters--collector--exporter--collector--endpoint"><a href="./values.yaml#L41">tracetest.telemetry.exporters.collector.exporter.collector.endpoint</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
