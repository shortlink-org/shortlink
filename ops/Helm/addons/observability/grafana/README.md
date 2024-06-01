# grafana

![Version: 0.7.3](https://img.shields.io/badge/Version-0.7.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://grafana.github.io/helm-charts | grafana | 7.3.11 |

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
			<td id="grafana--"grafana--ini"--auth--anonymous--enabled"><a href="./values.yaml#L194">grafana."grafana.ini".auth.anonymous.enabled</a></td>
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
			<td id="grafana--"grafana--ini"--auth--anonymous--hide_version"><a href="./values.yaml#L203">grafana."grafana.ini".auth.anonymous.hide_version</a></td>
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
			<td id="grafana--"grafana--ini"--auth--anonymous--org_name"><a href="./values.yaml#L197">grafana."grafana.ini".auth.anonymous.org_name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Main Org."
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--"grafana--ini"--auth--anonymous--org_role"><a href="./values.yaml#L200">grafana."grafana.ini".auth.anonymous.org_role</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Viewer"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--"grafana--ini"--reporting--enabled"><a href="./values.yaml#L190">grafana."grafana.ini".reporting.enabled</a></td>
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
			<td id="grafana--assertNoLeakedSecrets"><a href="./values.yaml#L18">grafana.assertNoLeakedSecrets</a></td>
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
			<td id="grafana--defaultDashboardsEnabled"><a href="./values.yaml#L22">grafana.defaultDashboardsEnabled</a></td>
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
			<td id="grafana--downloadDashboards--resources--limits--cpu"><a href="./values.yaml#L173">grafana.downloadDashboards.resources.limits.cpu</a></td>
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
			<td id="grafana--downloadDashboards--resources--limits--memory"><a href="./values.yaml#L174">grafana.downloadDashboards.resources.limits.memory</a></td>
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
			<td id="grafana--downloadDashboards--resources--requests--cpu"><a href="./values.yaml#L176">grafana.downloadDashboards.resources.requests.cpu</a></td>
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
			<td id="grafana--downloadDashboards--resources--requests--memory"><a href="./values.yaml#L177">grafana.downloadDashboards.resources.requests.memory</a></td>
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
			<td id="grafana--enabled"><a href="./values.yaml#L7">grafana.enabled</a></td>
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
			<td id="grafana--env--GF_DIAGNOSTICS_PROFILING_ADDR"><a href="./values.yaml#L95">grafana.env.GF_DIAGNOSTICS_PROFILING_ADDR</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.0.0.0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--env--GF_DIAGNOSTICS_PROFILING_ENABLED"><a href="./values.yaml#L94">grafana.env.GF_DIAGNOSTICS_PROFILING_ENABLED</a></td>
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
			<td id="grafana--env--GF_DIAGNOSTICS_PROFILING_PORT"><a href="./values.yaml#L96">grafana.env.GF_DIAGNOSTICS_PROFILING_PORT</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
6060
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--env--GF_FEATURE_TOGGLES_ENABLE"><a href="./values.yaml#L93">grafana.env.GF_FEATURE_TOGGLES_ENABLE</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"flameGraph,traceToProfiles,accessControlOnCall"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraContainerVolumes[0]--emptyDir"><a href="./values.yaml#L164">grafana.extraContainerVolumes[0].emptyDir</a></td>
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
			<td id="grafana--extraContainerVolumes[0]--name"><a href="./values.yaml#L163">grafana.extraContainerVolumes[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-plugins"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraInitContainers[0]--args[0]"><a href="./values.yaml#L140">grafana.extraInitContainers[0].args[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"# Get pluginName zip\ncurl -L \\\n-H \"Accept: application/octet-stream\" \\\n-o /tmp/grafana-lokiexplore-app-latest.zip \\\nhttps://storage.googleapis.com/integration-artifacts/grafana-lokiexplore-app/grafana-lokiexplore-app-latest.zip --insecure\n\n# Unzip pluginName where grafana read plugins\nunzip /tmp/grafana-lokiexplore-app-latest.zip -d /var/lib/grafana/plugins\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraInitContainers[0]--command[0]"><a href="./values.yaml#L138">grafana.extraInitContainers[0].command[0]</a></td>
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
			<td id="grafana--extraInitContainers[0]--command[1]"><a href="./values.yaml#L138">grafana.extraInitContainers[0].command[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"-c"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraInitContainers[0]--image"><a href="./values.yaml#L137">grafana.extraInitContainers[0].image</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"quay.io/curl/curl:latest"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraInitContainers[0]--name"><a href="./values.yaml#L136">grafana.extraInitContainers[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"plugin-init"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraInitContainers[0]--resources--limits--cpu"><a href="./values.yaml#L152">grafana.extraInitContainers[0].resources.limits.cpu</a></td>
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
			<td id="grafana--extraInitContainers[0]--resources--limits--memory"><a href="./values.yaml#L153">grafana.extraInitContainers[0].resources.limits.memory</a></td>
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
			<td id="grafana--extraInitContainers[0]--resources--requests--cpu"><a href="./values.yaml#L155">grafana.extraInitContainers[0].resources.requests.cpu</a></td>
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
			<td id="grafana--extraInitContainers[0]--resources--requests--memory"><a href="./values.yaml#L156">grafana.extraInitContainers[0].resources.requests.memory</a></td>
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
			<td id="grafana--extraInitContainers[0]--volumeMounts[0]--mountPath"><a href="./values.yaml#L160">grafana.extraInitContainers[0].volumeMounts[0].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/var/lib/grafana/plugins"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraInitContainers[0]--volumeMounts[0]--name"><a href="./values.yaml#L159">grafana.extraInitContainers[0].volumeMounts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-plugins"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraVolumeMounts[0]--mountPath"><a href="./values.yaml#L168">grafana.extraVolumeMounts[0].mountPath</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/var/lib/grafana/plugins"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--extraVolumeMounts[0]--name"><a href="./values.yaml#L167">grafana.extraVolumeMounts[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-plugins"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--imageRenderer--enabled"><a href="./values.yaml#L51">grafana.imageRenderer.enabled</a></td>
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
			<td id="grafana--imageRenderer--resources--limits--cpu"><a href="./values.yaml#L68">grafana.imageRenderer.resources.limits.cpu</a></td>
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
			<td id="grafana--imageRenderer--resources--limits--memory"><a href="./values.yaml#L69">grafana.imageRenderer.resources.limits.memory</a></td>
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
			<td id="grafana--imageRenderer--resources--requests--cpu"><a href="./values.yaml#L71">grafana.imageRenderer.resources.requests.cpu</a></td>
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
			<td id="grafana--imageRenderer--resources--requests--memory"><a href="./values.yaml#L72">grafana.imageRenderer.resources.requests.memory</a></td>
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
			<td id="grafana--imageRenderer--revisionHistoryLimit"><a href="./values.yaml#L53">grafana.imageRenderer.revisionHistoryLimit</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
2
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--imageRenderer--securityContext--fsGroup"><a href="./values.yaml#L59">grafana.imageRenderer.securityContext.fsGroup</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
472
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--imageRenderer--securityContext--runAsGroup"><a href="./values.yaml#L58">grafana.imageRenderer.securityContext.runAsGroup</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
472
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--imageRenderer--securityContext--runAsNonRoot"><a href="./values.yaml#L56">grafana.imageRenderer.securityContext.runAsNonRoot</a></td>
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
			<td id="grafana--imageRenderer--securityContext--runAsUser"><a href="./values.yaml#L57">grafana.imageRenderer.securityContext.runAsUser</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
472
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--imageRenderer--serviceMonitor--enabled"><a href="./values.yaml#L62">grafana.imageRenderer.serviceMonitor.enabled</a></td>
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
			<td id="grafana--imageRenderer--serviceMonitor--interval"><a href="./values.yaml#L64">grafana.imageRenderer.serviceMonitor.interval</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--imageRenderer--serviceMonitor--selfMonitor"><a href="./values.yaml#L63">grafana.imageRenderer.serviceMonitor.selfMonitor</a></td>
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
			<td id="grafana--ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L78">grafana.ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="grafana--ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L80">grafana.ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="grafana--ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L79">grafana.ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="grafana--ingress--enabled"><a href="./values.yaml#L75">grafana.ingress.enabled</a></td>
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
			<td id="grafana--ingress--hosts[0]"><a href="./values.yaml#L83">grafana.ingress.hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--ingress--path"><a href="./values.yaml#L85">grafana.ingress.path</a></td>
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
			<td id="grafana--ingress--tls[0]--hosts[0]"><a href="./values.yaml#L90">grafana.ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--ingress--tls[0]--secretName"><a href="./values.yaml#L88">grafana.ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-ingress-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--initChownData--resources--limits--cpu"><a href="./values.yaml#L182">grafana.initChownData.resources.limits.cpu</a></td>
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
			<td id="grafana--initChownData--resources--limits--memory"><a href="./values.yaml#L183">grafana.initChownData.resources.limits.memory</a></td>
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
			<td id="grafana--initChownData--resources--requests--cpu"><a href="./values.yaml#L185">grafana.initChownData.resources.requests.cpu</a></td>
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
			<td id="grafana--initChownData--resources--requests--memory"><a href="./values.yaml#L186">grafana.initChownData.resources.requests.memory</a></td>
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
			<td id="grafana--namespaceOverride"><a href="./values.yaml#L8">grafana.namespaceOverride</a></td>
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
			<td id="grafana--persistence--enabled"><a href="./values.yaml#L35">grafana.persistence.enabled</a></td>
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
			<td id="grafana--persistence--inMemory--enabled"><a href="./values.yaml#L38">grafana.persistence.inMemory.enabled</a></td>
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
			<td id="grafana--persistence--storageClassName"><a href="./values.yaml#L36">grafana.persistence.storageClassName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local-path"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--plugins[0]"><a href="./values.yaml#L130">grafana.plugins[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-polystat-panel"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--plugins[1]"><a href="./values.yaml#L131">grafana.plugins[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-oncall-app"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--plugins[2]"><a href="./values.yaml#L132">grafana.plugins[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"cloudflare-app"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--plugins[3]"><a href="./values.yaml#L133">grafana.plugins[3]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"hamedkarbasi93-kafka-datasource"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--podAnnotations--"pyroscope--grafana--com/port""><a href="./values.yaml#L12">grafana.podAnnotations."pyroscope.grafana.com/port"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"6060"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--podAnnotations--"pyroscope--grafana--com/scrape""><a href="./values.yaml#L11">grafana.podAnnotations."pyroscope.grafana.com/scrape"</a></td>
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
			<td id="grafana--resources--limits--cpu"><a href="./values.yaml#L28">grafana.resources.limits.cpu</a></td>
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
			<td id="grafana--resources--limits--memory"><a href="./values.yaml#L29">grafana.resources.limits.memory</a></td>
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
			<td id="grafana--resources--requests--cpu"><a href="./values.yaml#L31">grafana.resources.requests.cpu</a></td>
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
			<td id="grafana--resources--requests--memory"><a href="./values.yaml#L32">grafana.resources.requests.memory</a></td>
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
			<td id="grafana--revisionHistoryLimit"><a href="./values.yaml#L40">grafana.revisionHistoryLimit</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
2
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--serviceMonitor--enabled"><a href="./values.yaml#L43">grafana.serviceMonitor.enabled</a></td>
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
			<td id="grafana--serviceMonitor--interval"><a href="./values.yaml#L45">grafana.serviceMonitor.interval</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"1m"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--serviceMonitor--labels--release"><a href="./values.yaml#L48">grafana.serviceMonitor.labels.release</a></td>
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
			<td id="grafana--serviceMonitor--selfMonitor"><a href="./values.yaml#L44">grafana.serviceMonitor.selfMonitor</a></td>
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
			<td id="grafana--sidecar--alerts--enabled"><a href="./values.yaml#L100">grafana.sidecar.alerts.enabled</a></td>
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
			<td id="grafana--sidecar--alerts--searchNamespace"><a href="./values.yaml#L101">grafana.sidecar.alerts.searchNamespace</a></td>
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
			<td id="grafana--sidecar--dashboards--defaultFolderName"><a href="./values.yaml#L104">grafana.sidecar.dashboards.defaultFolderName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"General"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--sidecar--dashboards--enabled"><a href="./values.yaml#L103">grafana.sidecar.dashboards.enabled</a></td>
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
			<td id="grafana--sidecar--dashboards--folder"><a href="./values.yaml#L106">grafana.sidecar.dashboards.folder</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/tmp/dashboards"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="grafana--sidecar--dashboards--folderAnnotation"><a href="./values.yaml#L107">grafana.sidecar.dashboards.folderAnnotation</a></td>
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
			<td id="grafana--sidecar--dashboards--provider--foldersFromFilesStructure"><a href="./values.yaml#L109">grafana.sidecar.dashboards.provider.foldersFromFilesStructure</a></td>
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
			<td id="grafana--sidecar--dashboards--searchNamespace"><a href="./values.yaml#L105">grafana.sidecar.dashboards.searchNamespace</a></td>
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
			<td id="grafana--sidecar--datasources--enabled"><a href="./values.yaml#L111">grafana.sidecar.datasources.enabled</a></td>
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
			<td id="grafana--sidecar--datasources--searchNamespace"><a href="./values.yaml#L112">grafana.sidecar.datasources.searchNamespace</a></td>
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
			<td id="grafana--sidecar--notifiers--enabled"><a href="./values.yaml#L119">grafana.sidecar.notifiers.enabled</a></td>
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
			<td id="grafana--sidecar--notifiers--searchNamespace"><a href="./values.yaml#L120">grafana.sidecar.notifiers.searchNamespace</a></td>
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
			<td id="grafana--sidecar--plugins--enabled"><a href="./values.yaml#L114">grafana.sidecar.plugins.enabled</a></td>
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
			<td id="grafana--sidecar--plugins--initPlugins"><a href="./values.yaml#L117">grafana.sidecar.plugins.initPlugins</a></td>
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
			<td id="grafana--sidecar--plugins--searchNamespace"><a href="./values.yaml#L115">grafana.sidecar.plugins.searchNamespace</a></td>
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
			<td id="grafana--sidecar--plugins--skipReload"><a href="./values.yaml#L116">grafana.sidecar.plugins.skipReload</a></td>
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
			<td id="grafana--sidecar--resources--limits--cpu"><a href="./values.yaml#L123">grafana.sidecar.resources.limits.cpu</a></td>
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
			<td id="grafana--sidecar--resources--limits--memory"><a href="./values.yaml#L124">grafana.sidecar.resources.limits.memory</a></td>
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
			<td id="grafana--sidecar--resources--requests--cpu"><a href="./values.yaml#L126">grafana.sidecar.resources.requests.cpu</a></td>
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
			<td id="grafana--sidecar--resources--requests--memory"><a href="./values.yaml#L127">grafana.sidecar.resources.requests.memory</a></td>
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
			<td id="grafana--testFramework--enabled"><a href="./values.yaml#L15">grafana.testFramework.enabled</a></td>
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
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
