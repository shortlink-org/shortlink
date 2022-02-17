# SHORTLINK TASKS ======================================================================================================
SHORTLINK_NAMESPACE := shortlink
SHORTLINK_HELM_INGRESS := ops/Helm/shortlink-ingress
HELM_CHART_NGINX_INGRESS := ops/Helm/addons/gateway/nginx-ingress

helm-export-env: ## export env variables for Helm
	echo ISTIO_NAMESPACE=${ISTIO_NAMESPACE}
	echo SHORTLINK_NAMESPACE=${SHORTLINK_NAMESPACE}
	echo SHORTLINK_HELM_INGRESS=${SHORTLINK_HELM_INGRESS}
	echo HELM_CHART_NGINX_INGRESS=${HELM_CHART_NGINX_INGRESS}
	echo SHORTLINK_HELM_API=${SHORTLINK_HELM_API}
	echo SHORTLINK_HELM_LOGGER=${SHORTLINK_HELM_LOGGER}
	echo SHORTLINK_HELM_METADATA=${SHORTLINK_HELM_METADATA}
	echo SHORTLINK_HELM_LINK=${SHORTLINK_HELM_LINK}
	echo SHORTLINK_HELM_BOT=${SHORTLINK_HELM_BOT}
	echo SHORTLINK_HELM_UI=${SHORTLINK_HELM_UI}
	echo SHORTLINK_HELM_LANDING=${SHORTLINK_HELM_LANDING}
	echo SHORTLINK_HELM_PROXY=${SHORTLINK_HELM_PROXY}
	echo SHORTLINK_HELM_BOT=${SHORTLINK_HELM_BOT}

helm-shortlink-dep: ## set need dependencies for shortlink
	-kubectl create namespace ${SHORTLINK_NAMESPACE}
	-kubectl label namespace ${SHORTLINK_NAMESPACE} istio-injection=enabled

helm-shortlink-up: ## run shortlink in k8s by Helm
	-make helm-shortlink-dep

	@helm upgrade nginx-ingress ${HELM_CHART_NGINX_INGRESS} \
		--install \
		--namespace=nginx-ingress \
		--create-namespace=true \
		--set ingress-nginx.controller.metrics.enabled=false \
		--wait

	@helm upgrade shortlink-ingress ${SHORTLINK_HELM_INGRESS} \
		--install \
		--namespace=${ISTIO_NAMESPACE} \
		--create-namespace=true \
		--wait

	@helm upgrade rabbitmq ops/Helm/addons/rabbitmq \
		--install \
		--namespace=rabbitmq \
		--create-namespace=true \
		--set rabbitmq.ingress.tls=false \
		--set rabbitmq.metrics.enabled=false \
		--wait

	@helm upgrade shortlink-api ${SHORTLINK_HELM_API} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=true \
		--set ingress.enabled=true \
		--set deploy.env.MQ_ENABLED=true \
		--set host=shortlink.local \
		--wait

	@helm upgrade shortlink-metadata ${SHORTLINK_HELM_METADATA} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set deploy.image.tag=0.10.26 --set deploy.image.repository=registry.gitlab.com/shortlink-org/shortlink/metadata \
		--wait

	@helm upgrade shortlink-link ${SHORTLINK_HELM_LINK} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--wait

	@helm upgrade shortlink-landing ${SHORTLINK_HELM_LANDING} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set host=shortlink.local \
		--set ingress.enabled=true \
		--set ingress.type=nginx \
		--wait

	@helm upgrade shortlink-next ${SHORTLINK_HELM_UI} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set path=next \
		--set ingress.enabled=true \
		--set host=shortlink.local \
		--wait

	@helm upgrade shortlink-proxy ${SHORTLINK_HELM_PROXY} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set path=s \
		--set ingress.enabled=true \
		--set host=shortlink.local \
		--wait

	@helm upgrade shortlink-bot ${SHORTLINK_HELM_BOT} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set deploy.env.TELEGRAM_BOT_TOKEN=${TELEGRAM_BOT_TOKEN} \
		--set deploy.env.TELEGRAM_BOT_USERNAME=${TELEGRAM_BOT_USERNAME} \
		--wait

helm-shortlink-down: ## Clean artifact from K8S
	for i in $(SHORTLINK_SERVICES); do \
		helm -n ${SHORTLINK_NAMESPACE} del shortlink-$$i; \
  	done
