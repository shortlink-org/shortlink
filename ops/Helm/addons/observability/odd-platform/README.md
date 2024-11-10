# odd-platform

![Version: 0.1.2](https://img.shields.io/badge/Version-0.1.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.30.0 || >= v1.30.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://opendatadiscovery.github.io/charts | odd-platform | 0.1.10 |
| oci://registry-1.docker.io/bitnamicharts | postgresql | 16.1.2 |

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
			<td id="odd-platform--config--env[0]--name"><a href="./values.yaml#L6">odd-platform.config.env[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"OTEL_INSTRUMENTATION_COMMON_DB_STATEMENT_SANITIZER_ENABLED"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--config--env[0]--value"><a href="./values.yaml#L7">odd-platform.config.env[0].value</a></td>
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
			<td id="odd-platform--config--env[1]--name"><a href="./values.yaml#L8">odd-platform.config.env[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"SPRING_DATASOURCE_URL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--config--env[1]--value"><a href="./values.yaml#L9">odd-platform.config.env[1].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"jdbc:postgresql://odd-platform-postgresql:5432/odd-platform"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--config--env[2]--name"><a href="./values.yaml#L10">odd-platform.config.env[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"SPRING_DATASOURCE_USERNAME"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--config--env[2]--value"><a href="./values.yaml#L11">odd-platform.config.env[2].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"odd-platform"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--config--env[3]--name"><a href="./values.yaml#L12">odd-platform.config.env[3].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"SPRING_DATASOURCE_PASSWORD"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--config--env[3]--value"><a href="./values.yaml#L13">odd-platform.config.env[3].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"odd-platform"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--enabled"><a href="./values.yaml#L2">odd-platform.enabled</a></td>
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
			<td id="odd-platform--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L20">odd-platform.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="odd-platform--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L22">odd-platform.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="odd-platform--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L21">odd-platform.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="odd-platform--ingress--className"><a href="./values.yaml#L17">odd-platform.ingress.className</a></td>
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
			<td id="odd-platform--ingress--enabled"><a href="./values.yaml#L16">odd-platform.ingress.enabled</a></td>
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
			<td id="odd-platform--ingress--hosts[0]--host"><a href="./values.yaml#L25">odd-platform.ingress.hosts[0].host</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"odd.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--ingress--hosts[0]--paths[0]--path"><a href="./values.yaml#L27">odd-platform.ingress.hosts[0].paths[0].path</a></td>
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
			<td id="odd-platform--ingress--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L28">odd-platform.ingress.hosts[0].paths[0].pathType</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"ImplementationSpecific"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L33">odd-platform.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"odd.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--ingress--tls[0]--secretName"><a href="./values.yaml#L31">odd-platform.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"odd-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="odd-platform--resources--limits--cpu"><a href="./values.yaml#L37">odd-platform.resources.limits.cpu</a></td>
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
			<td id="odd-platform--resources--limits--memory"><a href="./values.yaml#L38">odd-platform.resources.limits.memory</a></td>
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
			<td id="odd-platform--resources--requests--cpu"><a href="./values.yaml#L40">odd-platform.resources.requests.cpu</a></td>
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
			<td id="odd-platform--resources--requests--memory"><a href="./values.yaml#L41">odd-platform.resources.requests.memory</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
