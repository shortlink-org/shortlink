# admin

![Version: 0.2.1](https://img.shields.io/badge/Version-0.2.1-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.0.0](https://img.shields.io/badge/AppVersion-1.0.0-informational?style=flat-square)

ShortLink Shop Admin

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
| file://../../../shortlink-template | shortlink-template | 0.8.18 |

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
			<td id="deploy--env--DEBUG"><a href="./values.yaml#L56">deploy.env.DEBUG</a></td>
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
			<td id="deploy--env--LOGIN_URL"><a href="./values.yaml#L50">deploy.env.LOGIN_URL</a></td>
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
			<td id="deploy--env--ORY_SDK_URL"><a href="./values.yaml#L48">deploy.env.ORY_SDK_URL</a></td>
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
			<td id="deploy--env--ORY_UI_URL"><a href="./values.yaml#L49">deploy.env.ORY_UI_URL</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"https://shortlink.best/next/auth"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--env--REDIS_URL"><a href="./values.yaml#L53">deploy.env.REDIS_URL</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"redis://redis-master.shortlink-shop:6379/0"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[0]--name"><a href="./values.yaml#L59">deploy.envSecret[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_DB"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[0]--secretKeyRef--key"><a href="./values.yaml#L62">deploy.envSecret[0].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"dbname"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[0]--secretKeyRef--name"><a href="./values.yaml#L61">deploy.envSecret[0].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[1]--name"><a href="./values.yaml#L63">deploy.envSecret[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_USER"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[1]--secretKeyRef--key"><a href="./values.yaml#L66">deploy.envSecret[1].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"user"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[1]--secretKeyRef--name"><a href="./values.yaml#L65">deploy.envSecret[1].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[2]--name"><a href="./values.yaml#L67">deploy.envSecret[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_PASSWORD"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[2]--secretKeyRef--key"><a href="./values.yaml#L70">deploy.envSecret[2].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"password"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[2]--secretKeyRef--name"><a href="./values.yaml#L69">deploy.envSecret[2].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[3]--name"><a href="./values.yaml#L71">deploy.envSecret[3].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_HOST"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[3]--secretKeyRef--key"><a href="./values.yaml#L74">deploy.envSecret[3].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"host"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--envSecret[3]--secretKeyRef--name"><a href="./values.yaml#L73">deploy.envSecret[3].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--image--pullPolicy"><a href="./values.yaml#L82">deploy.image.pullPolicy</a></td>
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
			<td>Global imagePullPolicy Default: 'Always' if image tag is 'latest', else 'IfNotPresent' Ref: http://kubernetes.io/docs/user-guide/images/#pre-pulling-images</td>
		</tr>
		<tr>
			<td id="deploy--image--repository"><a href="./values.yaml#L77">deploy.image.repository</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"registry.gitlab.com/shortlink-org/shortlink/shop_admin"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--image--tag"><a href="./values.yaml#L78">deploy.image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.18.1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="deploy--livenessProbe"><a href="./values.yaml#L93">deploy.livenessProbe</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "enabled": false,
  "httpGet": {
    "path": "/healthz/",
    "port": 8000
  },
  "initialDelaySeconds": 30
}
</pre>
</div>
			</td>
			<td>define a liveness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--readinessProbe"><a href="./values.yaml#L101">deploy.readinessProbe</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "enabled": false,
  "httpGet": {
    "path": "/healthz/",
    "port": 8000
  },
  "initialDelaySeconds": 30
}
</pre>
</div>
			</td>
			<td>define a readiness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--resources--limits"><a href="./values.yaml#L113">deploy.resources.limits</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "cpu": "100m",
  "memory": "128Mi"
}
</pre>
</div>
			</td>
			<td>We usually recommend not to specify default resources and to leave this as a conscious choice for the user. This also increases chances charts run on environments with little resources, such as Minikube. If you do want to specify resources, uncomment the following lines, adjust them as necessary, and remove the curly braces after 'resources:'.</td>
		</tr>
		<tr>
			<td id="deploy--resources--requests--cpu"><a href="./values.yaml#L117">deploy.resources.requests.cpu</a></td>
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
			<td id="deploy--resources--requests--memory"><a href="./values.yaml#L118">deploy.resources.requests.memory</a></td>
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
			<td id="deploy--securityContext"><a href="./values.yaml#L123">deploy.securityContext</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "allowPrivilegeEscalation": false,
  "capabilities": {
    "drop": [
      "ALL"
    ]
  },
  "readOnlyRootFilesystem": "true",
  "runAsGroup": 1000,
  "runAsNonRoot": true,
  "runAsUser": 1000
}
</pre>
</div>
			</td>
			<td>Security Context policies for controller pods See https://kubernetes.io/docs/tasks/administer-cluster/sysctl-cluster/ for notes on enabling and using sysctls</td>
		</tr>
		<tr>
			<td id="deploy--startupProbe"><a href="./values.yaml#L85">deploy.startupProbe</a></td>
			<td>
