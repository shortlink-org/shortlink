# argo

![Version: 0.4.4](https://img.shields.io/badge/Version-0.4.4-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 2.10.0](https://img.shields.io/badge/AppVersion-2.10.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://argoproj.github.io/argo-helm | argo-cd | 6.7.3 |
| https://argoproj.github.io/argo-helm | argocd-apps | 2.0.0 |
| https://argoproj.github.io/argo-helm | argocd-image-updater | 0.9.6 |
| oci://ghcr.io/dragonflydb/dragonfly/helm | dragonfly | v1.15.1 |

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
			<td id="argo-cd--applicationSet--metrics--enabled"><a href="./values.yaml#L261">argo-cd.applicationSet.metrics.enabled</a></td>
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
			<td id="argo-cd--applicationSet--metrics--serviceMonitor--enabled"><a href="./values.yaml#L263">argo-cd.applicationSet.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--applicationSet--resources--limits--cpu"><a href="./values.yaml#L254">argo-cd.applicationSet.resources.limits.cpu</a></td>
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
			<td id="argo-cd--applicationSet--resources--limits--memory"><a href="./values.yaml#L255">argo-cd.applicationSet.resources.limits.memory</a></td>
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
			<td id="argo-cd--applicationSet--resources--requests--cpu"><a href="./values.yaml#L257">argo-cd.applicationSet.resources.requests.cpu</a></td>
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
			<td id="argo-cd--applicationSet--resources--requests--memory"><a href="./values.yaml#L258">argo-cd.applicationSet.resources.requests.memory</a></td>
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
			<td id="argo-cd--configs--cm--"controller--diff--server--side""><a href="./values.yaml#L271">argo-cd.configs.cm."controller.diff.server.side"</a></td>
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
			<td id="argo-cd--configs--cm--"exec--enabled""><a href="./values.yaml#L273">argo-cd.configs.cm."exec.enabled"</a></td>
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
			<td id="argo-cd--configs--cm--"helm--valuesFileSchemes""><a href="./values.yaml#L275">argo-cd.configs.cm."helm.valuesFileSchemes"</a></td>
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
			<td id="argo-cd--configs--cm--"resource--compareoptions""><a href="./values.yaml#L282">argo-cd.configs.cm."resource.compareoptions"</a></td>
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
			<td id="argo-cd--configs--cm--"resource--exclusions""><a href="./values.yaml#L286">argo-cd.configs.cm."resource.exclusions"</a></td>
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
			<td id="argo-cd--configs--cm--"statusbadge--enabled""><a href="./values.yaml#L269">argo-cd.configs.cm."statusbadge.enabled"</a></td>
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
			<td id="argo-cd--configs--cm--url"><a href="./values.yaml#L267">argo-cd.configs.cm.url</a></td>
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
			<td id="argo-cd--configs--params--"dexserver--disable--tls""><a href="./values.yaml#L303">argo-cd.configs.params."dexserver.disable.tls"</a></td>
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
			<td id="argo-cd--configs--params--"otlp--address""><a href="./values.yaml#L301">argo-cd.configs.params."otlp.address"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana:4317"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--configs--repositories--shortlink--name"><a href="./values.yaml#L297">argo-cd.configs.repositories.shortlink.name</a></td>
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
			<td id="argo-cd--configs--repositories--shortlink--type"><a href="./values.yaml#L298">argo-cd.configs.repositories.shortlink.type</a></td>
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
			<td id="argo-cd--configs--repositories--shortlink--url"><a href="./values.yaml#L296">argo-cd.configs.repositories.shortlink.url</a></td>
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
			<td id="argo-cd--controller--metrics--applicationLabels--enabled"><a href="./values.yaml#L33">argo-cd.controller.metrics.applicationLabels.enabled</a></td>
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
			<td id="argo-cd--controller--metrics--enabled"><a href="./values.yaml#L31">argo-cd.controller.metrics.enabled</a></td>
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
			<td id="argo-cd--controller--metrics--serviceMonitor--enabled"><a href="./values.yaml#L35">argo-cd.controller.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--controller--replicas"><a href="./values.yaml#L20">argo-cd.controller.replicas</a></td>
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
			<td id="argo-cd--controller--resources--limits--cpu"><a href="./values.yaml#L24">argo-cd.controller.resources.limits.cpu</a></td>
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
			<td id="argo-cd--controller--resources--limits--memory"><a href="./values.yaml#L25">argo-cd.controller.resources.limits.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"2Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--resources--requests--cpu"><a href="./values.yaml#L27">argo-cd.controller.resources.requests.cpu</a></td>
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
			<td id="argo-cd--controller--resources--requests--memory"><a href="./values.yaml#L28">argo-cd.controller.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1Gi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--controller--rules--enabled"><a href="./values.yaml#L38">argo-cd.controller.rules.enabled</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--alert"><a href="./values.yaml#L40">argo-cd.controller.rules.spec[0].alert</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--annotations--description"><a href="./values.yaml#L48">argo-cd.controller.rules.spec[0].annotations.description</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--annotations--summary"><a href="./values.yaml#L47">argo-cd.controller.rules.spec[0].annotations.summary</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--expr"><a href="./values.yaml#L41">argo-cd.controller.rules.spec[0].expr</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--for"><a href="./values.yaml#L43">argo-cd.controller.rules.spec[0].for</a></td>
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
			<td id="argo-cd--controller--rules--spec[0]--labels--severity"><a href="./values.yaml#L45">argo-cd.controller.rules.spec[0].labels.severity</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--alert"><a href="./values.yaml#L52">argo-cd.controller.rules.spec[1].alert</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--annotations--description"><a href="./values.yaml#L60">argo-cd.controller.rules.spec[1].annotations.description</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--annotations--summary"><a href="./values.yaml#L59">argo-cd.controller.rules.spec[1].annotations.summary</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--expr"><a href="./values.yaml#L53">argo-cd.controller.rules.spec[1].expr</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--for"><a href="./values.yaml#L55">argo-cd.controller.rules.spec[1].for</a></td>
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
			<td id="argo-cd--controller--rules--spec[1]--labels--severity"><a href="./values.yaml#L57">argo-cd.controller.rules.spec[1].labels.severity</a></td>
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
			<td id="argo-cd--dex--enabled"><a href="./values.yaml#L66">argo-cd.dex.enabled</a></td>
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
			<td id="argo-cd--dex--env[0]--name"><a href="./values.yaml#L72">argo-cd.dex.env[0].name</a></td>
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
			<td id="argo-cd--dex--env[0]--valueFrom--secretKeyRef--key"><a href="./values.yaml#L76">argo-cd.dex.env[0].valueFrom.secretKeyRef.key</a></td>
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
			<td id="argo-cd--dex--env[0]--valueFrom--secretKeyRef--name"><a href="./values.yaml#L75">argo-cd.dex.env[0].valueFrom.secretKeyRef.name</a></td>
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
			<td id="argo-cd--dex--image--tag"><a href="./values.yaml#L69">argo-cd.dex.image.tag</a></td>
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
			<td id="argo-cd--dex--metrics--enabled"><a href="./values.yaml#L87">argo-cd.dex.metrics.enabled</a></td>
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
			<td id="argo-cd--dex--metrics--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L91">argo-cd.dex.metrics.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="argo-cd--dex--metrics--serviceMonitor--enabled"><a href="./values.yaml#L89">argo-cd.dex.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--dex--resources--limits--cpu"><a href="./values.yaml#L80">argo-cd.dex.resources.limits.cpu</a></td>
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
			<td id="argo-cd--dex--resources--limits--memory"><a href="./values.yaml#L81">argo-cd.dex.resources.limits.memory</a></td>
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
			<td id="argo-cd--dex--resources--requests--cpu"><a href="./values.yaml#L83">argo-cd.dex.resources.requests.cpu</a></td>
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
			<td id="argo-cd--dex--resources--requests--memory"><a href="./values.yaml#L84">argo-cd.dex.resources.requests.memory</a></td>
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
			<td id="argo-cd--externalRedis--host"><a href="./values.yaml#L97">argo-cd.externalRedis.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"argocd-dragonfly.argocd"
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
			<td id="argo-cd--global--logging--format"><a href="./values.yaml#L16">argo-cd.global.logging.format</a></td>
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
			<td id="argo-cd--global--logging--level"><a href="./values.yaml#L17">argo-cd.global.logging.level</a></td>
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
			<td id="argo-cd--notifications--metrics--enabled"><a href="./values.yaml#L315">argo-cd.notifications.metrics.enabled</a></td>
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
			<td id="argo-cd--notifications--metrics--serviceMonitor--enabled"><a href="./values.yaml#L317">argo-cd.notifications.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--notifications--resources--limits--cpu"><a href="./values.yaml#L308">argo-cd.notifications.resources.limits.cpu</a></td>
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
			<td id="argo-cd--notifications--resources--limits--memory"><a href="./values.yaml#L309">argo-cd.notifications.resources.limits.memory</a></td>
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
			<td id="argo-cd--notifications--resources--requests--cpu"><a href="./values.yaml#L311">argo-cd.notifications.resources.requests.cpu</a></td>
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
			<td id="argo-cd--notifications--resources--requests--memory"><a href="./values.yaml#L312">argo-cd.notifications.resources.requests.memory</a></td>
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
			<td id="argo-cd--redis--enabled"><a href="./values.yaml#L94">argo-cd.redis.enabled</a></td>
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
			<td id="argo-cd--repoServer--env[0]--name"><a href="./values.yaml#L163">argo-cd.repoServer.env[0].name</a></td>
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
			<td id="argo-cd--repoServer--env[0]--value"><a href="./values.yaml#L164">argo-cd.repoServer.env[0].value</a></td>
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
			<td id="argo-cd--repoServer--env[1]--name"><a href="./values.yaml#L165">argo-cd.repoServer.env[1].name</a></td>
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
			<td id="argo-cd--repoServer--env[1]--value"><a href="./values.yaml#L166">argo-cd.repoServer.env[1].value</a></td>
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
			<td id="argo-cd--repoServer--env[2]--name"><a href="./values.yaml#L167">argo-cd.repoServer.env[2].name</a></td>
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
			<td id="argo-cd--repoServer--env[2]--value"><a href="./values.yaml#L168">argo-cd.repoServer.env[2].value</a></td>
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
			<td id="argo-cd--repoServer--env[3]--name"><a href="./values.yaml#L169">argo-cd.repoServer.env[3].name</a></td>
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
			<td id="argo-cd--repoServer--env[3]--value"><a href="./values.yaml#L170">argo-cd.repoServer.env[3].value</a></td>
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
			<td id="argo-cd--repoServer--env[4]--name"><a href="./values.yaml#L171">argo-cd.repoServer.env[4].name</a></td>
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
			<td id="argo-cd--repoServer--env[4]--value"><a href="./values.yaml#L172">argo-cd.repoServer.env[4].value</a></td>
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
			<td id="argo-cd--repoServer--env[5]--name"><a href="./values.yaml#L174">argo-cd.repoServer.env[5].name</a></td>
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
			<td id="argo-cd--repoServer--env[5]--value"><a href="./values.yaml#L175">argo-cd.repoServer.env[5].value</a></td>
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
			<td id="argo-cd--repoServer--env[6]--name"><a href="./values.yaml#L176">argo-cd.repoServer.env[6].name</a></td>
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
			<td id="argo-cd--repoServer--env[6]--value"><a href="./values.yaml#L177">argo-cd.repoServer.env[6].value</a></td>
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
			<td id="argo-cd--repoServer--env[7]--name"><a href="./values.yaml#L178">argo-cd.repoServer.env[7].name</a></td>
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
			<td id="argo-cd--repoServer--env[7]--value"><a href="./values.yaml#L179">argo-cd.repoServer.env[7].value</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--args[0]"><a href="./values.yaml#L234">argo-cd.repoServer.initContainers[0].args[0]</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--command[0]"><a href="./values.yaml#L214">argo-cd.repoServer.initContainers[0].command[0]</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--command[1]"><a href="./values.yaml#L214">argo-cd.repoServer.initContainers[0].command[1]</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--env[0]--name"><a href="./values.yaml#L225">argo-cd.repoServer.initContainers[0].env[0].name</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--env[0]--value"><a href="./values.yaml#L226">argo-cd.repoServer.initContainers[0].env[0].value</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--env[1]--name"><a href="./values.yaml#L227">argo-cd.repoServer.initContainers[0].env[1].name</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--env[1]--value"><a href="./values.yaml#L228">argo-cd.repoServer.initContainers[0].env[1].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1.29.1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[2]--name"><a href="./values.yaml#L229">argo-cd.repoServer.initContainers[0].env[2].name</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--env[2]--value"><a href="./values.yaml#L230">argo-cd.repoServer.initContainers[0].env[2].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.36.0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--repoServer--initContainers[0]--env[3]--name"><a href="./values.yaml#L231">argo-cd.repoServer.initContainers[0].env[3].name</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--env[3]--value"><a href="./values.yaml#L232">argo-cd.repoServer.initContainers[0].env[3].value</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--image"><a href="./values.yaml#L213">argo-cd.repoServer.initContainers[0].image</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--name"><a href="./values.yaml#L212">argo-cd.repoServer.initContainers[0].name</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--allowPrivilegeEscalation"><a href="./values.yaml#L218">argo-cd.repoServer.initContainers[0].securityContext.allowPrivilegeEscalation</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--capabilities--drop[0]"><a href="./values.yaml#L221">argo-cd.repoServer.initContainers[0].securityContext.capabilities.drop[0]</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--runAsNonRoot"><a href="./values.yaml#L216">argo-cd.repoServer.initContainers[0].securityContext.runAsNonRoot</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--runAsUser"><a href="./values.yaml#L217">argo-cd.repoServer.initContainers[0].securityContext.runAsUser</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--securityContext--seccompProfile--type"><a href="./values.yaml#L223">argo-cd.repoServer.initContainers[0].securityContext.seccompProfile.type</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--volumeMounts[0]--mountPath"><a href="./values.yaml#L248">argo-cd.repoServer.initContainers[0].volumeMounts[0].mountPath</a></td>
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
			<td id="argo-cd--repoServer--initContainers[0]--volumeMounts[0]--name"><a href="./values.yaml#L249">argo-cd.repoServer.initContainers[0].volumeMounts[0].name</a></td>
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
			<td id="argo-cd--repoServer--metrics--enabled"><a href="./values.yaml#L185">argo-cd.repoServer.metrics.enabled</a></td>
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
			<td id="argo-cd--repoServer--metrics--serviceMonitor--enabled"><a href="./values.yaml#L187">argo-cd.repoServer.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--repoServer--rbac[0]--apiGroups[0]"><a href="./values.yaml#L150">argo-cd.repoServer.rbac[0].apiGroups[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[0]--resources[0]"><a href="./values.yaml#L152">argo-cd.repoServer.rbac[0].resources[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[0]--verbs[0]"><a href="./values.yaml#L154">argo-cd.repoServer.rbac[0].verbs[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[1]--apiGroups[0]"><a href="./values.yaml#L156">argo-cd.repoServer.rbac[1].apiGroups[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[1]--resources[0]"><a href="./values.yaml#L158">argo-cd.repoServer.rbac[1].resources[0]</a></td>
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
			<td id="argo-cd--repoServer--rbac[1]--verbs[0]"><a href="./values.yaml#L160">argo-cd.repoServer.rbac[1].verbs[0]</a></td>
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
			<td id="argo-cd--repoServer--resources--requests--cpu"><a href="./values.yaml#L208">argo-cd.repoServer.resources.requests.cpu</a></td>
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
			<td id="argo-cd--repoServer--resources--requests--memory"><a href="./values.yaml#L209">argo-cd.repoServer.resources.requests.memory</a></td>
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
			<td id="argo-cd--repoServer--serviceAccount--create"><a href="./values.yaml#L145">argo-cd.repoServer.serviceAccount.create</a></td>
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
			<td id="argo-cd--repoServer--serviceAccount--name"><a href="./values.yaml#L146">argo-cd.repoServer.serviceAccount.name</a></td>
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
			<td id="argo-cd--repoServer--volumeMounts[0]--mountPath"><a href="./values.yaml#L197">argo-cd.repoServer.volumeMounts[0].mountPath</a></td>
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
			<td id="argo-cd--repoServer--volumeMounts[0]--name"><a href="./values.yaml#L198">argo-cd.repoServer.volumeMounts[0].name</a></td>
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
			<td id="argo-cd--repoServer--volumeMounts[1]--mountPath"><a href="./values.yaml#L199">argo-cd.repoServer.volumeMounts[1].mountPath</a></td>
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
			<td id="argo-cd--repoServer--volumeMounts[1]--name"><a href="./values.yaml#L200">argo-cd.repoServer.volumeMounts[1].name</a></td>
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
			<td id="argo-cd--repoServer--volumes[0]--emptyDir"><a href="./values.yaml#L191">argo-cd.repoServer.volumes[0].emptyDir</a></td>
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
			<td id="argo-cd--repoServer--volumes[0]--name"><a href="./values.yaml#L190">argo-cd.repoServer.volumes[0].name</a></td>
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
			<td id="argo-cd--repoServer--volumes[1]--name"><a href="./values.yaml#L192">argo-cd.repoServer.volumes[1].name</a></td>
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
			<td id="argo-cd--repoServer--volumes[1]--secret--secretName"><a href="./values.yaml#L194">argo-cd.repoServer.volumes[1].secret.secretName</a></td>
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
			<td id="argo-cd--server--extensions--enabled"><a href="./values.yaml#L141">argo-cd.server.extensions.enabled</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L111">argo-cd.server.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/backend-protocol""><a href="./values.yaml#L112">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/backend-protocol"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"HTTPS"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/configuration-snippet""><a href="./values.yaml#L113">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/configuration-snippet"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L119">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L118">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/secure-backends""><a href="./values.yaml#L117">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/secure-backends"</a></td>
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
			<td id="argo-cd--server--ingress--annotations--"nginx--ingress--kubernetes--io/ssl-redirect""><a href="./values.yaml#L116">argo-cd.server.ingress.annotations."nginx.ingress.kubernetes.io/ssl-redirect"</a></td>
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
			<td id="argo-cd--server--ingress--enabled"><a href="./values.yaml#L106">argo-cd.server.ingress.enabled</a></td>
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
			<td id="argo-cd--server--ingress--extraTls[0]--hosts[0]"><a href="./values.yaml#L126">argo-cd.server.ingress.extraTls[0].hosts[0]</a></td>
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
			<td id="argo-cd--server--ingress--extraTls[0]--secretName"><a href="./values.yaml#L124">argo-cd.server.ingress.extraTls[0].secretName</a></td>
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
			<td id="argo-cd--server--ingress--hostname"><a href="./values.yaml#L121">argo-cd.server.ingress.hostname</a></td>
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
			<td id="argo-cd--server--ingress--ingressClassName"><a href="./values.yaml#L108">argo-cd.server.ingress.ingressClassName</a></td>
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
			<td id="argo-cd--server--metrics--enabled"><a href="./values.yaml#L101">argo-cd.server.metrics.enabled</a></td>
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
			<td id="argo-cd--server--metrics--serviceMonitor--enabled"><a href="./values.yaml#L103">argo-cd.server.metrics.serviceMonitor.enabled</a></td>
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
			<td id="argo-cd--server--rbac--"policy--csv""><a href="./values.yaml#L129">argo-cd.server.rbac."policy.csv"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"p, role:org-admin, applications, *, */*, allow\np, role:org-admin, clusters, get, *, allow\np, role:org-admin, repositories, get, *, allow\np, role:org-admin, repositories, create, *, allow\np, role:org-admin, repositories, update, *, allow\np, role:org-admin, repositories, delete, *, allow\np, role:org-admin, exec, create, */*, allow\ng, shortlink-org:devops, role:org-admin\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="argo-cd--server--rbac--"policy--default""><a href="./values.yaml#L138">argo-cd.server.rbac."policy.default"</a></td>
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
