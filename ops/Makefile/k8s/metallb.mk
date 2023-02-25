# MetalLB TASKS ========================================================================================================
METALLB_SECRET := "$(openssl rand -base64 128)"
METALLB_VERSION := v0.13.9

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
