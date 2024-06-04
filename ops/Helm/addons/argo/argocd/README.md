# argo

![Version: 0.4.6](https://img.shields.io/badge/Version-0.4.6-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.11.0](https://img.shields.io/badge/AppVersion-2.11.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-cd | 7.1.1 |
| https://argoproj.github.io/argo-helm | argocd-apps | 2.0.0 |
| https://argoproj.github.io/argo-helm | argocd-image-updater | 0.10.1 |
| oci://registry-1.docker.io/bitnamicharts | redis | 19.5.1 |

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
			<td id="argo-cd--applicationSet--metrics--enabled"><a href="./values.yaml#L282">argo-cd.applicationSet.metrics.enabled</a></td>
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
			<td id="argo-cd--applicationSet--metrics--serviceMonitor--enabled"><a href="./values.yaml#L284">argo-cd.applicationSet.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--applicationSet--resources--limits--cpu"><a href="./values.yaml#L275">argo-cd.applicationSet.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--applicationSet--resources--limits--memory"><a href="./values.yaml#L276">argo-cd.applicationSet.resources.limits.memory</a></td>
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
			<td id="argo-cd--applicationSet--resources--requests--cpu"><a href="./values.yaml#L278">argo-cd.applicationSet.resources.requests.cpu</a></td>
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
			<td id="argo-cd--applicationSet--resources--requests--memory"><a href="./values.yaml#L279">argo-cd.applicationSet.resources.requests.memory</a></td>
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
			<td id="argo-cd--configs--cm--"controller--diff--server--side""><a href="./values.yaml#L292">argo-cd.configs.cm."controller.diff.server.side"</a></td>
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
			<td id="argo-cd--configs--cm--"exec--enabled""><a href="./values.yaml#L294">argo-cd.configs.cm."exec.enabled"</a></td>
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
			<td id="argo-cd--configs--cm--"helm--valuesFileSchemes""><a href="./values.yaml#L298">argo-cd.configs.cm."helm.valuesFileSchemes"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"secrets+gpg-import, secrets+gpg-import-kubernetes, secrets+age-import, secrets+age-import-kubernetes, secrets,secrets+literal, https"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--configs--cm--"resource--compareoptions""><a href="./values.yaml#L305">argo-cd.configs.cm."resource.compareoptions"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"# disables status field diffing in specified resource types\n# ignoreAggregatedRoles: true\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--configs--cm--"resource--exclusions""><a href="./values.yaml#L309">argo-cd.configs.cm."resource.exclusions"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"- apiGroups:\n    - cilium.io\n  kinds:\n    - CiliumIdentity\n  clusters:\n    - \"*\"\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--configs--cm--"server--enable--proxy--extension""><a href="./values.yaml#L296">argo-cd.configs.cm."server.enable.proxy.extension"</a></td>
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
			<td id="argo-cd--configs--cm--"statusbadge--enabled""><a href="./values.yaml#L290">argo-cd.configs.cm."statusbadge.enabled"</a></td>
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
			<td id="argo-cd--configs--cm--url"><a href="./values.yaml#L288">argo-cd.configs.cm.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://argo.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--configs--params--"dexserver--disable--tls""><a href="./values.yaml#L326">argo-cd.configs.params."dexserver.disable.tls"</a></td>
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
			<td id="argo-cd--configs--params--"server--insecure""><a href="./values.yaml#L327">argo-cd.configs.params."server.insecure"</a></td>
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
			<td id="argo-cd--configs--repositories--shortlink--name"><a href="./values.yaml#L320">argo-cd.configs.repositories.shortlink.name</a></td>
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
			<td id="argo-cd--configs--repositories--shortlink--type"><a href="./values.yaml#L321">argo-cd.configs.repositories.shortlink.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"git"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--configs--repositories--shortlink--url"><a href="./values.yaml#L319">argo-cd.configs.repositories.shortlink.url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://github.com/shortlink-org/shortlink"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--metrics--applicationLabels--enabled"><a href="./values.yaml#L36">argo-cd.controller.metrics.applicationLabels.enabled</a></td>
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
			<td id="argo-cd--controller--metrics--enabled"><a href="./values.yaml#L34">argo-cd.controller.metrics.enabled</a></td>
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
			<td id="argo-cd--controller--metrics--serviceMonitor--enabled"><a href="./values.yaml#L38">argo-cd.controller.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--controller--replicas"><a href="./values.yaml#L23">argo-cd.controller.replicas</a></td>
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
			<td id="argo-cd--controller--resources--limits--cpu"><a href="./values.yaml#L27">argo-cd.controller.resources.limits.cpu</a></td>
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
			<td id="argo-cd--controller--resources--limits--memory"><a href="./values.yaml#L28">argo-cd.controller.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"3Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--resources--requests--cpu"><a href="./values.yaml#L30">argo-cd.controller.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"250m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--resources--requests--memory"><a href="./values.yaml#L31">argo-cd.controller.resources.requests.memory</a></td>
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
			<td id="argo-cd--controller--rules--enabled"><a href="./values.yaml#L41">argo-cd.controller.rules.enabled</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--alert"><a href="./values.yaml#L43">argo-cd.controller.rules.spec[0].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ArgoAppMissing"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[0]--annotations--description"><a href="./values.yaml#L51">argo-cd.controller.rules.spec[0].annotations.description</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Argo CD has not reported any applications data for the past 15 minutes which means that it must be down or not functioning properly.  This needs to be resolved for this cloud to continue to maintain state.\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[0]--annotations--summary"><a href="./values.yaml#L50">argo-cd.controller.rules.spec[0].annotations.summary</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"[Argo CD] No reported applications"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[0]--expr"><a href="./values.yaml#L44">argo-cd.controller.rules.spec[0].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"absent(argocd_app_info)\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[0]--for"><a href="./values.yaml#L46">argo-cd.controller.rules.spec[0].for</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--labels--severity"><a href="./values.yaml#L48">argo-cd.controller.rules.spec[0].labels.severity</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--alert"><a href="./values.yaml#L55">argo-cd.controller.rules.spec[1].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ArgoAppNotSynced"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[1]--annotations--description"><a href="./values.yaml#L63">argo-cd.controller.rules.spec[1].annotations.description</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"The application [{{`{{$labels.name}}`}} has not been synchronized for over\n 12 hours which means that the state of this cloud has drifted away from the\n state inside Git.\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[1]--annotations--summary"><a href="./values.yaml#L62">argo-cd.controller.rules.spec[1].annotations.summary</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"[{{`{{$labels.name}}`}}] Application not synchronized"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[1]--expr"><a href="./values.yaml#L56">argo-cd.controller.rules.spec[1].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argocd_app_info{sync_status!=\"Synced\"} == 1\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[1]--for"><a href="./values.yaml#L58">argo-cd.controller.rules.spec[1].for</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"12h"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--spec[1]--labels--severity"><a href="./values.yaml#L60">argo-cd.controller.rules.spec[1].labels.severity</a></td>
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
			<td id="argo-cd--dex--enabled"><a href="./values.yaml#L69">argo-cd.dex.enabled</a></td>
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
			<td id="argo-cd--dex--env[0]--name"><a href="./values.yaml#L75">argo-cd.dex.env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ARGO_WORKFLOWS_SSO_CLIENT_SECRET"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--dex--env[0]--valueFrom--secretKeyRef--key"><a href="./values.yaml#L79">argo-cd.dex.env[0].valueFrom.secretKeyRef.key</a></td>
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
			<td id="argo-cd--dex--env[0]--valueFrom--secretKeyRef--name"><a href="./values.yaml#L78">argo-cd.dex.env[0].valueFrom.secretKeyRef.name</a></td>
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
			<td id="argo-cd--dex--image--tag"><a href="./values.yaml#L72">argo-cd.dex.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"latest-alpine"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--dex--metrics--enabled"><a href="./values.yaml#L90">argo-cd.dex.metrics.enabled</a></td>
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
			<td id="argo-cd--dex--metrics--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L94">argo-cd.dex.metrics.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="argo-cd--dex--metrics--serviceMonitor--enabled"><a href="./values.yaml#L92">argo-cd.dex.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--dex--resources--limits--cpu"><a href="./values.yaml#L83">argo-cd.dex.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--dex--resources--limits--memory"><a href="./values.yaml#L84">argo-cd.dex.resources.limits.memory</a></td>
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
			<td id="argo-cd--dex--resources--requests--cpu"><a href="./values.yaml#L86">argo-cd.dex.resources.requests.cpu</a></td>
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
			<td id="argo-cd--dex--resources--requests--memory"><a href="./values.yaml#L87">argo-cd.dex.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"32Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--enabled"><a href="./values.yaml#L7">argo-cd.enabled</a></td>
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
			<td id="argo-cd--externalRedis--host"><a href="./values.yaml#L100">argo-cd.externalRedis.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redis-master.argocd"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--fullnameOverride"><a href="./values.yaml#L9">argo-cd.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argocd"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--global--logging--format"><a href="./values.yaml#L19">argo-cd.global.logging.format</a></td>
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
			<td id="argo-cd--global--logging--level"><a href="./values.yaml#L20">argo-cd.global.logging.level</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"warn"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--global--networkPolicy--create"><a href="./values.yaml#L16">argo-cd.global.networkPolicy.create</a></td>
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
			<td id="argo-cd--notifications--metrics--enabled"><a href="./values.yaml#L339">argo-cd.notifications.metrics.enabled</a></td>
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
			<td id="argo-cd--notifications--metrics--serviceMonitor--enabled"><a href="./values.yaml#L341">argo-cd.notifications.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--notifications--resources--limits--cpu"><a href="./values.yaml#L332">argo-cd.notifications.resources.limits.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"300m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--notifications--resources--limits--memory"><a href="./values.yaml#L333">argo-cd.notifications.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"256Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--notifications--resources--requests--cpu"><a href="./values.yaml#L335">argo-cd.notifications.resources.requests.cpu</a></td>
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
			<td id="argo-cd--notifications--resources--requests--memory"><a href="./values.yaml#L336">argo-cd.notifications.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--redis--enabled"><a href="./values.yaml#L97">argo-cd.redis.enabled</a></td>
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
			<td id="argo-cd--repoServer--env[0]--name"><a href="./values.yaml#L184">argo-cd.repoServer.env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_PLUGINS"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[0]--value"><a href="./values.yaml#L185">argo-cd.repoServer.env[0].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools/helm-plugins/"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[1]--name"><a href="./values.yaml#L186">argo-cd.repoServer.env[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_SOPS_PATH"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[1]--value"><a href="./values.yaml#L187">argo-cd.repoServer.env[1].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools/sops"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[2]--name"><a href="./values.yaml#L188">argo-cd.repoServer.env[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_VALS_PATH"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[2]--value"><a href="./values.yaml#L189">argo-cd.repoServer.env[2].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools/vals"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[3]--name"><a href="./values.yaml#L190">argo-cd.repoServer.env[3].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_KUBECTL_PATH"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[3]--value"><a href="./values.yaml#L191">argo-cd.repoServer.env[3].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools/kubectl"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[4]--name"><a href="./values.yaml#L192">argo-cd.repoServer.env[4].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_CURL_PATH"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[4]--value"><a href="./values.yaml#L193">argo-cd.repoServer.env[4].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools/curl"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[5]--name"><a href="./values.yaml#L195">argo-cd.repoServer.env[5].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_VALUES_ALLOW_SYMLINKS"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[5]--value"><a href="./values.yaml#L196">argo-cd.repoServer.env[5].value</a></td>
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
			<td id="argo-cd--repoServer--env[6]--name"><a href="./values.yaml#L197">argo-cd.repoServer.env[6].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_VALUES_ALLOW_ABSOLUTE_PATH"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[6]--value"><a href="./values.yaml#L198">argo-cd.repoServer.env[6].value</a></td>
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
			<td id="argo-cd--repoServer--env[7]--name"><a href="./values.yaml#L199">argo-cd.repoServer.env[7].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_VALUES_ALLOW_PATH_TRAVERSAL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--env[7]--value"><a href="./values.yaml#L200">argo-cd.repoServer.env[7].value</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--args[0]"><a href="./values.yaml#L255">argo-cd.repoServer.initContainers[0].args[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"mkdir -p /custom-tools/helm-plugins\nwget -qO- https://github.com/jkroepke/helm-secrets/releases/download/v${HELM_SECRETS_VERSION}/helm-secrets.tar.gz | tar -C /custom-tools/helm-plugins -xzf-;\n\nwget -qO /custom-tools/sops https://github.com/mozilla/sops/releases/download/v${SOPS_VERSION}/sops-v${SOPS_VERSION}.linux.amd64\nwget -qO /custom-tools/kubectl https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl\n\nwget -qO- https://github.com/variantdev/vals/releases/download/v${VALS_VERSION}/vals_${VALS_VERSION}_linux_amd64.tar.gz | tar -xzf- -C /custom-tools/ vals;\n\n# helm secrets wrapper mode installation (optional)\n# RUN printf '#!/usr/bin/env sh\\nexec %s secrets \"$@\"' \"${HELM_SECRETS_HELM_PATH}\" \u003e\"/usr/local/sbin/helm\" \u0026\u0026 chmod +x \"/custom-tools/helm\"\n\nchmod +x /custom-tools/*\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--command[0]"><a href="./values.yaml#L235">argo-cd.repoServer.initContainers[0].command[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sh"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--command[1]"><a href="./values.yaml#L235">argo-cd.repoServer.initContainers[0].command[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"-ec"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[0]--name"><a href="./values.yaml#L246">argo-cd.repoServer.initContainers[0].env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HELM_SECRETS_VERSION"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[0]--value"><a href="./values.yaml#L247">argo-cd.repoServer.initContainers[0].env[0].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"4.6.0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[1]--name"><a href="./values.yaml#L248">argo-cd.repoServer.initContainers[0].env[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"KUBECTL_VERSION"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[1]--value"><a href="./values.yaml#L249">argo-cd.repoServer.initContainers[0].env[1].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1.30.0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[2]--name"><a href="./values.yaml#L250">argo-cd.repoServer.initContainers[0].env[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"VALS_VERSION"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[2]--value"><a href="./values.yaml#L251">argo-cd.repoServer.initContainers[0].env[2].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.37.0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[3]--name"><a href="./values.yaml#L252">argo-cd.repoServer.initContainers[0].env[3].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"SOPS_VERSION"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[3]--value"><a href="./values.yaml#L253">argo-cd.repoServer.initContainers[0].env[3].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"3.8.1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--image"><a href="./values.yaml#L234">argo-cd.repoServer.initContainers[0].image</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"alpine:latest"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--name"><a href="./values.yaml#L233">argo-cd.repoServer.initContainers[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"download-tools"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--allowPrivilegeEscalation"><a href="./values.yaml#L239">argo-cd.repoServer.initContainers[0].securityContext.allowPrivilegeEscalation</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--capabilities--drop[0]"><a href="./values.yaml#L242">argo-cd.repoServer.initContainers[0].securityContext.capabilities.drop[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ALL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--runAsNonRoot"><a href="./values.yaml#L237">argo-cd.repoServer.initContainers[0].securityContext.runAsNonRoot</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--runAsUser"><a href="./values.yaml#L238">argo-cd.repoServer.initContainers[0].securityContext.runAsUser</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--seccompProfile--type"><a href="./values.yaml#L244">argo-cd.repoServer.initContainers[0].securityContext.seccompProfile.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"RuntimeDefault"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--volumeMounts[0]--mountPath"><a href="./values.yaml#L269">argo-cd.repoServer.initContainers[0].volumeMounts[0].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--volumeMounts[0]--name"><a href="./values.yaml#L270">argo-cd.repoServer.initContainers[0].volumeMounts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"custom-tools"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--metrics--enabled"><a href="./values.yaml#L206">argo-cd.repoServer.metrics.enabled</a></td>
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
			<td id="argo-cd--repoServer--metrics--serviceMonitor--enabled"><a href="./values.yaml#L208">argo-cd.repoServer.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--repoServer--rbac[0]--apiGroups[0]"><a href="./values.yaml#L171">argo-cd.repoServer.rbac[0].apiGroups[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[0]--resources[0]"><a href="./values.yaml#L173">argo-cd.repoServer.rbac[0].resources[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"secrets"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--rbac[0]--verbs[0]"><a href="./values.yaml#L175">argo-cd.repoServer.rbac[0].verbs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"get"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--rbac[1]--apiGroups[0]"><a href="./values.yaml#L177">argo-cd.repoServer.rbac[1].apiGroups[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[1]--resources[0]"><a href="./values.yaml#L179">argo-cd.repoServer.rbac[1].resources[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"pods/exec"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--rbac[1]--verbs[0]"><a href="./values.yaml#L181">argo-cd.repoServer.rbac[1].verbs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"create"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--resources--requests--cpu"><a href="./values.yaml#L229">argo-cd.repoServer.resources.requests.cpu</a></td>
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
			<td id="argo-cd--repoServer--resources--requests--memory"><a href="./values.yaml#L230">argo-cd.repoServer.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"64Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--serviceAccount--create"><a href="./values.yaml#L166">argo-cd.repoServer.serviceAccount.create</a></td>
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
			<td id="argo-cd--repoServer--serviceAccount--name"><a href="./values.yaml#L167">argo-cd.repoServer.serviceAccount.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argocd-repo-server"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumeMounts[0]--mountPath"><a href="./values.yaml#L218">argo-cd.repoServer.volumeMounts[0].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/custom-tools"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumeMounts[0]--name"><a href="./values.yaml#L219">argo-cd.repoServer.volumeMounts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"custom-tools"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumeMounts[1]--mountPath"><a href="./values.yaml#L220">argo-cd.repoServer.volumeMounts[1].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/sops-gpg/"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumeMounts[1]--name"><a href="./values.yaml#L221">argo-cd.repoServer.volumeMounts[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sops-gpg"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumes[0]--emptyDir"><a href="./values.yaml#L212">argo-cd.repoServer.volumes[0].emptyDir</a></td>
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
			<td id="argo-cd--repoServer--volumes[0]--name"><a href="./values.yaml#L211">argo-cd.repoServer.volumes[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"custom-tools"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumes[1]--name"><a href="./values.yaml#L213">argo-cd.repoServer.volumes[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sops-gpg"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--volumes[1]--secret--secretName"><a href="./values.yaml#L215">argo-cd.repoServer.volumes[1].secret.secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"sops-gpg"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--enabled"><a href="./values.yaml#L146">argo-cd.server.extensions.enabled</a></td>
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
			<td id="argo-cd--server--extensions--extensionList[0]--env[0]--name"><a href="./values.yaml#L151">argo-cd.server.extensions.extensionList[0].env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"EXTENSION_URL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[0]--env[0]--value"><a href="./values.yaml#L152">argo-cd.server.extensions.extensionList[0].env[0].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://github.com/argoproj-labs/rollout-extension/releases/download/v0.3.5/extension.tar"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[0]--name"><a href="./values.yaml#L149">argo-cd.server.extensions.extensionList[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"rollout-extension"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[1]--env[0]--name"><a href="./values.yaml#L155">argo-cd.server.extensions.extensionList[1].env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"EXTENSION_URL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[1]--env[0]--value"><a href="./values.yaml#L156">argo-cd.server.extensions.extensionList[1].env[0].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://github.com/argoproj-labs/argocd-extension-metrics/releases/download/v1.0.1/extension.tar.gz"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[1]--env[1]--name"><a href="./values.yaml#L157">argo-cd.server.extensions.extensionList[1].env[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"EXTENSION_CHECKSUM_URL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[1]--env[1]--value"><a href="./values.yaml#L158">argo-cd.server.extensions.extensionList[1].env[1].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://github.com/argoproj-labs/argocd-extension-metrics/releases/download/v1.0.1/extension_checksums.txt"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[1]--name"><a href="./values.yaml#L153">argo-cd.server.extensions.extensionList[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"extension-metrics"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[2]--env[0]--name"><a href="./values.yaml#L161">argo-cd.server.extensions.extensionList[2].env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"EXTENSION_URL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[2]--env[0]--value"><a href="./values.yaml#L162">argo-cd.server.extensions.extensionList[2].env[0].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://github.com/shortlink-org/shortlink/raw/main/boundaries/platform/argocd-extension-docs/ui/dist/extension.tar"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--extensions--extensionList[2]--name"><a href="./values.yaml#L159">argo-cd.server.extensions.extensionList[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink-extension-docs"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L114">argo-cd.server.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/backend-protocol""><a href="./values.yaml#L115">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/configuration-snippet""><a href="./values.yaml#L117">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"proxy_ssl_server_name on;\nproxy_ssl_name $host;"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L123">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L122">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/secure-backends""><a href="./values.yaml#L121">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/secure-backends"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/ssl-passthrough""><a href="./values.yaml#L116">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/ssl-passthrough"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/ssl-redirect""><a href="./values.yaml#L120">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/ssl-redirect"</a></td>
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
			<td id="argo-cd--server--ingress--enabled"><a href="./values.yaml#L109">argo-cd.server.ingress.enabled</a></td>
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
			<td id="argo-cd--server--ingress--extraTls[0]--hosts[0]"><a href="./values.yaml#L130">argo-cd.server.ingress.extraTls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--ingress--extraTls[0]--secretName"><a href="./values.yaml#L128">argo-cd.server.ingress.extraTls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--ingress--hostname"><a href="./values.yaml#L125">argo-cd.server.ingress.hostname</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argo.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--ingress--ingressClassName"><a href="./values.yaml#L111">argo-cd.server.ingress.ingressClassName</a></td>
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
			<td id="argo-cd--server--metrics--enabled"><a href="./values.yaml#L104">argo-cd.server.metrics.enabled</a></td>
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
			<td id="argo-cd--server--metrics--serviceMonitor--enabled"><a href="./values.yaml#L106">argo-cd.server.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--server--rbac--"policy--csv""><a href="./values.yaml#L133">argo-cd.server.rbac."policy.csv"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"p, role:readonly, extensions, invoke, httpbin, allow\np, role:org-admin, applications, *, */*, allow\np, role:org-admin, clusters, get, *, allow\np, role:org-admin, repositories, get, *, allow\np, role:org-admin, repositories, create, *, allow\np, role:org-admin, repositories, update, *, allow\np, role:org-admin, repositories, delete, *, allow\np, role:org-admin, exec, create, */*, allow\ng, shortlink-org:devops, role:org-admin\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--rbac--"policy--default""><a href="./values.yaml#L143">argo-cd.server.rbac."policy.default"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"role:readonly"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