object
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
{
  "enabled": false,
  "httpGet": {
    "path": "/healthz/",
    "port": 8000
  },
  "initialDelaySeconds": 30
}
</pre>
</div>
			</td>
			<td>define a liveness probe that checks every 5 seconds, starting after 5 seconds</td>
		</tr>
		<tr>
			<td id="deploy--type"><a href="./values.yaml#L44">deploy.type</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"Deployment"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--annotations--"cert-manager--io/cluster-issuer""><a href="./values.yaml#L25">ingress.annotations."cert-manager.io/cluster-issuer"</a></td>
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
			<td id="ingress--annotations--"nginx--ingress--kubernetes--io/enable-opentelemetry""><a href="./values.yaml#L27">ingress.annotations."nginx.ingress.kubernetes.io/enable-opentelemetry"</a></td>
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
			<td id="ingress--annotations--"nginx--ingress--kubernetes--io/enable-owasp-core-rules""><a href="./values.yaml#L26">ingress.annotations."nginx.ingress.kubernetes.io/enable-owasp-core-rules"</a></td>
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
			<td id="ingress--enabled"><a href="./values.yaml#L20">ingress.enabled</a></td>
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
			<td id="ingress--hostname"><a href="./values.yaml#L29">ingress.hostname</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--ingressClassName"><a href="./values.yaml#L22">ingress.ingressClassName</a></td>
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
			<td id="ingress--paths[0]--path"><a href="./values.yaml#L31">ingress.paths[0].path</a></td>
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
			<td id="ingress--paths[0]--service--name"><a href="./values.yaml#L33">ingress.paths[0].service.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shortlink-shop-admin"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--paths[0]--service--port"><a href="./values.yaml#L34">ingress.paths[0].service.port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
8000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--tls[0]--hosts[0]"><a href="./values.yaml#L39">ingress.tls[0].hosts[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop.shortlink.best"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="ingress--tls[0]--secretName"><a href="./values.yaml#L37">ingress.tls[0].secretName</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-admin-tls"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--command[0]"><a href="./values.yaml#L155">initContainers[0].command[0]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"python"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--command[1]"><a href="./values.yaml#L155">initContainers[0].command[1]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"src/migration.py"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--command[2]"><a href="./values.yaml#L155">initContainers[0].command[2]</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"migrate"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[0]--name"><a href="./values.yaml#L158">initContainers[0].envSecret[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_DB"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[0]--secretKeyRef--key"><a href="./values.yaml#L161">initContainers[0].envSecret[0].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"dbname"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[0]--secretKeyRef--name"><a href="./values.yaml#L160">initContainers[0].envSecret[0].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[1]--name"><a href="./values.yaml#L162">initContainers[0].envSecret[1].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_USER"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[1]--secretKeyRef--key"><a href="./values.yaml#L165">initContainers[0].envSecret[1].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"user"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[1]--secretKeyRef--name"><a href="./values.yaml#L164">initContainers[0].envSecret[1].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[2]--name"><a href="./values.yaml#L166">initContainers[0].envSecret[2].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_PASSWORD"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[2]--secretKeyRef--key"><a href="./values.yaml#L169">initContainers[0].envSecret[2].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"password"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[2]--secretKeyRef--name"><a href="./values.yaml#L168">initContainers[0].envSecret[2].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[3]--name"><a href="./values.yaml#L170">initContainers[0].envSecret[3].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"POSTGRES_HOST"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[3]--secretKeyRef--key"><a href="./values.yaml#L173">initContainers[0].envSecret[3].secretKeyRef.key</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"host"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--envSecret[3]--secretKeyRef--name"><a href="./values.yaml#L172">initContainers[0].envSecret[3].secretKeyRef.name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"shop-postgres-pguser-shop"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--image--repository"><a href="./values.yaml#L153">initContainers[0].image.repository</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"registry.gitlab.com/shortlink-org/shortlink/shop_admin"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--image--tag"><a href="./values.yaml#L154">initContainers[0].image.tag</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"0.18.1"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="initContainers[0]--name"><a href="./values.yaml#L151">initContainers[0].name</a></td>
			<td>
string
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
"migration"
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="monitoring--enabled"><a href="./values.yaml#L178">monitoring.enabled</a></td>
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
			<td id="networkPolicy--enabled"><a href="./values.yaml#L183">networkPolicy.enabled</a></td>
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
			<td id="service--ports[0]--name"><a href="./values.yaml#L140">service.ports[0].name</a></td>
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
			<td id="service--ports[0]--port"><a href="./values.yaml#L141">service.ports[0].port</a></td>
			<td>
int
</td>
			<td>
				<div style="max-width: 300px;">
<pre lang="json">
8000
</pre>
</div>
			</td>
			<td></td>
		</tr>
		<tr>
			<td id="service--ports[0]--protocol"><a href="./values.yaml#L142">service.ports[0].protocol</a></td>
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
			<td id="service--ports[0]--public"><a href="./values.yaml#L143">service.ports[0].public</a></td>
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
			<td id="service--ports[1]--name"><a href="./values.yaml#L144">service.ports[1].name</a></td>
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
			<td id="service--ports[1]--port"><a href="./values.yaml#L145">service.ports[1].port</a></td>
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
			<td id="service--ports[1]--protocol"><a href="./values.yaml#L146">service.ports[1].protocol</a></td>
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
			<td id="service--ports[1]--public"><a href="./values.yaml#L147">service.ports[1].public</a></td>
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
			<td id="service--type"><a href="./values.yaml#L138">service.type</a></td>
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
	</tbody>
</table>

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
