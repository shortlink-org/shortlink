# pyroscope

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://pyroscope-io.github.io/helm-chart | pyroscope | 0.2.92 |
| https://pyroscope-io.github.io/helm-chart | pyroscope-ebpf | 0.1.31 |

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
			<td id="pyroscope-ebpf--args[0]"><a href="./values.yaml#L35">pyroscope-ebpf.args[0]</a></td>
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
			<td id="pyroscope-ebpf--args[1]"><a href="./values.yaml#L36">pyroscope-ebpf.args[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"--application-name"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyroscope-ebpf--args[2]"><a href="./values.yaml#L37">pyroscope-ebpf.args[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"k8s.ebpf"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyroscope-ebpf--args[3]"><a href="./values.yaml#L38">pyroscope-ebpf.args[3]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"--server-address"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyroscope-ebpf--args[4]"><a href="./values.yaml#L39">pyroscope-ebpf.args[4]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://pyroscope-server:4040"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyroscope-ebpf--enabled"><a href="./values.yaml#L32">pyroscope-ebpf.enabled</a></td>
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
			<td id="pyroscope--enabled"><a href="./values.yaml#L2">pyroscope.enabled</a></td>
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
			<td id="pyroscope--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L8">pyroscope.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="pyroscope--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L10">pyroscope.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="pyroscope--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L9">pyroscope.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="pyroscope--ingress--className"><a href="./values.yaml#L6">pyroscope.ingress.className</a></td>
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
			<td id="pyroscope--ingress--enabled"><a href="./values.yaml#L5">pyroscope.ingress.enabled</a></td>
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
			<td id="pyroscope--ingress--hosts[0]--host"><a href="./values.yaml#L13">pyroscope.ingress.hosts[0].host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"pyroscope.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyroscope--ingress--hosts[0]--paths[0]--path"><a href="./values.yaml#L15">pyroscope.ingress.hosts[0].paths[0].path</a></td>
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
			<td id="pyroscope--ingress--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L16">pyroscope.ingress.hosts[0].paths[0].pathType</a></td>
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
			<td id="pyroscope--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L21">pyroscope.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"pyroscope.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="pyroscope--ingress--tls[0]--secretName"><a href="./values.yaml#L19">pyroscope.ingress.tls[0].secretName</a></td>
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
			<td id="pyroscope--resources--limits--cpu"><a href="./values.yaml#L25">pyroscope.resources.limits.cpu</a></td>
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
			<td id="pyroscope--resources--limits--memory"><a href="./values.yaml#L26">pyroscope.resources.limits.memory</a></td>
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
			<td id="pyroscope--resources--requests--cpu"><a href="./values.yaml#L28">pyroscope.resources.requests.cpu</a></td>
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
			<td id="pyroscope--resources--requests--memory"><a href="./values.yaml#L29">pyroscope.resources.requests.memory</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
