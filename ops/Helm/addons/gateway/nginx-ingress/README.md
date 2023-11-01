# nginx-ingress

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kubernetes.github.io/ingress-nginx | ingress-nginx | 4.8.3 |

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
			<td id="ingress-nginx--controller--admissionWebhooks--enabled"><a href="./values.yaml#L87">ingress-nginx.controller.admissionWebhooks.enabled</a></td>
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
			<td id="ingress-nginx--controller--config--enable-opentelemetry"><a href="./values.yaml#L43">ingress-nginx.controller.config.enable-opentelemetry</a></td>
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
			<td id="ingress-nginx--controller--config--opentelemetry-config"><a href="./values.yaml#L44">ingress-nginx.controller.config.opentelemetry-config</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/etc/nginx/opentelemetry.toml"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--opentelemetry-operation-name"><a href="./values.yaml#L45">ingress-nginx.controller.config.opentelemetry-operation-name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HTTP $request_method $service_name $uri"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--opentelemetry-trust-incoming-span"><a href="./values.yaml#L46">ingress-nginx.controller.config.opentelemetry-trust-incoming-span</a></td>
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
			<td id="ingress-nginx--controller--config--otel-max-export-batch-size"><a href="./values.yaml#L51">ingress-nginx.controller.config.otel-max-export-batch-size</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"512"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otel-max-queuesize"><a href="./values.yaml#L49">ingress-nginx.controller.config.otel-max-queuesize</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"2048"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otel-sampler"><a href="./values.yaml#L53">ingress-nginx.controller.config.otel-sampler</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"AlwaysOn"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otel-sampler-parent-based"><a href="./values.yaml#L55">ingress-nginx.controller.config.otel-sampler-parent-based</a></td>
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
			<td id="ingress-nginx--controller--config--otel-sampler-ratio"><a href="./values.yaml#L54">ingress-nginx.controller.config.otel-sampler-ratio</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1.0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otel-schedule-delay-millis"><a href="./values.yaml#L50">ingress-nginx.controller.config.otel-schedule-delay-millis</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"5000"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otel-service-name"><a href="./values.yaml#L52">ingress-nginx.controller.config.otel-service-name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"nginx-ingress"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otlp-collector-host"><a href="./values.yaml#L47">ingress-nginx.controller.config.otlp-collector-host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--otlp-collector-port"><a href="./values.yaml#L48">ingress-nginx.controller.config.otlp-collector-port</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"4317"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--config--server-snippet"><a href="./values.yaml#L56">ingress-nginx.controller.config.server-snippet</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"opentelemetry_attribute \"ingress.namespace\" \"$namespace\";\nopentelemetry_attribute \"ingress.service_name\" \"$service_name\";\nopentelemetry_attribute \"ingress.name\" \"$ingress_name\";\nopentelemetry_attribute \"ingress.upstream\" \"$proxy_upstream_name\";\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--extraEnvs[0]--name"><a href="./values.yaml#L36">ingress-nginx.controller.extraEnvs[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"NODE_IP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--extraEnvs[0]--valueFrom--fieldRef--fieldPath"><a href="./values.yaml#L39">ingress-nginx.controller.extraEnvs[0].valueFrom.fieldRef.fieldPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"status.hostIP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--hostNetwork"><a href="./values.yaml#L65">ingress-nginx.controller.hostNetwork</a></td>
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
			<td id="ingress-nginx--controller--ingressClassResource--default"><a href="./values.yaml#L71">ingress-nginx.controller.ingressClassResource.default</a></td>
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
			<td id="ingress-nginx--controller--kind"><a href="./values.yaml#L75">ingress-nginx.controller.kind</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Deployment"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--enabled"><a href="./values.yaml#L91">ingress-nginx.controller.metrics.enabled</a></td>
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
			<td id="ingress-nginx--controller--metrics--prometheusRule--additionalLabels--app"><a href="./values.yaml#L104">ingress-nginx.controller.metrics.prometheusRule.additionalLabels.app</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kube-prometheus-stack"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--additionalLabels--release"><a href="./values.yaml#L105">ingress-nginx.controller.metrics.prometheusRule.additionalLabels.release</a></td>
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
			<td id="ingress-nginx--controller--metrics--prometheusRule--enabled"><a href="./values.yaml#L102">ingress-nginx.controller.metrics.prometheusRule.enabled</a></td>
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
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[0]--alert"><a href="./values.yaml#L108">ingress-nginx.controller.metrics.prometheusRule.rules[0].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"NGINXConfigFailed"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[0]--annotations--description"><a href="./values.yaml#L114">ingress-nginx.controller.metrics.prometheusRule.rules[0].annotations.description</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"bad ingress config - nginx config test failed"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[0]--annotations--summary"><a href="./values.yaml#L115">ingress-nginx.controller.metrics.prometheusRule.rules[0].annotations.summary</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"uninstall the latest ingress changes to allow config reloads to resume"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[0]--expr"><a href="./values.yaml#L109">ingress-nginx.controller.metrics.prometheusRule.rules[0].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"count(nginx_ingress_controller_config_last_reload_successful == 0) \u003e 0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[0]--for"><a href="./values.yaml#L110">ingress-nginx.controller.metrics.prometheusRule.rules[0].for</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1s"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[0]--labels--severity"><a href="./values.yaml#L112">ingress-nginx.controller.metrics.prometheusRule.rules[0].labels.severity</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"critical"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[1]--alert"><a href="./values.yaml#L116">ingress-nginx.controller.metrics.prometheusRule.rules[1].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"NGINXCertificateExpiry"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[1]--annotations--description"><a href="./values.yaml#L122">ingress-nginx.controller.metrics.prometheusRule.rules[1].annotations.description</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ssl certificate(s) will expire in less then a week"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[1]--annotations--summary"><a href="./values.yaml#L123">ingress-nginx.controller.metrics.prometheusRule.rules[1].annotations.summary</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"renew expiring certificates to avoid downtime"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[1]--expr"><a href="./values.yaml#L117">ingress-nginx.controller.metrics.prometheusRule.rules[1].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"(avg(nginx_ingress_controller_ssl_expire_time_seconds) by (host) - time()) \u003c 604800"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[1]--for"><a href="./values.yaml#L118">ingress-nginx.controller.metrics.prometheusRule.rules[1].for</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1s"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[1]--labels--severity"><a href="./values.yaml#L120">ingress-nginx.controller.metrics.prometheusRule.rules[1].labels.severity</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"critical"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[2]--alert"><a href="./values.yaml#L124">ingress-nginx.controller.metrics.prometheusRule.rules[2].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"NGINXTooMany500s"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[2]--annotations--description"><a href="./values.yaml#L130">ingress-nginx.controller.metrics.prometheusRule.rules[2].annotations.description</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Too many 5XXs"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[2]--annotations--summary"><a href="./values.yaml#L131">ingress-nginx.controller.metrics.prometheusRule.rules[2].annotations.summary</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"More than 5% of all requests returned 5XX, this requires your attention"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[2]--expr"><a href="./values.yaml#L125">ingress-nginx.controller.metrics.prometheusRule.rules[2].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"100 * ( sum( nginx_ingress_controller_requests{status=~\"5.+\"} ) / sum(nginx_ingress_controller_requests) ) \u003e 5"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[2]--for"><a href="./values.yaml#L126">ingress-nginx.controller.metrics.prometheusRule.rules[2].for</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[2]--labels--severity"><a href="./values.yaml#L128">ingress-nginx.controller.metrics.prometheusRule.rules[2].labels.severity</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"warning"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[3]--alert"><a href="./values.yaml#L132">ingress-nginx.controller.metrics.prometheusRule.rules[3].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"NGINXTooMany400s"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[3]--annotations--description"><a href="./values.yaml#L138">ingress-nginx.controller.metrics.prometheusRule.rules[3].annotations.description</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Too many 4XXs"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[3]--annotations--summary"><a href="./values.yaml#L139">ingress-nginx.controller.metrics.prometheusRule.rules[3].annotations.summary</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"More than 5% of all requests returned 4XX, this requires your attention"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[3]--expr"><a href="./values.yaml#L133">ingress-nginx.controller.metrics.prometheusRule.rules[3].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"100 * ( sum( nginx_ingress_controller_requests{status=~\"4.+\"} ) / sum(nginx_ingress_controller_requests) ) \u003e 5"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[3]--for"><a href="./values.yaml#L134">ingress-nginx.controller.metrics.prometheusRule.rules[3].for</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--prometheusRule--rules[3]--labels--severity"><a href="./values.yaml#L136">ingress-nginx.controller.metrics.prometheusRule.rules[3].labels.severity</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"warning"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--metrics--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L96">ingress-nginx.controller.metrics.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="ingress-nginx--controller--metrics--serviceMonitor--enabled"><a href="./values.yaml#L94">ingress-nginx.controller.metrics.serviceMonitor.enabled</a></td>
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
			<td id="ingress-nginx--controller--metrics--serviceMonitor--namespaceSelector--matchNames[0]"><a href="./values.yaml#L99">ingress-nginx.controller.metrics.serviceMonitor.namespaceSelector.matchNames[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"nginx-ingress"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--nodeSelector--"kubernetes--io/hostname""><a href="./values.yaml#L25">ingress-nginx.controller.nodeSelector."kubernetes.io/hostname"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"node1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--opentelemetry--enabled"><a href="./values.yaml#L68">ingress-nginx.controller.opentelemetry.enabled</a></td>
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
			<td id="ingress-nginx--controller--podSecurityContext--fsGroup"><a href="./values.yaml#L78">ingress-nginx.controller.podSecurityContext.fsGroup</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1001
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--resources--requests--cpu"><a href="./values.yaml#L32">ingress-nginx.controller.resources.requests.cpu</a></td>
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
			<td id="ingress-nginx--controller--resources--requests--memory"><a href="./values.yaml#L33">ingress-nginx.controller.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"90Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--service--nodePorts--http"><a href="./values.yaml#L83">ingress-nginx.controller.service.nodePorts.http</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
80
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--service--nodePorts--https"><a href="./values.yaml#L84">ingress-nginx.controller.service.nodePorts.https</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
443
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--controller--service--type"><a href="./values.yaml#L81">ingress-nginx.controller.service.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"NodePort"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--defaultBackend--enabled"><a href="./values.yaml#L10">ingress-nginx.defaultBackend.enabled</a></td>
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
			<td id="ingress-nginx--defaultBackend--resources--limits--cpu"><a href="./values.yaml#L14">ingress-nginx.defaultBackend.resources.limits.cpu</a></td>
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
			<td id="ingress-nginx--defaultBackend--resources--limits--memory"><a href="./values.yaml#L15">ingress-nginx.defaultBackend.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--defaultBackend--resources--requests--cpu"><a href="./values.yaml#L17">ingress-nginx.defaultBackend.resources.requests.cpu</a></td>
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
			<td id="ingress-nginx--defaultBackend--resources--requests--memory"><a href="./values.yaml#L18">ingress-nginx.defaultBackend.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress-nginx--enabled"><a href="./values.yaml#L6">ingress-nginx.enabled</a></td>
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
