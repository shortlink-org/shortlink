# kyverno

![Version: 0.2.1](https://img.shields.io/badge/Version-0.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://kyverno.github.io/kyverno | kyverno | 3.2.1 |
| https://kyverno.github.io/kyverno | kyverno-policies | 3.2.1 |
| https://kyverno.github.io/policy-reporter | policy-reporter | 2.22.5 |

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
			<td id="kyverno-policies--background"><a href="./values.yaml#L92">kyverno-policies.background</a></td>
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
			<td id="kyverno-policies--enabled"><a href="./values.yaml#L88">kyverno-policies.enabled</a></td>
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
			<td id="kyverno-policies--failurePolicy"><a href="./values.yaml#L99">kyverno-policies.failurePolicy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Ignore"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--podSecuritySeverity"><a href="./values.yaml#L90">kyverno-policies.podSecuritySeverity</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"low"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--validationFailureActionByPolicy--disallow-capabilities-strict"><a href="./values.yaml#L95">kyverno-policies.validationFailureActionByPolicy.disallow-capabilities-strict</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"audit"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--validationFailureActionByPolicy--disallow-host-path"><a href="./values.yaml#L96">kyverno-policies.validationFailureActionByPolicy.disallow-host-path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"audit"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno-policies--validationFailureActionByPolicy--disallow-host-ports"><a href="./values.yaml#L97">kyverno-policies.validationFailureActionByPolicy.disallow-host-ports</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"audit"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--admissionController--hostNetwork"><a href="./values.yaml#L14">kyverno.admissionController.hostNetwork</a></td>
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
			<td id="kyverno--admissionController--networkPolicy--enabled"><a href="./values.yaml#L22">kyverno.admissionController.networkPolicy.enabled</a></td>
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
			<td id="kyverno--admissionController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L28">kyverno.admissionController.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="kyverno--admissionController--serviceMonitor--enabled"><a href="./values.yaml#L25">kyverno.admissionController.serviceMonitor.enabled</a></td>
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
			<td id="kyverno--admissionController--tracing--address"><a href="./values.yaml#L18">kyverno.admissionController.tracing.address</a></td>
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
			<td id="kyverno--admissionController--tracing--enabled"><a href="./values.yaml#L17">kyverno.admissionController.tracing.enabled</a></td>
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
			<td id="kyverno--admissionController--tracing--port"><a href="./values.yaml#L19">kyverno.admissionController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--backgroundController--enabled"><a href="./values.yaml#L51">kyverno.backgroundController.enabled</a></td>
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
			<td id="kyverno--backgroundController--networkPolicy--enabled"><a href="./values.yaml#L54">kyverno.backgroundController.networkPolicy.enabled</a></td>
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
			<td id="kyverno--backgroundController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L65">kyverno.backgroundController.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="kyverno--backgroundController--serviceMonitor--enabled"><a href="./values.yaml#L62">kyverno.backgroundController.serviceMonitor.enabled</a></td>
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
			<td id="kyverno--backgroundController--tracing--address"><a href="./values.yaml#L58">kyverno.backgroundController.tracing.address</a></td>
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
			<td id="kyverno--backgroundController--tracing--enabled"><a href="./values.yaml#L57">kyverno.backgroundController.tracing.enabled</a></td>
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
			<td id="kyverno--backgroundController--tracing--port"><a href="./values.yaml#L59">kyverno.backgroundController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--cleanupController--enabled"><a href="./values.yaml#L68">kyverno.cleanupController.enabled</a></td>
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
			<td id="kyverno--cleanupController--logging--format"><a href="./values.yaml#L80">kyverno.cleanupController.logging.format</a></td>
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
			<td id="kyverno--cleanupController--networkPolicy--enabled"><a href="./values.yaml#L71">kyverno.cleanupController.networkPolicy.enabled</a></td>
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
			<td id="kyverno--cleanupController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L77">kyverno.cleanupController.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="kyverno--cleanupController--serviceMonitor--enabled"><a href="./values.yaml#L74">kyverno.cleanupController.serviceMonitor.enabled</a></td>
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
			<td id="kyverno--cleanupController--tracing--address"><a href="./values.yaml#L84">kyverno.cleanupController.tracing.address</a></td>
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
			<td id="kyverno--cleanupController--tracing--enabled"><a href="./values.yaml#L83">kyverno.cleanupController.tracing.enabled</a></td>
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
			<td id="kyverno--cleanupController--tracing--port"><a href="./values.yaml#L85">kyverno.cleanupController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--config--resourceFiltersExcludeNamespaces[0]"><a href="./values.yaml#L10">kyverno.config.resourceFiltersExcludeNamespaces[0]</a></td>
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
			<td id="kyverno--config--resourceFiltersExcludeNamespaces[1]"><a href="./values.yaml#L11">kyverno.config.resourceFiltersExcludeNamespaces[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kube-system"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--enabled"><a href="./values.yaml#L6">kyverno.enabled</a></td>
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
			<td id="kyverno--reportsController--enabled"><a href="./values.yaml#L34">kyverno.reportsController.enabled</a></td>
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
			<td id="kyverno--reportsController--networkPolicy--enabled"><a href="./values.yaml#L37">kyverno.reportsController.networkPolicy.enabled</a></td>
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
			<td id="kyverno--reportsController--serviceMonitor--additionalLabels--release"><a href="./values.yaml#L48">kyverno.reportsController.serviceMonitor.additionalLabels.release</a></td>
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
			<td id="kyverno--reportsController--serviceMonitor--enabled"><a href="./values.yaml#L45">kyverno.reportsController.serviceMonitor.enabled</a></td>
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
			<td id="kyverno--reportsController--tracing--address"><a href="./values.yaml#L41">kyverno.reportsController.tracing.address</a></td>
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
			<td id="kyverno--reportsController--tracing--enabled"><a href="./values.yaml#L40">kyverno.reportsController.tracing.enabled</a></td>
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
			<td id="kyverno--reportsController--tracing--port"><a href="./values.yaml#L42">kyverno.reportsController.tracing.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
4317
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kyverno--webhooksCleanup--enabled"><a href="./values.yaml#L31">kyverno.webhooksCleanup.enabled</a></td>
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
			<td id="policy-reporter--enabled"><a href="./values.yaml#L102">policy-reporter.enabled</a></td>
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
			<td id="policy-reporter--global--plugins--kyverno"><a href="./values.yaml#L176">policy-reporter.global.plugins.kyverno</a></td>
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
			<td id="policy-reporter--kyvernoPlugin--enabled"><a href="./values.yaml#L164">policy-reporter.kyvernoPlugin.enabled</a></td>
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
			<td id="policy-reporter--logging--encoding"><a href="./values.yaml#L116">policy-reporter.logging.encoding</a></td>
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
			<td id="policy-reporter--metrics--enabled"><a href="./values.yaml#L122">policy-reporter.metrics.enabled</a></td>
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
			<td id="policy-reporter--monitoring--enabled"><a href="./values.yaml#L167">policy-reporter.monitoring.enabled</a></td>
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
			<td id="policy-reporter--monitoring--grafana--folder--annotation"><a href="./values.yaml#L171">policy-reporter.monitoring.grafana.folder.annotation</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana_dashboard_folder"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--monitoring--grafana--folder--name"><a href="./values.yaml#L172">policy-reporter.monitoring.grafana.folder.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Security"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--networkPolicy--enabled"><a href="./values.yaml#L113">policy-reporter.networkPolicy.enabled</a></td>
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
			<td id="policy-reporter--profiling--enabled"><a href="./values.yaml#L125">policy-reporter.profiling.enabled</a></td>
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
			<td id="policy-reporter--resources--limits--cpu"><a href="./values.yaml#L106">policy-reporter.resources.limits.cpu</a></td>
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
			<td id="policy-reporter--resources--limits--memory"><a href="./values.yaml#L107">policy-reporter.resources.limits.memory</a></td>
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
			<td id="policy-reporter--resources--requests--cpu"><a href="./values.yaml#L109">policy-reporter.resources.requests.cpu</a></td>
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
			<td id="policy-reporter--resources--requests--memory"><a href="./values.yaml#L110">policy-reporter.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"75Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--rest--enabled"><a href="./values.yaml#L119">policy-reporter.rest.enabled</a></td>
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
			<td id="policy-reporter--target--loki--host"><a href="./values.yaml#L180">policy-reporter.target.loki.host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-loki.grafana:3100"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--target--loki--minimumPriority"><a href="./values.yaml#L181">policy-reporter.target.loki.minimumPriority</a></td>
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
			<td id="policy-reporter--target--loki--skipExistingOnStartup"><a href="./values.yaml#L182">policy-reporter.target.loki.skipExistingOnStartup</a></td>
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
			<td id="policy-reporter--target--loki--sources[0]"><a href="./values.yaml#L184">policy-reporter.target.loki.sources[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kyverno"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--enabled"><a href="./values.yaml#L128">policy-reporter.ui.enabled</a></td>
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
			<td id="policy-reporter--ui--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L146">policy-reporter.ui.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L148">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L147">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/rewrite-target""><a href="./values.yaml#L149">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/rewrite-target"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/$1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--annotations--"nginx--ingress--kubernetes--io/use-regex""><a href="./values.yaml#L150">policy-reporter.ui.ingress.annotations."nginx.ingress.kubernetes.io/use-regex"</a></td>
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
			<td id="policy-reporter--ui--ingress--className"><a href="./values.yaml#L143">policy-reporter.ui.ingress.className</a></td>
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
			<td id="policy-reporter--ui--ingress--enabled"><a href="./values.yaml#L142">policy-reporter.ui.ingress.enabled</a></td>
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
			<td id="policy-reporter--ui--ingress--hosts[0]--host"><a href="./values.yaml#L153">policy-reporter.ui.ingress.hosts[0].host</a></td>
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
			<td id="policy-reporter--ui--ingress--hosts[0]--paths[0]--path"><a href="./values.yaml#L155">policy-reporter.ui.ingress.hosts[0].paths[0].path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/kyverno/?(.*)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="policy-reporter--ui--ingress--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L156">policy-reporter.ui.ingress.hosts[0].paths[0].pathType</a></td>
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
			<td id="policy-reporter--ui--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L161">policy-reporter.ui.ingress.tls[0].hosts[0]</a></td>
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
			<td id="policy-reporter--ui--ingress--tls[0]--secretName"><a href="./values.yaml#L159">policy-reporter.ui.ingress.tls[0].secretName</a></td>
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
			<td id="policy-reporter--ui--plugins--kyverno"><a href="./values.yaml#L131">policy-reporter.ui.plugins.kyverno</a></td>
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
			<td id="policy-reporter--ui--resources--limits--cpu"><a href="./values.yaml#L135">policy-reporter.ui.resources.limits.cpu</a></td>
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
			<td id="policy-reporter--ui--resources--limits--memory"><a href="./values.yaml#L136">policy-reporter.ui.resources.limits.memory</a></td>
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
			<td id="policy-reporter--ui--resources--requests--cpu"><a href="./values.yaml#L138">policy-reporter.ui.resources.requests.cpu</a></td>
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
			<td id="policy-reporter--ui--resources--requests--memory"><a href="./values.yaml#L139">policy-reporter.ui.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"45Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
