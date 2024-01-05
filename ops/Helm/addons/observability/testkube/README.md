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
| https://kubeshop.github.io/helm-charts | testkube | 1.16.26 |

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
			<td id="testkube--mongodb--enabled"><a href="./values.yaml#L18">testkube.mongodb.enabled</a></td>
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
			<td id="testkube--nats--enabled"><a href="./values.yaml#L21">testkube.nats.enabled</a></td>
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
			<td id="testkube--nats--exporter--resources--limits--cpu"><a href="./values.yaml#L35">testkube.nats.exporter.resources.limits.cpu</a></td>
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
			<td id="testkube--nats--exporter--resources--limits--memory"><a href="./values.yaml#L36">testkube.nats.exporter.resources.limits.memory</a></td>
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
			<td id="testkube--nats--exporter--resources--requests--cpu"><a href="./values.yaml#L38">testkube.nats.exporter.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--nats--exporter--resources--requests--memory"><a href="./values.yaml#L39">testkube.nats.exporter.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--nats--nats--resources--limits--cpu"><a href="./values.yaml#L26">testkube.nats.nats.resources.limits.cpu</a></td>
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
			<td id="testkube--nats--nats--resources--limits--memory"><a href="./values.yaml#L27">testkube.nats.nats.resources.limits.memory</a></td>
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
			<td id="testkube--nats--nats--resources--requests--cpu"><a href="./values.yaml#L29">testkube.nats.nats.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--nats--nats--resources--requests--memory"><a href="./values.yaml#L30">testkube.nats.nats.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--preUpgradeHook--resources--limits--cpu"><a href="./values.yaml#L11">testkube.preUpgradeHook.resources.limits.cpu</a></td>
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
			<td id="testkube--preUpgradeHook--resources--limits--memory"><a href="./values.yaml#L12">testkube.preUpgradeHook.resources.limits.memory</a></td>
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
			<td id="testkube--preUpgradeHook--resources--requests--cpu"><a href="./values.yaml#L14">testkube.preUpgradeHook.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--preUpgradeHook--resources--requests--memory"><a href="./values.yaml#L15">testkube.preUpgradeHook.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--minio--enabled"><a href="./values.yaml#L49">testkube.testkube-api.minio.enabled</a></td>
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
			<td id="testkube--testkube-api--mongodb--secretKey"><a href="./values.yaml#L53">testkube.testkube-api.mongodb.secretKey</a></td>
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
			<td id="testkube--testkube-api--mongodb--secretName"><a href="./values.yaml#L52">testkube.testkube-api.mongodb.secretName</a></td>
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
			<td id="testkube--testkube-api--multinamespace--enabled"><a href="./values.yaml#L46">testkube.testkube-api.multinamespace.enabled</a></td>
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
			<td id="testkube--testkube-api--nats--enabled"><a href="./values.yaml#L43">testkube.testkube-api.nats.enabled</a></td>
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
			<td id="testkube--testkube-api--prometheus--enabled"><a href="./values.yaml#L78">testkube.testkube-api.prometheus.enabled</a></td>
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
			<td id="testkube--testkube-api--prometheus--monitoringLabels--release"><a href="./values.yaml#L80">testkube.testkube-api.prometheus.monitoringLabels.release</a></td>
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
			<td id="testkube--testkube-api--resources--limits--cpu"><a href="./values.yaml#L84">testkube.testkube-api.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"500m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--resources--limits--memory"><a href="./values.yaml#L85">testkube.testkube-api.resources.limits.memory</a></td>
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
			<td id="testkube--testkube-api--resources--requests--cpu"><a href="./values.yaml#L87">testkube.testkube-api.resources.requests.cpu</a></td>
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
			<td id="testkube--testkube-api--resources--requests--memory"><a href="./values.yaml#L88">testkube.testkube-api.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"200Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--storage--SSL"><a href="./values.yaml#L74">testkube.testkube-api.storage.SSL</a></td>
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
			<td id="testkube--testkube-api--storage--expiration"><a href="./values.yaml#L73">testkube.testkube-api.storage.expiration</a></td>
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
			<td id="testkube--testkube-api--storage--region"><a href="./values.yaml#L71">testkube.testkube-api.storage.region</a></td>
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
			<td id="testkube--testkube-api--storage--scrapperEnabled"><a href="./values.yaml#L75">testkube.testkube-api.storage.scrapperEnabled</a></td>
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
			<td id="testkube--testkube-api--storage--token"><a href="./values.yaml#L72">testkube.testkube-api.storage.token</a></td>
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
			<td id="testkube--testkube-api--testConnection--resources--limits--cpu"><a href="./values.yaml#L93">testkube.testkube-api.testConnection.resources.limits.cpu</a></td>
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
			<td id="testkube--testkube-api--testConnection--resources--limits--memory"><a href="./values.yaml#L94">testkube.testkube-api.testConnection.resources.limits.memory</a></td>
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
			<td id="testkube--testkube-api--testConnection--resources--requests--cpu"><a href="./values.yaml#L96">testkube.testkube-api.testConnection.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--testConnection--resources--requests--memory"><a href="./values.yaml#L97">testkube.testkube-api.testConnection.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-api--uiIngress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L59">testkube.testkube-api.uiIngress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L61">testkube.testkube-api.uiIngress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L60">testkube.testkube-api.uiIngress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="testkube--testkube-api--uiIngress--className"><a href="./values.yaml#L57">testkube.testkube-api.uiIngress.className</a></td>
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
			<td id="testkube--testkube-api--uiIngress--enabled"><a href="./values.yaml#L56">testkube.testkube-api.uiIngress.enabled</a></td>
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
			<td id="testkube--testkube-api--uiIngress--hosts[0]"><a href="./values.yaml#L63">testkube.testkube-api.uiIngress.hosts[0]</a></td>
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
			<td id="testkube--testkube-api--uiIngress--tls[0]--hosts[0]"><a href="./values.yaml#L67">testkube.testkube-api.uiIngress.tls[0].hosts[0]</a></td>
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
			<td id="testkube--testkube-api--uiIngress--tls[0]--secretName"><a href="./values.yaml#L68">testkube.testkube-api.uiIngress.tls[0].secretName</a></td>
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
			<td id="testkube--testkube-api--uiIngress--tlsenabled"><a href="./values.yaml#L64">testkube.testkube-api.uiIngress.tlsenabled</a></td>
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
			<td id="testkube--testkube-dashboard--apiServerEndpoint"><a href="./values.yaml#L100">testkube.testkube-dashboard.apiServerEndpoint</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L106">testkube.testkube-dashboard.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L108">testkube.testkube-dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L107">testkube.testkube-dashboard.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--enabled"><a href="./values.yaml#L103">testkube.testkube-dashboard.ingress.enabled</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--hosts[0]"><a href="./values.yaml#L111">testkube.testkube-dashboard.ingress.hosts[0]</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L116">testkube.testkube-dashboard.ingress.tls[0].hosts[0]</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--tls[0]--secretName"><a href="./values.yaml#L117">testkube.testkube-dashboard.ingress.tls[0].secretName</a></td>
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
			<td id="testkube--testkube-dashboard--ingress--tlsenabled"><a href="./values.yaml#L113">testkube.testkube-dashboard.ingress.tlsenabled</a></td>
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
			<td id="testkube--testkube-dashboard--oauth2--enabled"><a href="./values.yaml#L120">testkube.testkube-dashboard.oauth2.enabled</a></td>
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
			<td id="testkube--testkube-dashboard--resources--limits--cpu"><a href="./values.yaml#L124">testkube.testkube-dashboard.resources.limits.cpu</a></td>
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
			<td id="testkube--testkube-dashboard--resources--limits--memory"><a href="./values.yaml#L125">testkube.testkube-dashboard.resources.limits.memory</a></td>
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
			<td id="testkube--testkube-dashboard--resources--requests--cpu"><a href="./values.yaml#L127">testkube.testkube-dashboard.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--resources--requests--memory"><a href="./values.yaml#L128">testkube.testkube-dashboard.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--testConnection--resources--limits--cpu"><a href="./values.yaml#L133">testkube.testkube-dashboard.testConnection.resources.limits.cpu</a></td>
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
			<td id="testkube--testkube-dashboard--testConnection--resources--limits--memory"><a href="./values.yaml#L134">testkube.testkube-dashboard.testConnection.resources.limits.memory</a></td>
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
			<td id="testkube--testkube-dashboard--testConnection--resources--requests--cpu"><a href="./values.yaml#L136">testkube.testkube-dashboard.testConnection.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-dashboard--testConnection--resources--requests--memory"><a href="./values.yaml#L137">testkube.testkube-dashboard.testConnection.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-operator--enabled"><a href="./values.yaml#L140">testkube.testkube-operator.enabled</a></td>
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
			<td id="testkube--testkube-operator--proxy--resources--limits--cpu"><a href="./values.yaml#L155">testkube.testkube-operator.proxy.resources.limits.cpu</a></td>
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
			<td id="testkube--testkube-operator--proxy--resources--limits--memory"><a href="./values.yaml#L156">testkube.testkube-operator.proxy.resources.limits.memory</a></td>
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
			<td id="testkube--testkube-operator--proxy--resources--requests--cpu"><a href="./values.yaml#L158">testkube.testkube-operator.proxy.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-operator--proxy--resources--requests--memory"><a href="./values.yaml#L159">testkube.testkube-operator.proxy.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-operator--resources--limits--cpu"><a href="./values.yaml#L144">testkube.testkube-operator.resources.limits.cpu</a></td>
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
			<td id="testkube--testkube-operator--resources--limits--memory"><a href="./values.yaml#L145">testkube.testkube-operator.resources.limits.memory</a></td>
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
			<td id="testkube--testkube-operator--resources--requests--cpu"><a href="./values.yaml#L147">testkube.testkube-operator.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"20m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-operator--resources--requests--memory"><a href="./values.yaml#L148">testkube.testkube-operator.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"56Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="testkube--testkube-operator--webhook--enabled"><a href="./values.yaml#L162">testkube.testkube-operator.webhook.enabled</a></td>
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
			<td id="testkube--testkube-operator--webhook--patch--enabled"><a href="./values.yaml#L165">testkube.testkube-operator.webhook.patch.enabled</a></td>
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
