# testkube

![Version: 1.13.27](https://img.shields.io/badge/Version-1.13.27-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kubeshop.github.io/helm-charts | testkube | 1.15.28 |

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
			<td id="testkube--enabled"><a href="./values.yaml#L6">testkube.enabled</a></td>
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
			<td id="testkube--mongodb--enabled"><a href="./values.yaml#L9">testkube.mongodb.enabled</a></td>
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
			<td id="testkube--nats--enabled"><a href="./values.yaml#L12">testkube.nats.enabled</a></td>
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
			<td id="testkube--testkube-api--minio--enabled"><a href="./values.yaml#L22">testkube.testkube-api.minio.enabled</a></td>
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
			<td id="testkube--testkube-api--mongodb--secretKey"><a href="./values.yaml#L26">testkube.testkube-api.mongodb.secretKey</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"connectionString.standard"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--mongodb--secretName"><a href="./values.yaml#L25">testkube.testkube-api.mongodb.secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"mongodb-testkube-testkube"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--multinamespace--enabled"><a href="./values.yaml#L19">testkube.testkube-api.multinamespace.enabled</a></td>
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
			<td id="testkube--testkube-api--nats--enabled"><a href="./values.yaml#L16">testkube.testkube-api.nats.enabled</a></td>
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
			<td id="testkube--testkube-api--prometheus--enabled"><a href="./values.yaml#L53">testkube.testkube-api.prometheus.enabled</a></td>
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
			<td id="testkube--testkube-api--prometheus--monitoringLabels--release"><a href="./values.yaml#L55">testkube.testkube-api.prometheus.monitoringLabels.release</a></td>
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
			<td id="testkube--testkube-api--storage--SSL"><a href="./values.yaml#L49">testkube.testkube-api.storage.SSL</a></td>
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
			<td id="testkube--testkube-api--storage--expiration"><a href="./values.yaml#L48">testkube.testkube-api.storage.expiration</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
0
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--storage--region"><a href="./values.yaml#L46">testkube.testkube-api.storage.region</a></td>
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
			<td id="testkube--testkube-api--storage--scrapperEnabled"><a href="./values.yaml#L50">testkube.testkube-api.storage.scrapperEnabled</a></td>
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
			<td id="testkube--testkube-api--storage--token"><a href="./values.yaml#L47">testkube.testkube-api.storage.token</a></td>
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
			<td id="testkube--testkube-api--uiIngress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L32">testkube.testkube-api.uiIngress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--annotations--"nginx--ingress--kubernetes--io/enable-modsecurity""><a href="./values.yaml#L33">testkube.testkube-api.uiIngress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L35">testkube.testkube-api.uiIngress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L34">testkube.testkube-api.uiIngress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--className"><a href="./values.yaml#L30">testkube.testkube-api.uiIngress.className</a></td>
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
			<td id="testkube--testkube-api--uiIngress--enabled"><a href="./values.yaml#L29">testkube.testkube-api.uiIngress.enabled</a></td>
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
			<td id="testkube--testkube-api--uiIngress--hosts[0]"><a href="./values.yaml#L37">testkube.testkube-api.uiIngress.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"testkube.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--uiIngress--path"><a href="./values.yaml#L38">testkube.testkube-api.uiIngress.path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/results"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--uiIngress--tls[0]--hosts[0]"><a href="./values.yaml#L42">testkube.testkube-api.uiIngress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"testkube.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--uiIngress--tls[0]--secretName"><a href="./values.yaml#L43">testkube.testkube-api.uiIngress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"testkube-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--uiIngress--tlsenabled"><a href="./values.yaml#L39">testkube.testkube-api.uiIngress.tlsenabled</a></td>
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
			<td id="testkube--testkube-dashboard--apiServerEndpoint"><a href="./values.yaml#L58">testkube.testkube-dashboard.apiServerEndpoint</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://testkube.shortlink.best/results/v1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L64">testkube.testkube-dashboard.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-modsecurity""><a href="./values.yaml#L65">testkube.testkube-dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-modsecurity"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L67">testkube.testkube-dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L66">testkube.testkube-dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--enabled"><a href="./values.yaml#L61">testkube.testkube-dashboard.ingress.enabled</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--hosts[0]"><a href="./values.yaml#L70">testkube.testkube-dashboard.ingress.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"testkube.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L75">testkube.testkube-dashboard.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"testkube.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--ingress--tls[0]--secretName"><a href="./values.yaml#L76">testkube.testkube-dashboard.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"testkube-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--ingress--tlsenabled"><a href="./values.yaml#L72">testkube.testkube-dashboard.ingress.tlsenabled</a></td>
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
			<td id="testkube--testkube-dashboard--oauth2--enabled"><a href="./values.yaml#L79">testkube.testkube-dashboard.oauth2.enabled</a></td>
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
			<td id="testkube--testkube-operator"><a href="./values.yaml#L81">testkube.testkube-operator</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
