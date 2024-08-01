# istio

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kiali.org/helm-charts | kiali-operator | 1.88.0 |

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
			<td id="kiali-operator--cr--create"><a href="./values.yaml#L7">kiali-operator.cr.create</a></td>
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
			<td id="kiali-operator--cr--namespace"><a href="./values.yaml#L8">kiali-operator.cr.namespace</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
""
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--auth--strategy"><a href="./values.yaml#L16">kiali-operator.cr.spec.auth.strategy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"anonymous"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--deployment--accessible_namespaces[0]"><a href="./values.yaml#L18">kiali-operator.cr.spec.deployment.accessible_namespaces[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"**"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--deployment--view_only_mode"><a href="./values.yaml#L19">kiali-operator.cr.spec.deployment.view_only_mode</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--custom_dashboards--enabled"><a href="./values.yaml#L39">kiali-operator.cr.spec.external_services.custom_dashboards.enabled</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--custom_dashboards--folder"><a href="./values.yaml#L42">kiali-operator.cr.spec.external_services.custom_dashboards.folder</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Kiali"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--custom_dashboards--label_selector"><a href="./values.yaml#L40">kiali-operator.cr.spec.external_services.custom_dashboards.label_selector</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"app=grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--custom_dashboards--namespace"><a href="./values.yaml#L41">kiali-operator.cr.spec.external_services.custom_dashboards.namespace</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--grafana--enabled"><a href="./values.yaml#L32">kiali-operator.cr.spec.external_services.grafana.enabled</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--grafana--in_cluster_url"><a href="./values.yaml#L33">kiali-operator.cr.spec.external_services.grafana.in_cluster_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana.grafana:80"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--grafana--url"><a href="./values.yaml#L35">kiali-operator.cr.spec.external_services.grafana.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/grafana"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[0]--app_label"><a href="./values.yaml#L51">kiali-operator.cr.spec.external_services.istio.component_status.components[0].app_label</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istiod"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[0]--is_core"><a href="./values.yaml#L52">kiali-operator.cr.spec.external_services.istio.component_status.components[0].is_core</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[0]--is_proxy"><a href="./values.yaml#L53">kiali-operator.cr.spec.external_services.istio.component_status.components[0].is_proxy</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[1]--app_label"><a href="./values.yaml#L54">kiali-operator.cr.spec.external_services.istio.component_status.components[1].app_label</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istio-ingress"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[1]--is_core"><a href="./values.yaml#L55">kiali-operator.cr.spec.external_services.istio.component_status.components[1].is_core</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[1]--is_proxy"><a href="./values.yaml#L56">kiali-operator.cr.spec.external_services.istio.component_status.components[1].is_proxy</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--components[1]--namespace"><a href="./values.yaml#L57">kiali-operator.cr.spec.external_services.istio.component_status.components[1].namespace</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istio-ingress"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--component_status--enabled"><a href="./values.yaml#L49">kiali-operator.cr.spec.external_services.istio.component_status.enabled</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--istio--config_map_name"><a href="./values.yaml#L44">kiali-operator.cr.spec.external_services.istio.config_map_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istio"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--istio_sidecar_injector_config_map_name"><a href="./values.yaml#L46">kiali-operator.cr.spec.external_services.istio.istio_sidecar_injector_config_map_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istio-sidecar-injector"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--istiod_deployment_name"><a href="./values.yaml#L45">kiali-operator.cr.spec.external_services.istio.istiod_deployment_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istiod"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--istio--root_namespace"><a href="./values.yaml#L47">kiali-operator.cr.spec.external_services.istio.root_namespace</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istio-system"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--prometheus--url"><a href="./values.yaml#L37">kiali-operator.cr.spec.external_services.prometheus.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://prometheus-prometheus.prometheus-operator:9090/prometheus"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--tracing--auth--type"><a href="./values.yaml#L26">kiali-operator.cr.spec.external_services.tracing.auth.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"none"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--tracing--enabled"><a href="./values.yaml#L24">kiali-operator.cr.spec.external_services.tracing.enabled</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--tracing--in_cluster_url"><a href="./values.yaml#L28">kiali-operator.cr.spec.external_services.tracing.in_cluster_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-tempo.grafana:16686"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--tracing--namespace_selector"><a href="./values.yaml#L27">kiali-operator.cr.spec.external_services.tracing.namespace_selector</a></td>
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
			<td id="kiali-operator--cr--spec--external_services--tracing--url"><a href="./values.yaml#L29">kiali-operator.cr.spec.external_services.tracing.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-tempo.grafana:16686/"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--external_services--tracing--use_grpc"><a href="./values.yaml#L30">kiali-operator.cr.spec.external_services.tracing.use_grpc</a></td>
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
			<td id="kiali-operator--cr--spec--istio_labels--app_label_name"><a href="./values.yaml#L13">kiali-operator.cr.spec.istio_labels.app_label_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"app.kubernetes.io/name"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--istio_labels--version_label_name"><a href="./values.yaml#L14">kiali-operator.cr.spec.istio_labels.version_label_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"app.kubernetes.io/version"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--istio_namespace"><a href="./values.yaml#L11">kiali-operator.cr.spec.istio_namespace</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"istio-system"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--cr--spec--server--web_root"><a href="./values.yaml#L21">kiali-operator.cr.spec.server.web_root</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/kiali"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kiali-operator--enabled"><a href="./values.yaml#L2">kiali-operator.enabled</a></td>
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
			<td id="kiali-operator--onlyViewOnlyMode"><a href="./values.yaml#L4">kiali-operator.onlyViewOnlyMode</a></td>
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
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
