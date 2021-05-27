helm-install-jaeger:
	@helm upgrade jaeger-operator ops/Helm/addons/jaeger-operator \
		--install \
		--namespace=jaeger-operator \
		--create-namespace=true \
		--wait
