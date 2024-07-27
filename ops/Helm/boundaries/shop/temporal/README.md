# temporal

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://go.temporal.io/helm-charts | temporal | 0.44.0 |
| https://scylla-operator-charts.storage.googleapis.com/stable | scylla | v1.13.0 |

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
			<td id="temporal--cassandra--enabled"><a href="./values.yaml#L78">temporal.cassandra.enabled</a></td>
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
			<td id="temporal--elasticsearch--enabled"><a href="./values.yaml#L74">temporal.elasticsearch.enabled</a></td>
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
			<td id="temporal--elasticsearch--replicas"><a href="./values.yaml#L75">temporal.elasticsearch.replicas</a></td>
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
			<td id="temporal--enabled"><a href="./values.yaml#L6">temporal.enabled</a></td>
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
			<td id="temporal--fullnameOverride"><a href="./values.yaml#L8">temporal.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"temporal"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--grafana--enabled"><a href="./values.yaml#L84">temporal.grafana.enabled</a></td>
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
			<td id="temporal--prometheus--enabled"><a href="./values.yaml#L81">temporal.prometheus.enabled</a></td>
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
			<td id="temporal--server--config--persistence--default--cassandra--existingSecret"><a href="./values.yaml#L21">temporal.server.config.persistence.default.cassandra.existingSecret</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"scylla-cluster-auth-token"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--server--config--persistence--default--cassandra--hosts[0]"><a href="./values.yaml#L20">temporal.server.config.persistence.default.cassandra.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"scylla-cluster.temporal"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--server--config--persistence--default--cassandra--port"><a href="./values.yaml#L18">temporal.server.config.persistence.default.cassandra.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
9042
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--server--config--persistence--visibility--cassandra--existingSecret"><a href="./values.yaml#L28">temporal.server.config.persistence.visibility.cassandra.existingSecret</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"scylla-cluster-auth-token"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--server--config--persistence--visibility--cassandra--hosts[0]"><a href="./values.yaml#L27">temporal.server.config.persistence.visibility.cassandra.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"scylla-cluster.temporal"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--server--config--persistence--visibility--cassandra--port"><a href="./values.yaml#L25">temporal.server.config.persistence.visibility.cassandra.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
9042
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--server--frontend--metrics--serviceMonitor--enabled"><a href="./values.yaml#L40">temporal.server.frontend.metrics.serviceMonitor.enabled</a></td>
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
			<td id="temporal--server--history--metrics--serviceMonitor--enabled"><a href="./values.yaml#L45">temporal.server.history.metrics.serviceMonitor.enabled</a></td>
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
			<td id="temporal--server--matching--metrics--serviceMonitor--enabled"><a href="./values.yaml#L50">temporal.server.matching.metrics.serviceMonitor.enabled</a></td>
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
			<td id="temporal--server--metrics--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L35">temporal.server.metrics.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="temporal--server--metrics--serviceMonitor--enabled"><a href="./values.yaml#L32">temporal.server.metrics.serviceMonitor.enabled</a></td>
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
			<td id="temporal--server--worker--metrics--serviceMonitor--enabled"><a href="./values.yaml#L55">temporal.server.worker.metrics.serviceMonitor.enabled</a></td>
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
			<td id="temporal--serviceAccount--create"><a href="./values.yaml#L11">temporal.serviceAccount.create</a></td>
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
			<td id="temporal--web--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L61">temporal.web.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="temporal--web--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L63">temporal.web.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="temporal--web--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L62">temporal.web.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="temporal--web--ingress--enabled"><a href="./values.yaml#L59">temporal.web.ingress.enabled</a></td>
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
			<td id="temporal--web--ingress--hosts[0]"><a href="./values.yaml#L66">temporal.web.ingress.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"temporal.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--web--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L71">temporal.web.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"temporal.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="temporal--web--ingress--tls[0]--secretName"><a href="./values.yaml#L69">temporal.web.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"temporal-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
