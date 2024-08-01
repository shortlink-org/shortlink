# falcosecurity

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://falcosecurity.github.io/charts | falco | 4.7.2 |
| https://falcosecurity.github.io/charts | falco-exporter | 0.12.1 |

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
			<td id="events--deletePodEvent"><a href="./values.yaml#L6">events.deletePodEvent</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"{{inputs.parameters.falco-event}}"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco-exporter--enabled"><a href="./values.yaml#L72">falco-exporter.enabled</a></td>
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
			<td id="falco-exporter--fullnameOverride"><a href="./values.yaml#L74">falco-exporter.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"falco-exporter"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco-exporter--grafanaDashboard--enabled"><a href="./values.yaml#L83">falco-exporter.grafanaDashboard.enabled</a></td>
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
			<td id="falco-exporter--grafanaDashboard--folder"><a href="./values.yaml#L85">falco-exporter.grafanaDashboard.folder</a></td>
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
			<td id="falco-exporter--grafanaDashboard--namespace"><a href="./values.yaml#L84">falco-exporter.grafanaDashboard.namespace</a></td>
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
			<td id="falco-exporter--prometheusRules--enabled"><a href="./values.yaml#L88">falco-exporter.prometheusRules.enabled</a></td>
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
			<td id="falco-exporter--resources--limits--cpu"><a href="./values.yaml#L92">falco-exporter.resources.limits.cpu</a></td>
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
			<td id="falco-exporter--resources--limits--memory"><a href="./values.yaml#L93">falco-exporter.resources.limits.memory</a></td>
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
			<td id="falco-exporter--resources--requests--cpu"><a href="./values.yaml#L95">falco-exporter.resources.requests.cpu</a></td>
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
			<td id="falco-exporter--resources--requests--memory"><a href="./values.yaml#L96">falco-exporter.resources.requests.memory</a></td>
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
			<td id="falco-exporter--serviceMonitor--enabled"><a href="./values.yaml#L77">falco-exporter.serviceMonitor.enabled</a></td>
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
			<td id="falco-exporter--serviceMonitor--labels--release"><a href="./values.yaml#L80">falco-exporter.serviceMonitor.labels.release</a></td>
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
			<td id="falco--collectors--crio--enabled"><a href="./values.yaml#L18">falco.collectors.crio.enabled</a></td>
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
			<td id="falco--collectors--docker--enabled"><a href="./values.yaml#L16">falco.collectors.docker.enabled</a></td>
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
			<td id="falco--driver--kind"><a href="./values.yaml#L21">falco.driver.kind</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ebpf"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--enabled"><a href="./values.yaml#L9">falco.enabled</a></td>
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
			<td id="falco--falco--grpc--enabled"><a href="./values.yaml#L33">falco.falco.grpc.enabled</a></td>
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
			<td id="falco--falco--grpc_output--enabled"><a href="./values.yaml#L36">falco.falco.grpc_output.enabled</a></td>
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
			<td id="falco--falco--json_output"><a href="./values.yaml#L30">falco.falco.json_output</a></td>
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
			<td id="falco--falco--metrics--enabled"><a href="./values.yaml#L39">falco.falco.metrics.enabled</a></td>
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
			<td id="falco--falco--rules_file[0]"><a href="./values.yaml#L25">falco.falco.rules_file[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/etc/falco/falco_rules.yaml"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falco--rules_file[1]"><a href="./values.yaml#L26">falco.falco.rules_file[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/etc/falco/falco_rules.local.yaml"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falco--rules_file[2]"><a href="./values.yaml#L27">falco.falco.rules_file[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/etc/falco/rules.d"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falco--rules_file[3]"><a href="./values.yaml#L28">falco.falco.rules_file[3]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/etc/falco/k8s_audit_rules.yaml"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falcoctl--config--artifact--follow--refs[0]"><a href="./values.yaml#L49">falco.falcoctl.config.artifact.follow.refs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"falco-rules:2"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falcoctl--config--artifact--install--refs[0]"><a href="./values.yaml#L46">falco.falcoctl.config.artifact.install.refs[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"falco-rules:2"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falcosidekick--config--webhook--address"><a href="./values.yaml#L56">falco.falcosidekick.config.webhook.address</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://webhook-falco-eventsource-svc.argocd.svc.cluster.local:12000/falco"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="falco--falcosidekick--enabled"><a href="./values.yaml#L52">falco.falcosidekick.enabled</a></td>
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
			<td id="falco--falcosidekick--webui--enabled"><a href="./values.yaml#L59">falco.falcosidekick.webui.enabled</a></td>
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
			<td id="falco--falcosidekick--webui--redis--resources--limits--cpu"><a href="./values.yaml#L65">falco.falcosidekick.webui.redis.resources.limits.cpu</a></td>
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
			<td id="falco--falcosidekick--webui--redis--resources--limits--memory"><a href="./values.yaml#L66">falco.falcosidekick.webui.redis.resources.limits.memory</a></td>
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
			<td id="falco--falcosidekick--webui--redis--resources--requests--cpu"><a href="./values.yaml#L68">falco.falcosidekick.webui.redis.resources.requests.cpu</a></td>
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
			<td id="falco--falcosidekick--webui--redis--resources--requests--memory"><a href="./values.yaml#L69">falco.falcosidekick.webui.redis.resources.requests.memory</a></td>
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
			<td id="falco--falcosidekick--webui--replicaCount"><a href="./values.yaml#L60">falco.falcosidekick.webui.replicaCount</a></td>
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
			<td id="falco--scc--create"><a href="./values.yaml#L12">falco.scc.create</a></td>
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
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
