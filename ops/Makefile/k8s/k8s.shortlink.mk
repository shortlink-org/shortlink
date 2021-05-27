# SHORTLINK TASKS ======================================================================================================
SHORTLINK_NAMESPACE := shortlink
SHORTLINK_HELM_INGRESS := ops/Helm/shortlink-ingress
HELM_CHART_NGINX_INGRESS := ops/Helm/addons/nginx-ingress

helm-shortlink-dep:
	-kubectl create namespace ${SHORTLINK_NAMESPACE}
	-kubectl label namespace ${SHORTLINK_NAMESPACE} istio-injection=enabled

helm-shortlink-up: ## run shortlink in k8s by Helm
	@echo helm install/update ${PROJECT_NAME}

	-make helm-shortlink-dep

	@helm upgrade nginx-ingress ${HELM_CHART_NGINX_INGRESS} \
		--install \
		--namespace=nginx-ingress \
		--create-namespace=true \
		--set ingress-nginx.controller.metrics.enabled=false \
		--wait

	@helm upgrade shortlink-ingress ${SHORTLINK_HELM_INGRESS} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--wait

	@helm upgrade api ${SHORTLINK_HELM_API} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=true \
		--set ingress.enabled=true \
		--set deploy.env.MQ_ENABLED=true \
		--set host=shortlink.local \
		--wait

	@helm upgrade metadata ${SHORTLINK_HELM_METADATA} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set deploy.env.GRPC_SERVER_TLS_ENABLED=false \
		--wait

	@helm upgrade link ${SHORTLINK_HELM_LINK} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set deploy.env.GRPC_SERVER_TLS_ENABLED=false \
		--wait

	@helm upgrade landing ${SHORTLINK_HELM_LANDING} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set host=shortlink.local \
		--set ingress.enabled=true \
		--set ingress.type=nginx \
		--wait

	@helm upgrade next ${SHORTLINK_HELM_UI} \
		--install \
		--namespace=${SHORTLINK_NAMESPACE} \
		--create-namespace=true \
		--set serviceAccount.create=false \
		--set path=next \
		--set ingress.enabled=true \
		--set host=shortlink.local \
		--wait

helm-shortlink-down: ## Clean artifact from K8S
	for i in $(SHORTLINK_SERVICES); do \
		helm -n ${SHORTLINK_NAMESPACE} del $$i; \
  	done
