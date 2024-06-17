# grafana-oncall

![Version: 0.2.2](https://img.shields.io/badge/Version-0.2.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | oncall | 1.7.1 |
| oci://registry-1.docker.io/bitnamicharts | redis | 19.5.3 |

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
			<td id="oncall--celery--resources--limits--cpu"><a href="./values.yaml#L37">oncall.celery.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"200m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--celery--resources--limits--memory"><a href="./values.yaml#L38">oncall.celery.resources.limits.memory</a></td>
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
			<td id="oncall--celery--resources--requests--cpu"><a href="./values.yaml#L40">oncall.celery.resources.requests.cpu</a></td>
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
			<td id="oncall--celery--resources--requests--memory"><a href="./values.yaml#L41">oncall.celery.resources.requests.memory</a></td>
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
			<td id="oncall--cert-manager--enabled"><a href="./values.yaml#L64">oncall.cert-manager.enabled</a></td>
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
			<td id="oncall--database--type"><a href="./values.yaml#L67">oncall.database.type</a></td>
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
"1000m"
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
"512Mi"
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
"150m"
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
"128Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalGrafana--url"><a href="./values.yaml#L109">oncall.externalGrafana.url</a></td>
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
			<td id="oncall--externalPostgresql--db_name"><a href="./values.yaml#L75">oncall.externalPostgresql.db_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"oncall"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalPostgresql--existingSecret"><a href="./values.yaml#L79">oncall.externalPostgresql.existingSecret</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"oncall-postgres-pguser-oncall"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalPostgresql--host"><a href="./values.yaml#L73">oncall.externalPostgresql.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"oncall-postgres-ha.grafana-oncall"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalPostgresql--password"><a href="./values.yaml#L77">oncall.externalPostgresql.password</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
null
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalPostgresql--passwordKey"><a href="./values.yaml#L81">oncall.externalPostgresql.passwordKey</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"password"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalPostgresql--port"><a href="./values.yaml#L74">oncall.externalPostgresql.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
5432
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalPostgresql--user"><a href="./values.yaml#L76">oncall.externalPostgresql.user</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"oncall"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRabbitmq--existingSecret"><a href="./values.yaml#L93">oncall.externalRabbitmq.existingSecret</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-rabbitmq-default-user"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRabbitmq--host"><a href="./values.yaml#L90">oncall.externalRabbitmq.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"rabbitmq.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRabbitmq--passwordKey"><a href="./values.yaml#L94">oncall.externalRabbitmq.passwordKey</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"password"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRabbitmq--port"><a href="./values.yaml#L91">oncall.externalRabbitmq.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
5672
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRabbitmq--usernameKey"><a href="./values.yaml#L95">oncall.externalRabbitmq.usernameKey</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"username"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRedis--existingSecret"><a href="./values.yaml#L102">oncall.externalRedis.existingSecret</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redis"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRedis--host"><a href="./values.yaml#L101">oncall.externalRedis.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redis-master.grafana-oncall"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--externalRedis--passwordKey"><a href="./values.yaml#L103">oncall.externalRedis.passwordKey</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redis-password"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--grafana--enabled"><a href="./values.yaml#L106">oncall.grafana.enabled</a></td>
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
			<td id="oncall--ingress-nginx--enabled"><a href="./values.yaml#L61">oncall.ingress-nginx.enabled</a></td>
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
			<td id="oncall--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L47">oncall.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="oncall--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L49">oncall.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="oncall--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L48">oncall.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="oncall--ingress--enabled"><a href="./values.yaml#L44">oncall.ingress.enabled</a></td>
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
			<td id="oncall--ingress--extraPaths[0]--backend--service--name"><a href="./values.yaml#L56">oncall.ingress.extraPaths[0].backend.service.name</a></td>
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
			<td id="oncall--ingress--extraPaths[0]--backend--service--port--name"><a href="./values.yaml#L58">oncall.ingress.extraPaths[0].backend.service.port.name</a></td>
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
			<td id="oncall--ingress--extraPaths[0]--path"><a href="./values.yaml#L52">oncall.ingress.extraPaths[0].path</a></td>
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
			<td id="oncall--ingress--extraPaths[0]--pathType"><a href="./values.yaml#L53">oncall.ingress.extraPaths[0].pathType</a></td>
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
			<td id="oncall--mariadb--enabled"><a href="./values.yaml#L84">oncall.mariadb.enabled</a></td>
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
			<td id="oncall--migrate--annotations--"argocd--argoproj--io/hook""><a href="./values.yaml#L23">oncall.migrate.annotations."argocd.argoproj.io/hook"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Sync"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--migrate--annotations--"argocd--argoproj--io/hook-delete-policy""><a href="./values.yaml#L24">oncall.migrate.annotations."argocd.argoproj.io/hook-delete-policy"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HookSucceeded"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--migrate--enabled"><a href="./values.yaml#L20">oncall.migrate.enabled</a></td>
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
			<td id="oncall--migrate--resources--limits--cpu"><a href="./values.yaml#L28">oncall.migrate.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"2000m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--migrate--resources--limits--memory"><a href="./values.yaml#L29">oncall.migrate.resources.limits.memory</a></td>
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
			<td id="oncall--migrate--resources--requests--cpu"><a href="./values.yaml#L31">oncall.migrate.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1000m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="oncall--migrate--resources--requests--memory"><a href="./values.yaml#L32">oncall.migrate.resources.requests.memory</a></td>
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
			<td id="oncall--oncall"><a href="./values.yaml#L111">oncall.oncall</a></td>
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
			<td id="oncall--postgresql--enabled"><a href="./values.yaml#L70">oncall.postgresql.enabled</a></td>
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
			<td id="oncall--rabbitmq--enabled"><a href="./values.yaml#L87">oncall.rabbitmq.enabled</a></td>
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
			<td id="oncall--redis--enabled"><a href="./values.yaml#L98">oncall.redis.enabled</a></td>
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
