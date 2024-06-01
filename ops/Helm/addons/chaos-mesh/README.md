# chaos

![Version: 0.6.0](https://img.shields.io/badge/Version-0.6.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Chaos service

**Homepage:** <https://github.com/shortlink-org/shortlink>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Source Code

* <https://github.com/shortlink-org/shortlink>

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://charts.chaos-mesh.org | chaos-mesh | 2.6.3 |

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
			<td id="chaos-mesh--bpfki--create"><a href="./values.yaml#L73">chaos-mesh.bpfki.create</a></td>
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
			<td id="chaos-mesh--bpfki--resources--limits--cpu"><a href="./values.yaml#L77">chaos-mesh.bpfki.resources.limits.cpu</a></td>
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
			<td id="chaos-mesh--bpfki--resources--limits--memory"><a href="./values.yaml#L78">chaos-mesh.bpfki.resources.limits.memory</a></td>
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
			<td id="chaos-mesh--bpfki--resources--requests--cpu"><a href="./values.yaml#L80">chaos-mesh.bpfki.resources.requests.cpu</a></td>
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
			<td id="chaos-mesh--bpfki--resources--requests--memory"><a href="./values.yaml#L81">chaos-mesh.bpfki.resources.requests.memory</a></td>
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
			<td id="chaos-mesh--chaosDaemon--resources--limits--cpu"><a href="./values.yaml#L25">chaos-mesh.chaosDaemon.resources.limits.cpu</a></td>
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
			<td id="chaos-mesh--chaosDaemon--resources--limits--memory"><a href="./values.yaml#L26">chaos-mesh.chaosDaemon.resources.limits.memory</a></td>
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
			<td id="chaos-mesh--chaosDaemon--resources--requests--cpu"><a href="./values.yaml#L28">chaos-mesh.chaosDaemon.resources.requests.cpu</a></td>
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
			<td id="chaos-mesh--chaosDaemon--resources--requests--memory"><a href="./values.yaml#L29">chaos-mesh.chaosDaemon.resources.requests.memory</a></td>
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
			<td id="chaos-mesh--chaosDaemon--runtime"><a href="./values.yaml#L20">chaos-mesh.chaosDaemon.runtime</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"containerd"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--chaosDaemon--socketPath"><a href="./values.yaml#L21">chaos-mesh.chaosDaemon.socketPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/run/containerd/containerd.sock"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--chaosDlv--enable"><a href="./values.yaml#L84">chaos-mesh.chaosDlv.enable</a></td>
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
			<td id="chaos-mesh--controllerManager--enableFilterNamespace"><a href="./values.yaml#L9">chaos-mesh.controllerManager.enableFilterNamespace</a></td>
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
			<td id="chaos-mesh--controllerManager--replicaCount"><a href="./values.yaml#L7">chaos-mesh.controllerManager.replicaCount</a></td>
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
			<td id="chaos-mesh--controllerManager--resources--limits--cpu"><a href="./values.yaml#L13">chaos-mesh.controllerManager.resources.limits.cpu</a></td>
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
			<td id="chaos-mesh--controllerManager--resources--limits--memory"><a href="./values.yaml#L14">chaos-mesh.controllerManager.resources.limits.memory</a></td>
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
			<td id="chaos-mesh--controllerManager--resources--requests--cpu"><a href="./values.yaml#L16">chaos-mesh.controllerManager.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"25m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--controllerManager--resources--requests--memory"><a href="./values.yaml#L17">chaos-mesh.controllerManager.resources.requests.memory</a></td>
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
			<td id="chaos-mesh--dashboard--ingress--enabled"><a href="./values.yaml#L41">chaos-mesh.dashboard.ingress.enabled</a></td>
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
			<td id="chaos-mesh--dashboard--ingress--hosts[0]--name"><a href="./values.yaml#L46">chaos-mesh.dashboard.ingress.hosts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"dashboard.local"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--dashboard--ingress--hosts[0]--tls"><a href="./values.yaml#L47">chaos-mesh.dashboard.ingress.hosts[0].tls</a></td>
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
			<td id="chaos-mesh--dashboard--ingress--hosts[0]--tlsSecret"><a href="./values.yaml#L48">chaos-mesh.dashboard.ingress.hosts[0].tlsSecret</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"dashboard.local-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--dashboard--ingress--ingressClassName"><a href="./values.yaml#L43">chaos-mesh.dashboard.ingress.ingressClassName</a></td>
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
			<td id="chaos-mesh--dashboard--ingress--paths[0]"><a href="./values.yaml#L51">chaos-mesh.dashboard.ingress.paths[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/chaos"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--dashboard--resources--limits--cpu"><a href="./values.yaml#L55">chaos-mesh.dashboard.resources.limits.cpu</a></td>
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
			<td id="chaos-mesh--dashboard--resources--limits--memory"><a href="./values.yaml#L56">chaos-mesh.dashboard.resources.limits.memory</a></td>
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
			<td id="chaos-mesh--dashboard--resources--requests--cpu"><a href="./values.yaml#L58">chaos-mesh.dashboard.resources.requests.cpu</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"25m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="chaos-mesh--dashboard--resources--requests--memory"><a href="./values.yaml#L59">chaos-mesh.dashboard.resources.requests.memory</a></td>
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
			<td id="chaos-mesh--dashboard--securityMode"><a href="./values.yaml#L32">chaos-mesh.dashboard.securityMode</a></td>
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
			<td id="chaos-mesh--dnsServer--create"><a href="./values.yaml#L62">chaos-mesh.dnsServer.create</a></td>
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
			<td id="chaos-mesh--dnsServer--resources--limits--cpu"><a href="./values.yaml#L66">chaos-mesh.dnsServer.resources.limits.cpu</a></td>
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
			<td id="chaos-mesh--dnsServer--resources--limits--memory"><a href="./values.yaml#L67">chaos-mesh.dnsServer.resources.limits.memory</a></td>
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
			<td id="chaos-mesh--dnsServer--resources--requests--cpu"><a href="./values.yaml#L69">chaos-mesh.dnsServer.resources.requests.cpu</a></td>
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
			<td id="chaos-mesh--dnsServer--resources--requests--memory"><a href="./values.yaml#L70">chaos-mesh.dnsServer.resources.requests.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"70Mi"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
