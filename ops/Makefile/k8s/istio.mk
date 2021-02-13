# ISTIO TASKS ==========================================================================================================
istio-up: ## Run istio
	@istioctl manifest apply \
		--set profile=demo
	@kubectl label namespace default istio-injection=enabled

# MetalLB TASKS ========================================================================================================
METALLB_SECRET := "$(openssl rand -base64 128)"
METALLB_VERSION := v0.9.5

metallb-up: ## Run MetalLB
	@kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/namespace.yaml
	@kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/${METALLB_VERSION}/manifests/metallb.yaml
	# On first install only
	@kubectl create secret generic -n metallb-system memberlist --from-literal=secretkey="${METALLB_SECRET}"
	# Apply configuration
	@kubectl apply -f ops/Helm/addons/metallb/metallb.yaml
