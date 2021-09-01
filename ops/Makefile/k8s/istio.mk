# ISTIO TASKS ==========================================================================================================
ISTIO_VERSION := 1.11.1
ISTIO_NAMESPACE := istio-system
ISTIO_CHART_PATH := ops/Helm/addons/gateway/istio/charts

istio-dep: ## Install istio
	@curl -L https://istio.io/downloadIstio | ISTIO_VERSION=${ISTIO_VERSION} sh -
	@sudo mv istio-${ISTIO_VERSION}/bin/istioctl /usr/bin/istioctl
	@rsync -a -v istio-${ISTIO_VERSION}/manifests/charts ops/Helm/addons/gateway/istio
	@rm -rf istio-${ISTIO_VERSION}

istio-up: ## Run istio
	@helm upgrade istio-base ${ISTIO_CHART_PATH}/base  \
		--install \
		--namespace=${ISTIO_NAMESPACE} \
		--create-namespace=true \
		--wait
	@helm upgrade istiod ${ISTIO_CHART_PATH}/istio-control/istio-discovery \
		--install \
		--namespace=${ISTIO_NAMESPACE} \
		--create-namespace=true \
		--wait
	@helm upgrade istio-ingress ${ISTIO_CHART_PATH}/gateways/istio-ingress \
		--install \
		--namespace=${ISTIO_NAMESPACE} \
		--create-namespace=true \
		--wait
	@helm install istio-egress ${ISTIO_CHART_PATH}/gateways/istio-egress -n ${ISTIO_NAMESPACE}

istio-down: ## Delete istio
	@helm delete istio-egress -n ${ISTIO_NAMESPACE}
	@helm delete istio-ingress -n ${ISTIO_NAMESPACE}
	@helm delete istiod -n ${ISTIO_NAMESPACE}
	@helm delete istio-base -n ${ISTIO_NAMESPACE}
	@kubectl delete namespace ${ISTIO_NAMESPACE}
	# Drop istio CRD
	@kubectl get crd | grep --color=never 'istio.io' | awk '{print $1}' | xargs -n1 kubectl delete crd

# MetalLB TASKS ========================================================================================================
METALLB_SECRET := "$(openssl rand -base64 128)"
METALLB_VERSION := v0.9.6

metallb-up: ## Run MetalLB
	@kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/namespace.yaml
	@kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml
	# On first install only
	@kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="${METALLB_SECRET}"
	# Apply configuration
	@kubectl apply -f ops/Helm/addons/metallb/metallb.yaml

metallb-down: ## Down MetalLB
	@kubectl delete -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/namespace.yaml
	@kubectl delete -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml
	# On first install only
	@kubectl delete secret -n metallb-system memberlist
	# Apply configuration
	@kubectl delete -f ops/Helm/addons/metallb/metallb.yaml
