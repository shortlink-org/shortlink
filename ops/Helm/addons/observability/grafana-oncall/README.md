# grafana-oncall

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | oncall | 1.3.35 |
| oci://registry-1.docker.io/bitnamicharts | redis | 18.0.4 |

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
			<td id="oncall--base_url"><a href="./values.yaml#L8">oncall.base_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--celery--resources--limits--cpu"><a href="./values.yaml#L22">oncall.celery.resources.limits.cpu</a></td>
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
			<td id="oncall--celery--resources--limits--memory"><a href="./values.yaml#L23">oncall.celery.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"100Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--celery--resources--requests--cpu"><a href="./values.yaml#L25">oncall.celery.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"5m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--celery--resources--requests--memory"><a href="./values.yaml#L26">oncall.celery.resources.requests.memory</a></td>
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
			<td id="oncall--cert-manager--enabled"><a href="./values.yaml#L50">oncall.cert-manager.enabled</a></td>
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
			<td id="oncall--database--type"><a href="./values.yaml#L53">oncall.database.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"postgresql"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--enabled"><a href="./values.yaml#L6">oncall.enabled</a></td>
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
			<td id="oncall--engine--resources--limits--cpu"><a href="./values.yaml#L13">oncall.engine.resources.limits.cpu</a></td>
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
			<td id="oncall--engine--resources--limits--memory"><a href="./values.yaml#L14">oncall.engine.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"100Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--engine--resources--requests--cpu"><a href="./values.yaml#L16">oncall.engine.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"5m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--engine--resources--requests--memory"><a href="./values.yaml#L17">oncall.engine.resources.requests.memory</a></td>
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
			<td id="oncall--externalGrafana--url"><a href="./values.yaml#L74">oncall.externalGrafana.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://grafana.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRedis--host"><a href="./values.yaml#L68">oncall.externalRedis.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redis-master.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--grafana--enabled"><a href="./values.yaml#L71">oncall.grafana.enabled</a></td>
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
			<td id="oncall--ingress-nginx--enabled"><a href="./values.yaml#L47">oncall.ingress-nginx.enabled</a></td>
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
			<td id="oncall--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L32">oncall.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="oncall--ingress--annotations--"nginx--ingress--kubernetes--io/enable-modsecurity""><a href="./values.yaml#L33">oncall.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"false"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L35">oncall.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="oncall--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L34">oncall.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="oncall--ingress--enabled"><a href="./values.yaml#L29">oncall.ingress.enabled</a></td>
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
			<td id="oncall--ingress--extraPaths[0]--backend--service--name"><a href="./values.yaml#L42">oncall.ingress.extraPaths[0].backend.service.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ssl-redirect"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--ingress--extraPaths[0]--backend--service--port--name"><a href="./values.yaml#L44">oncall.ingress.extraPaths[0].backend.service.port.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"use-annotation"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--ingress--extraPaths[0]--path"><a href="./values.yaml#L38">oncall.ingress.extraPaths[0].path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/*"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--ingress--extraPaths[0]--pathType"><a href="./values.yaml#L39">oncall.ingress.extraPaths[0].pathType</a></td>
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
			<td id="oncall--mariadb--enabled"><a href="./values.yaml#L59">oncall.mariadb.enabled</a></td>
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
			<td id="oncall--oncall"><a href="./values.yaml#L76">oncall.oncall</a></td>
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
			<td id="oncall--postgresql--enabled"><a href="./values.yaml#L56">oncall.postgresql.enabled</a></td>
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
			<td id="oncall--rabbitmq--enabled"><a href="./values.yaml#L62">oncall.rabbitmq.enabled</a></td>
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
			<td id="oncall--redis--enabled"><a href="./values.yaml#L65">oncall.redis.enabled</a></td>
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
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
