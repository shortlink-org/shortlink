# grafana-agent

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | k8s-monitoring | 0.6.1 |

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
			<td id="k8s-monitoring--cluster--name"><a href="./values.yaml#L12">k8s-monitoring.cluster.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ShortLink"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--enabled"><a href="./values.yaml#L6">k8s-monitoring.enabled</a></td>
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
			<td id="k8s-monitoring--externalServices--loki--host"><a href="./values.yaml#L17">k8s-monitoring.externalServices.loki.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://logs-prod-012.grafana.net"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--externalServices--prometheus--host"><a href="./values.yaml#L15">k8s-monitoring.externalServices.prometheus.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://prometheus-prod-24-prod-eu-west-2.grafana.net"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--externalServices--tempo--host"><a href="./values.yaml#L19">k8s-monitoring.externalServices.tempo.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://tempo-prod-10-prod-eu-west-2.grafana.net:443"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[0]--name"><a href="./values.yaml#L32">k8s-monitoring.grafana-agent.agent.extraPorts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"otlp-grpc"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[0]--port"><a href="./values.yaml#L33">k8s-monitoring.grafana-agent.agent.extraPorts[0].port</a></td>
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
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[0]--protocol"><a href="./values.yaml#L35">k8s-monitoring.grafana-agent.agent.extraPorts[0].protocol</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TCP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[0]--targetPort"><a href="./values.yaml#L34">k8s-monitoring.grafana-agent.agent.extraPorts[0].targetPort</a></td>
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
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[1]--name"><a href="./values.yaml#L36">k8s-monitoring.grafana-agent.agent.extraPorts[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"otlp-http"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[1]--port"><a href="./values.yaml#L37">k8s-monitoring.grafana-agent.agent.extraPorts[1].port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4318
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[1]--protocol"><a href="./values.yaml#L39">k8s-monitoring.grafana-agent.agent.extraPorts[1].protocol</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TCP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[1]--targetPort"><a href="./values.yaml#L38">k8s-monitoring.grafana-agent.agent.extraPorts[1].targetPort</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4318
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[2]--name"><a href="./values.yaml#L40">k8s-monitoring.grafana-agent.agent.extraPorts[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"zipkin"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[2]--port"><a href="./values.yaml#L41">k8s-monitoring.grafana-agent.agent.extraPorts[2].port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
9411
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[2]--protocol"><a href="./values.yaml#L43">k8s-monitoring.grafana-agent.agent.extraPorts[2].protocol</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TCP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--grafana-agent--agent--extraPorts[2]--targetPort"><a href="./values.yaml#L42">k8s-monitoring.grafana-agent.agent.extraPorts[2].targetPort</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
9411
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--opencost--opencost--exporter--defaultClusterId"><a href="./values.yaml#L23">k8s-monitoring.opencost.opencost.exporter.defaultClusterId</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ShortLink"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--opencost--opencost--prometheus--external--url"><a href="./values.yaml#L26">k8s-monitoring.opencost.opencost.prometheus.external.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://prometheus-prod-24-prod-eu-west-2.grafana.net/api/prom"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="k8s-monitoring--prometheus-operator-crds--enabled"><a href="./values.yaml#L9">k8s-monitoring.prometheus-operator-crds.enabled</a></td>
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
			<td id="k8s-monitoring--traces--enabled"><a href="./values.yaml#L28">k8s-monitoring.traces.enabled</a></td>
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
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
