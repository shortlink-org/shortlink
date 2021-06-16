# AUTH TASKS ===========================================================================================================
AUTH_NAMESPACE := shortlink

auth-up: ## Run auth-stack
	@helm upgrade auth ops/Helm/shortlink-auth  \
		--install \
		--namespace=${AUTH_NAMESPACE} \
		--create-namespace=true \
		--wait
