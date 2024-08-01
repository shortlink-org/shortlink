# redis

![Version: 0.2.0](https://img.shields.io/badge/Version-0.2.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://vmware-tanzu.github.io/helm-charts | velero | 7.1.4 |

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
			<td id="velero--configuration--volumeSnapshotLocation[0]--name"><a href="./values.yaml#L28">velero.configuration.volumeSnapshotLocation[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"default"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--configuration--volumeSnapshotLocation[0]--provider"><a href="./values.yaml#L29">velero.configuration.volumeSnapshotLocation[0].provider</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"aws"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--containerSecurityContext--allowPrivilegeEscalation"><a href="./values.yaml#L20">velero.containerSecurityContext.allowPrivilegeEscalation</a></td>
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
			<td id="velero--containerSecurityContext--capabilities--add"><a href="./values.yaml#L23">velero.containerSecurityContext.capabilities.add</a></td>
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
			<td id="velero--containerSecurityContext--capabilities--drop[0]"><a href="./values.yaml#L22">velero.containerSecurityContext.capabilities.drop[0]</a></td>
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
			<td id="velero--containerSecurityContext--readOnlyRootFilesystem"><a href="./values.yaml#L24">velero.containerSecurityContext.readOnlyRootFilesystem</a></td>
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
			<td id="velero--enabled"><a href="./values.yaml#L6">velero.enabled</a></td>
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
			<td id="velero--initContainers[0]--image"><a href="./values.yaml#L10">velero.initContainers[0].image</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"velero/velero-plugin-for-csi:v0.7.1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--initContainers[0]--imagePullPolicy"><a href="./values.yaml#L11">velero.initContainers[0].imagePullPolicy</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"IfNotPresent"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--initContainers[0]--name"><a href="./values.yaml#L9">velero.initContainers[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"velero-plugin-for-csi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--initContainers[0]--volumeMounts[0]--mountPath"><a href="./values.yaml#L13">velero.initContainers[0].volumeMounts[0].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/target"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--initContainers[0]--volumeMounts[0]--name"><a href="./values.yaml#L14">velero.initContainers[0].volumeMounts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"plugins"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--nodeAgentPodMonitor--enabled"><a href="./values.yaml#L35">velero.metrics.nodeAgentPodMonitor.enabled</a></td>
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
			<td id="velero--metrics--prometheusRule--enabled"><a href="./values.yaml#L37">velero.metrics.prometheusRule.enabled</a></td>
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
			<td id="velero--metrics--prometheusRule--spec[0]--alert"><a href="./values.yaml#L39">velero.metrics.prometheusRule.spec[0].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"VeleroBackupPartialFailures"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--prometheusRule--spec[0]--annotations--message"><a href="./values.yaml#L41">velero.metrics.prometheusRule.spec[0].annotations.message</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Velero backup {{ $labels.schedule }} has {{ $value | humanizePercentage }} partialy failed backups."
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--prometheusRule--spec[0]--expr"><a href="./values.yaml#L42">velero.metrics.prometheusRule.spec[0].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"velero_backup_partial_failure_total{schedule!=\"\"} / velero_backup_attempt_total{schedule!=\"\"} \u003e 0.25"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--prometheusRule--spec[0]--for"><a href="./values.yaml#L44">velero.metrics.prometheusRule.spec[0].for</a></td>
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
			<td id="velero--metrics--prometheusRule--spec[0]--labels--severity"><a href="./values.yaml#L46">velero.metrics.prometheusRule.spec[0].labels.severity</a></td>
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
			<td id="velero--metrics--prometheusRule--spec[1]--alert"><a href="./values.yaml#L47">velero.metrics.prometheusRule.spec[1].alert</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"VeleroBackupFailures"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--prometheusRule--spec[1]--annotations--message"><a href="./values.yaml#L49">velero.metrics.prometheusRule.spec[1].annotations.message</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Velero backup {{ $labels.schedule }} has {{ $value | humanizePercentage }} failed backups."
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--prometheusRule--spec[1]--expr"><a href="./values.yaml#L50">velero.metrics.prometheusRule.spec[1].expr</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"velero_backup_failure_total{schedule!=\"\"} / velero_backup_attempt_total{schedule!=\"\"} \u003e 0.25"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="velero--metrics--prometheusRule--spec[1]--for"><a href="./values.yaml#L52">velero.metrics.prometheusRule.spec[1].for</a></td>
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
			<td id="velero--metrics--prometheusRule--spec[1]--labels--severity"><a href="./values.yaml#L54">velero.metrics.prometheusRule.spec[1].labels.severity</a></td>
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
			<td id="velero--metrics--serviceMonitor--enabled"><a href="./values.yaml#L33">velero.metrics.serviceMonitor.enabled</a></td>
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
			<td id="velero--podSecurityContext--fsGroup"><a href="./values.yaml#L17">velero.podSecurityContext.fsGroup</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
1337
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
