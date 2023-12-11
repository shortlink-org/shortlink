# shortlink-support

![Version: 0.1.1](https://img.shields.io/badge/Version-0.1.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

Shortlink support service

**Homepage:** <https://batazor.github.io/shortlink/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Source Code

* <https://github.com/shortlink-org/shortlink>

## Requirements

Kubernetes: `>= 1.28.0 || >= v1.28.0-0`

| Repository | Name | Version |
|------------|------|---------|
| file://../shortlink-template | shortlink-template | 0.7.0 |

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
			<td id="deploy--env--TRACER_URI"><a href="./values.yaml#L21">deploy.env.TRACER_URI</a></td>
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
			<td id="deploy--image--repository"><a href="./values.yaml#L24">deploy.image.repository</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"registry.gitlab.com/shortlink-org/shortlink/support"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--image--tag"><a href="./values.yaml#L25">deploy.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.16.101"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="monitoring--enabled"><a href="./values.yaml#L54">monitoring.enabled</a></td>
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
			<td id="podDisruptionBudget--enabled"><a href="./values.yaml#L59">podDisruptionBudget.enabled</a></td>
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
			<td id="service--ports[0]--name"><a href="./values.yaml#L42">service.ports[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--port"><a href="./values.yaml#L43">service.ports[0].port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
8080
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--protocol"><a href="./values.yaml#L44">service.ports[0].protocol</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TCP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--public"><a href="./values.yaml#L45">service.ports[0].public</a></td>
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
			<td id="service--ports[1]--name"><a href="./values.yaml#L46">service.ports[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"metrics"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[1]--port"><a href="./values.yaml#L47">service.ports[1].port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
9090
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[1]--protocol"><a href="./values.yaml#L48">service.ports[1].protocol</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TCP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[1]--public"><a href="./values.yaml#L49">service.ports[1].public</a></td>
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
			<td id="service--type"><a href="./values.yaml#L40">service.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ClusterIP"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="sidecars[0]--image--repository"><a href="./values.yaml#L35">sidecars[0].image.repository</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"registry.gitlab.com/shortlink-org/shortlink/support-proxy"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="sidecars[0]--image--tag"><a href="./values.yaml#L36">sidecars[0].image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.16.101"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="sidecars[0]--name"><a href="./values.yaml#L33">sidecars[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"proxy"
</pre>
</div>
			</td>
			<td></td>
		</tr>
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.11.0](https://github.com/norwoodj/helm-docs/releases/v1.11.0)
