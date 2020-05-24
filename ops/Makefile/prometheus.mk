# PROMETHEUS TASKS =====================================================================================================
PROMETHEUS_NAMESPACE := prometheus
export PROMETHEUS_SLACK := https://hooks.slack.com/services/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

# docs: https://www.robustperception.io/sending-email-with-the-alertmanager-via-gmail
export PROMETHEUS_SLACK_CHANNEL := \#monitoring
export PROMETHEUS_MAIL_ACCOUNT := user@gmail.com
export PROMETHEUS_MAIL_HOST := smtp.gmail.com:587
export PROMETHEUS_MAIL_PASSWORD := you_auth_token

helm-prometheus-up:
	@helm repo add stable https://kubernetes-charts.storage.googleapis.com
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_alertmanagers.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_podmonitors.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_prometheuses.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_servicemonitors.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_thanosrulers.yaml
	# Custom setting values
	@envsubst < ops/Helm/prometheus-operator.values.yaml > /tmp/prometheus-operator.values.yaml
	@helm upgrade prometheus stable/prometheus-operator \
		--install \
		--namespace=${PROMETHEUS_NAMESPACE} \
        --create-namespace=true \
		--wait \
		-f /tmp/prometheus-operator.values.yaml

helm-prometheus-down:
	@helm --namespace=${PROMETHEUS_NAMESPACE} delete prometheus
	@kubectl delete crd prometheuses.monitoring.coreos.com
	@kubectl delete crd prometheusrules.monitoring.coreos.com
	@kubectl delete crd servicemonitors.monitoring.coreos.com
	@kubectl delete crd podmonitors.monitoring.coreos.com
	@kubectl delete crd alertmanagers.monitoring.coreos.com
	@kubectl delete crd thanosrulers.monitoring.coreos.com
