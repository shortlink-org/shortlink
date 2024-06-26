# kratos

![Version: 0.3.3](https://img.shields.io/badge/Version-0.3.3-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| batazor | <batazor111@gmail.com> | <batazor.ru> |

## Requirements

Kubernetes: `>= 1.29.0 || >= v1.29.0-0`

| Repository | Name | Version |
|------------|------|---------|
| https://k8s.ory.sh/helm/charts | kratos | 0.45.0 |

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
			<td id="kratos--deployment--extraEnv[0]--name"><a href="./values.yaml#L18">kratos.deployment.extraEnv[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TRACING_PROVIDER"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[0]--value"><a href="./values.yaml#L19">kratos.deployment.extraEnv[0].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"jaeger"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[1]--name"><a href="./values.yaml#L20">kratos.deployment.extraEnv[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TRACING_PROVIDERS_JAEGER_SAMPLING_SERVER_URL"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[1]--value"><a href="./values.yaml#L21">kratos.deployment.extraEnv[1].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://grafana-tempo.grafana:14268/sampling"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[2]--name"><a href="./values.yaml#L22">kratos.deployment.extraEnv[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"TRACING_PROVIDERS_JAEGER_LOCAL_AGENT_ADDRESS"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[2]--value"><a href="./values.yaml#L23">kratos.deployment.extraEnv[2].value</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"grafana-tempo.grafana:6831"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[3]--name"><a href="./values.yaml#L24">kratos.deployment.extraEnv[3].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"DSN"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[3]--valueFrom--secretKeyRef--key"><a href="./values.yaml#L28">kratos.deployment.extraEnv[3].valueFrom.secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"uri"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--deployment--extraEnv[3]--valueFrom--secretKeyRef--name"><a href="./values.yaml#L27">kratos.deployment.extraEnv[3].valueFrom.secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kratos-postgres-pguser-kratos"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--enabled"><a href="./values.yaml#L8">kratos.enabled</a></td>
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
			<td id="kratos--fullnameOverride"><a href="./values.yaml#L10">kratos.fullnameOverride</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"kratos"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--ingress--admin--className"><a href="./values.yaml#L33">kratos.ingress.admin.className</a></td>
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
			<td id="kratos--ingress--admin--enabled"><a href="./values.yaml#L32">kratos.ingress.admin.enabled</a></td>
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
			<td id="kratos--ingress--public--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L38">kratos.ingress.public.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="kratos--ingress--public--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L40">kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="kratos--ingress--public--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L39">kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="kratos--ingress--public--annotations--"nginx--ingress--kubernetes--io/rewrite-target""><a href="./values.yaml#L41">kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/rewrite-target"</a></td>
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
			<td id="kratos--ingress--public--annotations--"nginx--ingress--kubernetes--io/use-regex""><a href="./values.yaml#L42">kratos.ingress.public.annotations."nginx.ingress.kubernetes.io/use-regex"</a></td>
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
			<td id="kratos--ingress--public--className"><a href="./values.yaml#L36">kratos.ingress.public.className</a></td>
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
			<td id="kratos--ingress--public--enabled"><a href="./values.yaml#L35">kratos.ingress.public.enabled</a></td>
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
			<td id="kratos--ingress--public--hosts[0]--host"><a href="./values.yaml#L44">kratos.ingress.public.hosts[0].host</a></td>
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
			<td id="kratos--ingress--public--hosts[0]--paths[0]--path"><a href="./values.yaml#L46">kratos.ingress.public.hosts[0].paths[0].path</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"/api/auth/?(.*)"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--ingress--public--hosts[0]--paths[0]--pathType"><a href="./values.yaml#L47">kratos.ingress.public.hosts[0].paths[0].pathType</a></td>
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
			<td id="kratos--ingress--public--tls[0]--hosts[0]"><a href="./values.yaml#L51">kratos.ingress.public.tls[0].hosts[0]</a></td>
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
			<td id="kratos--ingress--public--tls[0]--secretName"><a href="./values.yaml#L49">kratos.ingress.public.tls[0].secretName</a></td>
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
			<td id="kratos--kratos--automigration"><a href="./values.yaml#L181">kratos.kratos.automigration</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "enabled": true,
  "resources": {
    "limits": {
      "cpu": "100m",
      "memory": "128Mi"
    },
    "requests": {
      "cpu": "100m",
      "memory": "128Mi"
    }
  },
  "type": "initContainer"
}
</pre>
</div>
			</td>
			<td>Enables database migration</td>
		</tr>
		<tr>
			<td id="kratos--kratos--automigration--type"><a href="./values.yaml#L187">kratos.kratos.automigration.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"initContainer"
</pre>
</div>
			</td>
			<td>Configure the way to execute database migration. Possible values: job, initContainer When set to job, the migration will be executed as a job on release or upgrade. When set to initContainer, the migration will be executed when kratos pod is created Defaults to job</td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--hashers--argon2--iterations"><a href="./values.yaml#L174">kratos.kratos.config.hashers.argon2.iterations</a></td>
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
			<td id="kratos--kratos--config--hashers--argon2--key_length"><a href="./values.yaml#L176">kratos.kratos.config.hashers.argon2.key_length</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
16
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--hashers--argon2--memory"><a href="./values.yaml#L173">kratos.kratos.config.hashers.argon2.memory</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"128MB"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--hashers--argon2--parallelism"><a href="./values.yaml#L172">kratos.kratos.config.hashers.argon2.parallelism</a></td>
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
			<td id="kratos--kratos--config--hashers--argon2--salt_length"><a href="./values.yaml#L175">kratos.kratos.config.hashers.argon2.salt_length</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
16
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--identity--default_schema_id"><a href="./values.yaml#L56">kratos.kratos.config.identity.default_schema_id</a></td>
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
			<td id="kratos--kratos--config--identity--schemas[0]--id"><a href="./values.yaml#L58">kratos.kratos.config.identity.schemas[0].id</a></td>
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
			<td id="kratos--kratos--config--identity--schemas[0]--url"><a href="./values.yaml#L59">kratos.kratos.config.identity.schemas[0].url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"file:///etc/config/identity.default.schema.json"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--log--format"><a href="./values.yaml#L162">kratos.kratos.config.log.format</a></td>
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
			<td id="kratos--kratos--config--log--leak_sensitive_values"><a href="./values.yaml#L163">kratos.kratos.config.log.leak_sensitive_values</a></td>
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
			<td id="kratos--kratos--config--log--level"><a href="./values.yaml#L161">kratos.kratos.config.log.level</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"info"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--secrets--cookie[0]"><a href="./values.yaml#L168">kratos.kratos.config.secrets.cookie[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"PLEASE-CHANGE-ME-I-AM-VERY-INSECURE"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--allowed_return_urls[0]"><a href="./values.yaml#L96">kratos.kratos.config.selfservice.allowed_return_urls[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"*"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--allowed_return_urls[1]"><a href="./values.yaml#L97">kratos.kratos.config.selfservice.allowed_return_urls[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://*"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--allowed_return_urls[2]"><a href="./values.yaml#L98">kratos.kratos.config.selfservice.allowed_return_urls[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://*"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--default_browser_return_url"><a href="./values.yaml#L94">kratos.kratos.config.selfservice.default_browser_return_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--error--ui_url"><a href="./values.yaml#L122">kratos.kratos.config.selfservice.flows.error.ui_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/error"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--login--lifespan"><a href="./values.yaml#L147">kratos.kratos.config.selfservice.flows.login.lifespan</a></td>
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
			<td id="kratos--kratos--config--selfservice--flows--login--ui_url"><a href="./values.yaml#L146">kratos.kratos.config.selfservice.flows.login.ui_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/auth/login"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--logout--after--default_browser_return_url"><a href="./values.yaml#L143">kratos.kratos.config.selfservice.flows.logout.after.default_browser_return_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/auth/login"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--recovery--enabled"><a href="./values.yaml#L132">kratos.kratos.config.selfservice.flows.recovery.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--flows--recovery--ui_url"><a href="./values.yaml#L133">kratos.kratos.config.selfservice.flows.recovery.ui_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/auth/recovery"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--registration--after--oidc--hooks[0]--hook"><a href="./values.yaml#L158">kratos.kratos.config.selfservice.flows.registration.after.oidc.hooks[0].hook</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"session"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--registration--after--password--hooks[0]--hook"><a href="./values.yaml#L155">kratos.kratos.config.selfservice.flows.registration.after.password.hooks[0].hook</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"session"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--registration--lifespan"><a href="./values.yaml#L150">kratos.kratos.config.selfservice.flows.registration.lifespan</a></td>
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
			<td id="kratos--kratos--config--selfservice--flows--registration--ui_url"><a href="./values.yaml#L151">kratos.kratos.config.selfservice.flows.registration.ui_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/auth/registration"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--settings--privileged_session_max_age"><a href="./values.yaml#L128">kratos.kratos.config.selfservice.flows.settings.privileged_session_max_age</a></td>
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
			<td id="kratos--kratos--config--selfservice--flows--settings--required_aal"><a href="./values.yaml#L129">kratos.kratos.config.selfservice.flows.settings.required_aal</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"highest_available"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--settings--ui_url"><a href="./values.yaml#L125">kratos.kratos.config.selfservice.flows.settings.ui_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/user/profile"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--verification--after--default_browser_return_url"><a href="./values.yaml#L139">kratos.kratos.config.selfservice.flows.verification.after.default_browser_return_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--flows--verification--enabled"><a href="./values.yaml#L136">kratos.kratos.config.selfservice.flows.verification.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--flows--verification--ui_url"><a href="./values.yaml#L137">kratos.kratos.config.selfservice.flows.verification.ui_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/auth/verification"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--selfservice--methods--link--enabled"><a href="./values.yaml#L106">kratos.kratos.config.selfservice.methods.link.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--methods--lookup_secret--enabled"><a href="./values.yaml#L110">kratos.kratos.config.selfservice.methods.lookup_secret.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--methods--oidc--enabled"><a href="./values.yaml#L108">kratos.kratos.config.selfservice.methods.oidc.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--methods--password--enabled"><a href="./values.yaml#L102">kratos.kratos.config.selfservice.methods.password.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--methods--profile--enabled"><a href="./values.yaml#L104">kratos.kratos.config.selfservice.methods.profile.enabled</a></td>
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
			<td id="kratos--kratos--config--selfservice--methods--totp--config--issuer"><a href="./values.yaml#L115">kratos.kratos.config.selfservice.methods.totp.config.issuer</a></td>
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
			<td id="kratos--kratos--config--selfservice--methods--totp--enabled"><a href="./values.yaml#L112">kratos.kratos.config.selfservice.methods.totp.enabled</a></td>
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
			<td id="kratos--kratos--config--serve--admin--base_url"><a href="./values.yaml#L84">kratos.kratos.config.serve.admin.base_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://127.0.0.1:4434/"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--base_url"><a href="./values.yaml#L64">kratos.kratos.config.serve.public.base_url</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/api/auth"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allow_credentials"><a href="./values.yaml#L82">kratos.kratos.config.serve.public.cors.allow_credentials</a></td>
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
			<td id="kratos--kratos--config--serve--public--cors--allowed_headers[0]"><a href="./values.yaml#L78">kratos.kratos.config.serve.public.cors.allowed_headers[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Authorization"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_headers[1]"><a href="./values.yaml#L79">kratos.kratos.config.serve.public.cors.allowed_headers[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Cookie"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_headers[2]"><a href="./values.yaml#L80">kratos.kratos.config.serve.public.cors.allowed_headers[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Content-Type"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_headers[3]"><a href="./values.yaml#L81">kratos.kratos.config.serve.public.cors.allowed_headers[3]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Set-Cookie"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_methods[0]"><a href="./values.yaml#L72">kratos.kratos.config.serve.public.cors.allowed_methods[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POST"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_methods[1]"><a href="./values.yaml#L73">kratos.kratos.config.serve.public.cors.allowed_methods[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"GET"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_methods[2]"><a href="./values.yaml#L74">kratos.kratos.config.serve.public.cors.allowed_methods[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"PUT"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_methods[3]"><a href="./values.yaml#L75">kratos.kratos.config.serve.public.cors.allowed_methods[3]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"PATCH"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_methods[4]"><a href="./values.yaml#L76">kratos.kratos.config.serve.public.cors.allowed_methods[4]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"DELETE"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_origins[0]"><a href="./values.yaml#L69">kratos.kratos.config.serve.public.cors.allowed_origins[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"http://127.0.0.1:3000"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--allowed_origins[1]"><a href="./values.yaml#L70">kratos.kratos.config.serve.public.cors.allowed_origins[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--serve--public--cors--debug"><a href="./values.yaml#L67">kratos.kratos.config.serve.public.cors.debug</a></td>
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
			<td id="kratos--kratos--config--serve--public--cors--enabled"><a href="./values.yaml#L66">kratos.kratos.config.serve.public.cors.enabled</a></td>
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
			<td id="kratos--kratos--config--session--cookie--domain"><a href="./values.yaml#L90">kratos.kratos.config.session.cookie.domain</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--session--cookie--same_site"><a href="./values.yaml#L91">kratos.kratos.config.session.cookie.same_site</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Lax"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--config--session--lifespan"><a href="./values.yaml#L88">kratos.kratos.config.session.lifespan</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"720h"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--development"><a href="./values.yaml#L178">kratos.kratos.development</a></td>
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
			<td id="kratos--kratos--identitySchemas--"identity--default--schema--json""><a href="./values.yaml#L199">kratos.kratos.identitySchemas."identity.default.schema.json"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"{\n  \"$id\": \"https://schemas.ory.sh/presets/kratos/quickstart/email-password/identity.schema.json\",\n  \"$schema\": \"http://json-schema.org/draft-07/schema#\",\n  \"title\": \"Person\",\n  \"type\": \"object\",\n  \"properties\": {\n    \"traits\": {\n      \"type\": \"object\",\n      \"properties\": {\n        \"email\": {\n          \"type\": \"string\",\n          \"format\": \"email\",\n          \"title\": \"E-Mail\",\n          \"minLength\": 3,\n          \"ory.sh/kratos\": {\n            \"credentials\": {\n              \"password\": {\n                \"identifier\": true\n              },\n              \"totp\": {\n                \"account_name\": true\n              }\n            },\n            \"verification\": {\n              \"via\": \"email\"\n            },\n            \"recovery\": {\n              \"via\": \"email\"\n            }\n          }\n        },\n        \"name\": {\n          \"type\": \"object\",\n          \"properties\": {\n            \"first\": {\n              \"title\": \"First Name\",\n              \"type\": \"string\"\n            },\n            \"last\": {\n              \"title\": \"Last Name\",\n              \"type\": \"string\"\n            }\n          }\n        }\n      },\n      \"required\": [\n        \"email\"\n      ],\n      \"additionalProperties\": false\n    }\n  }\n}\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--identitySchemas--"oidc--github--jsonnet""><a href="./values.yaml#L252">kratos.kratos.identitySchemas."oidc.github.jsonnet"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local claims = {\n  email_verified: false,\n} + std.extVar('claims');\n\n{\n  identity: {\n    traits: {\n      // Allowing unverified email addresses enables account\n      // enumeration attacks, especially if the value is used for\n      // e.g. verification or as a password login identifier.\n      //\n      // Therefore we only return the email if it (a) exists and (b) is marked verified\n      // by GitHub.\n      [if 'email' in claims \u0026\u0026 claims.email_verified then 'email' else null]: claims.email,\n    },\n    metadata_public: {\n      github_username: claims.username,\n    }\n  },\n}\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--kratos--identitySchemas--"oidc--gitlab--jsonnet""><a href="./values.yaml#L273">kratos.kratos.identitySchemas."oidc.gitlab.jsonnet"</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"local claims = {\n  email_verified: false,\n} + std.extVar('claims');\n{\n  identity: {\n    traits: {\n      // Allowing unverified email addresses enables account\n      // enumeration attacks,  if the value is used for\n      // verification or as a password login identifier.\n      //\n      // Therefore we only return the email if it (a) exists and (b) is marked verified\n      // by GitLab.\n      [if 'email' in claims \u0026\u0026 claims.email_verified then 'email' else null]: claims.email,\n    },\n  },\n}\n"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="kratos--secret--enabled"><a href="./values.yaml#L13">kratos.secret.enabled</a></td>
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
			<td id="kratos--secret--hashSumEnabled"><a href="./values.yaml#L14">kratos.secret.hashSumEnabled</a></td>
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
			<td id="kratos--serviceMonitor--enabled"><a href="./values.yaml#L293">kratos.serviceMonitor.enabled</a></td>
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
			<td id="kratos--serviceMonitor--labels--release"><a href="./values.yaml#L296">kratos.serviceMonitor.labels.release</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.12.0](https://github.com/norwoodj/helm-docs/releases/v1.12.0)
