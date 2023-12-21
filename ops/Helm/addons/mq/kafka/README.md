# kafka

![Version: 0.3.0](https://img.shields.io/badge/Version-0.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://provectus.github.io/kafka-ui-charts | kafka-ui | 0.7.5 |
| https://strimzi.io/charts/ | strimzi-kafka-operator | 0.39.0 |

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
			<td id="kafka-ui--enabled"><a href="./values.yaml#L23">kafka-ui.enabled</a></td>
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
			<td id="kafka-ui--envs--config--KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS"><a href="./values.yaml#L35">kafka-ui.envs.config.KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink-kafka-bootstrap:9092"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--envs--config--KAFKA_CLUSTERS_0_NAME"><a href="./values.yaml#L34">kafka-ui.envs.config.KAFKA_CLUSTERS_0_NAME</a></td>
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
			<td id="kafka-ui--envs--config--KAFKA_CLUSTERS_0_READONLY"><a href="./values.yaml#L37">kafka-ui.envs.config.KAFKA_CLUSTERS_0_READONLY</a></td>
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
			<td id="kafka-ui--envs--config--KAFKA_CLUSTERS_0_ZOOKEEPER"><a href="./values.yaml#L36">kafka-ui.envs.config.KAFKA_CLUSTERS_0_ZOOKEEPER</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kafka-zookeeper-client:2181"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--envs--config--MANAGEMENT_HEALTH_LDAP_ENABLED"><a href="./values.yaml#L40">kafka-ui.envs.config.MANAGEMENT_HEALTH_LDAP_ENABLED</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"FALSE"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--envs--config--SERVER_SERVLET_CONTEXT_PATH"><a href="./values.yaml#L41">kafka-ui.envs.config.SERVER_SERVLET_CONTEXT_PATH</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/kafka-ui"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--envs--config--SPRING_SECURITY_USER_NAME"><a href="./values.yaml#L38">kafka-ui.envs.config.SPRING_SECURITY_USER_NAME</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redacted"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--envs--config--SPRING_SECURITY_USER_PASSWORD"><a href="./values.yaml#L39">kafka-ui.envs.config.SPRING_SECURITY_USER_PASSWORD</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redacted"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--envs--secret"><a href="./values.yaml#L32">kafka-ui.envs.secret</a></td>
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
			<td id="kafka-ui--fullnameOverride"><a href="./values.yaml#L25">kafka-ui.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kafka-ui"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--image--pullPolicy"><a href="./values.yaml#L29">kafka-ui.image.pullPolicy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Always"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--image--tag"><a href="./values.yaml#L28">kafka-ui.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"master"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L57">kafka-ui.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="kafka-ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-modsecurity""><a href="./values.yaml#L58">kafka-ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity"</a></td>
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
			<td id="kafka-ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L60">kafka-ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="kafka-ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L59">kafka-ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="kafka-ui--ingress--annotations--"nginx--ingress--kubernetes--io/rewrite-target""><a href="./values.yaml#L61">kafka-ui.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/kafka-ui/$2"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--ingress--annotations--"nginx--ingress--kubernetes--io/use-regex""><a href="./values.yaml#L62">kafka-ui.ingress.annotations."nginx.ingress.kubernetes.io/use-regex"</a></td>
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
			<td id="kafka-ui--ingress--enabled"><a href="./values.yaml#L54">kafka-ui.ingress.enabled</a></td>
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
			<td id="kafka-ui--ingress--host"><a href="./values.yaml#L64">kafka-ui.ingress.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--ingress--ingressClassName"><a href="./values.yaml#L55">kafka-ui.ingress.ingressClassName</a></td>
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
			<td id="kafka-ui--ingress--path"><a href="./values.yaml#L66">kafka-ui.ingress.path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/kafka-ui(/|$)(.*)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--ingress--tls--enabled"><a href="./values.yaml#L69">kafka-ui.ingress.tls.enabled</a></td>
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
			<td id="kafka-ui--ingress--tls--secretName"><a href="./values.yaml#L70">kafka-ui.ingress.tls.secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kafka-ui--probes--useHttpsScheme"><a href="./values.yaml#L44">kafka-ui.probes.useHttpsScheme</a></td>
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
			<td id="strimzi-kafka-operator--createAggregateRoles"><a href="./values.yaml#L13">strimzi-kafka-operator.createAggregateRoles</a></td>
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
			<td id="strimzi-kafka-operator--dashboards--annotations--grafana_dashboard_folder"><a href="./values.yaml#L20">strimzi-kafka-operator.dashboards.annotations.grafana_dashboard_folder</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Kafka"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="strimzi-kafka-operator--dashboards--enabled"><a href="./values.yaml#L16">strimzi-kafka-operator.dashboards.enabled</a></td>
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
			<td id="strimzi-kafka-operator--dashboards--label"><a href="./values.yaml#L17">strimzi-kafka-operator.dashboards.label</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana_dashboard"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="strimzi-kafka-operator--dashboards--labelValue"><a href="./values.yaml#L18">strimzi-kafka-operator.dashboards.labelValue</a></td>
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
			<td id="strimzi-kafka-operator--enabled"><a href="./values.yaml#L6">strimzi-kafka-operator.enabled</a></td>
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
			<td id="strimzi-kafka-operator--watchAnyNamespace"><a href="./values.yaml#L11">strimzi-kafka-operator.watchAnyNamespace</a></td>
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
			<td id="strimzi-kafka-operator--watchNamespaces[0]"><a href="./values.yaml#L9">strimzi-kafka-operator.watchNamespaces[0]</a></td>
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
			<td id="strimzi-kafka-operator--watchNamespaces[1]"><a href="./values.yaml#L10">strimzi-kafka-operator.watchNamespaces[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kafka"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
