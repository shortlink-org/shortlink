# ISTIO TASKS ==========================================================================================================
istio-up: ## Run istio
	@istioctl manifest apply \
		--set profile=demo \
		--set gateways.istio-ingressgateway.type=NodePort \
		--set gateways.istio-egressgateway.type=NodePort
	@kubectl label namespace default istio-injection=enabled
