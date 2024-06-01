# argo-workflows

![Version: 0.3.18](https://img.shields.io/badge/Version-0.3.18-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.9.0](https://img.shields.io/badge/AppVersion-2.9.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-workflows | 0.41.7 |

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
			<td id="argo-workflows--controller--logging"><a href="./values.yaml#L25">argo-workflows.controller.logging</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "format": "json"
}
</pre>
</div>
			</td>
			<td>enable persistence using postgres postgresql:  host: localhost  port: 5432  database: argo  tableName: argo_workflows</td>
		</tr>
		<tr>
			<td id="argo-workflows--controller--metricsConfig--enabled"><a href="./values.yaml#L8">argo-workflows.controller.metricsConfig.enabled</a></td>
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
			<td id="argo-workflows--controller--resources--limits--cpu"><a href="./values.yaml#L30">argo-workflows.controller.resources.limits.cpu</a></td>
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
			<td id="argo-workflows--controller--resources--limits--memory"><a href="./values.yaml#L31">argo-workflows.controller.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1024Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--controller--resources--requests--cpu"><a href="./values.yaml#L33">argo-workflows.controller.resources.requests.cpu</a></td>
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
			<td id="argo-workflows--controller--resources--requests--memory"><a href="./values.yaml#L34">argo-workflows.controller.resources.requests.memory</a></td>
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
			<td id="argo-workflows--controller--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L14">argo-workflows.controller.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="argo-workflows--controller--serviceMonitor--enabled"><a href="./values.yaml#L12">argo-workflows.controller.serviceMonitor.enabled</a></td>
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
			<td id="argo-workflows--controller--telemetryConfig--enabled"><a href="./values.yaml#L10">argo-workflows.controller.telemetryConfig.enabled</a></td>
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
			<td id="argo-workflows--controller--workflowNamespaces"><a href="./values.yaml#L16">argo-workflows.controller.workflowNamespaces</a></td>
			<td>
list
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
[]
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--enabled"><a href="./values.yaml#L2">argo-workflows.enabled</a></td>
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
			<td id="argo-workflows--fullnameOverride"><a href="./values.yaml#L4">argo-workflows.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo-workflows"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--authModes[0]"><a href="./values.yaml#L71">argo-workflows.server.authModes[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"server"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--baseHref"><a href="./values.yaml#L39">argo-workflows.server.baseHref</a></td>
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
			<td id="argo-workflows--server--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L49">argo-workflows.server.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="argo-workflows--server--ingress--annotations--"nginx--ingress--kubernetes--io/backend-protocol""><a href="./values.yaml#L50">argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HTTP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L52">argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="argo-workflows--server--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L51">argo-workflows.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="argo-workflows--server--ingress--enabled"><a href="./values.yaml#L44">argo-workflows.server.ingress.enabled</a></td>
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
			<td id="argo-workflows--server--ingress--hosts[0]"><a href="./values.yaml#L55">argo-workflows.server.ingress.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"workflows.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--ingress--ingressClassName"><a href="./values.yaml#L46">argo-workflows.server.ingress.ingressClassName</a></td>
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
			<td id="argo-workflows--server--ingress--paths[0]"><a href="./values.yaml#L58">argo-workflows.server.ingress.paths[0]</a></td>
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
			<td id="argo-workflows--server--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L63">argo-workflows.server.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"workflows.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--ingress--tls[0]--secretName"><a href="./values.yaml#L61">argo-workflows.server.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"workflows-argo-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--logging--format"><a href="./values.yaml#L66">argo-workflows.server.logging.format</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"json"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--revisionHistoryLimit"><a href="./values.yaml#L37">argo-workflows.server.revisionHistoryLimit</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--secure"><a href="./values.yaml#L68">argo-workflows.server.secure</a></td>
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
			<td id="argo-workflows--server--sso--clientId--key"><a href="./values.yaml#L86">argo-workflows.server.sso.clientId.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"client-id"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--clientId--name"><a href="./values.yaml#L85">argo-workflows.server.sso.clientId.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo-workflows-sso"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--clientSecret--key"><a href="./values.yaml#L89">argo-workflows.server.sso.clientSecret.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"client-secret"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--clientSecret--name"><a href="./values.yaml#L88">argo-workflows.server.sso.clientSecret.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo-workflows-sso"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--enabled"><a href="./values.yaml#L77">argo-workflows.server.sso.enabled</a></td>
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
			<td id="argo-workflows--server--sso--issuer"><a href="./values.yaml#L78">argo-workflows.server.sso.issuer</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://argo.shortlink.best/api/dex"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--rbac--enabled"><a href="./values.yaml#L81">argo-workflows.server.sso.rbac.enabled</a></td>
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
			<td id="argo-workflows--server--sso--redirectUrl"><a href="./values.yaml#L90">argo-workflows.server.sso.redirectUrl</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://workflows.shortlink.best/oauth2/callback"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--scopes[0]"><a href="./values.yaml#L83">argo-workflows.server.sso.scopes[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"groups"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-workflows--server--sso--sessionExpiry"><a href="./values.yaml#L79">argo-workflows.server.sso.sessionExpiry</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"240h"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
