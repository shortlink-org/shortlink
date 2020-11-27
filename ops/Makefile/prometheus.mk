# PROMETHEUS TASKS =====================================================================================================
PROMETHEUS_NAMESPACE := prometheus
export PROMETHEUS_SLACK := https://hooks.slack.com/services/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

# docs: https://www.robustperception.io/sending-email-with-the-alertmanager-via-gmail
export PROMETHEUS_SLACK_CHANNEL := \#monitoring
export PROMETHEUS_MAIL_ACCOUNT  := user@gmail.com
export PROMETHEUS_MAIL_HOST     := smtp.gmail.com:587
export PROMETHEUS_MAIL_PASSWORD := you_auth_token

# docs: https://github.com/metalmatze/alertmanager-bot
export PROMETHEUS_TELEGRAM_API   := 'WFhYWFhYWA==' # in base64
export PROMETHEUS_TELEGRAM_ADMIN := 'MTIzNA==' # in base64

helm-prometheus-up:
	@helm repo add stable https://charts.helm.sh/stable
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_alertmanagers.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_podmonitors.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_prometheuses.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_prometheusrules.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_servicemonitors.yaml
	@kubectl apply -f https://raw.githubusercontent.com/coreos/prometheus-operator/release-0.38/example/prometheus-operator-crd/monitoring.coreos.com_thanosrulers.yaml
	# Custom setting values
	@envsubst < ops/Helm/addons/monitoring/prometheus-operator.values.yaml > /tmp/prometheus-operator.values.yaml
	@helm upgrade prometheus stable/prometheus-operator \
		--install \
		--namespace=${PROMETHEUS_NAMESPACE} \
        --create-namespace=true \
		--wait \
		-f /tmp/prometheus-operator.values.yaml

prometheus-telegram-alert:
	@envsubst < ops/Helm/addons/monitoring/telegram-alert-bot.yaml > /tmp/telegram-alert-bot.yaml
	@kubectl apply -n ${PROMETHEUS_NAMESPACE} \
		-f /tmp/telegram-alert-bot.yaml

prometheus-telegram-alert-sown:
	@kubectl delete -n ${PROMETHEUS_NAMESPACE} \
		-f /tmp/telegram-alert-bot.yaml

helm-prometheus-down:
	@helm --namespace=${PROMETHEUS_NAMESPACE} delete prometheus
	@kubectl delete crd prometheuses.monitoring.coreos.com
	@kubectl delete crd prometheusrules.monitoring.coreos.com
	@kubectl delete crd servicemonitors.monitoring.coreos.com
	@kubectl delete crd podmonitors.monitoring.coreos.com
	@kubectl delete crd alertmanagers.monitoring.coreos.com
	@kubectl delete crd thanosrulers.monitoring.coreos.com
