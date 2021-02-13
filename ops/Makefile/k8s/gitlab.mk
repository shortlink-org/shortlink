# GITLAB TASKS =========================================================================================================
GITLAB_NAMESPACE := gitlab

gitlab-minikube: ## Install GitLab to minikube
	@helm repo add gitlab https://charts.gitlab.io/
	@helm repo update
	@helm upgrade -n gitlab --install gitlab gitlab/gitlab \
      --namespace=${GITLAB_NAMESPACE} \
      --create-namespace=true \
	  --set global.hosts.domain=172.17.0.2.nip.io \
 	  --set global.hosts.externalIP=172.17.0.2 \
	  -f ops/docker-compose/tooling/gitlab/helm-value.yaml
	@kubectl -n ${GITLAB_NAMESPACE} get secret gitlab-wildcard-tls-ca -ojsonpath='{.data.cfssl_ca}' | base64 --decode > gitlab.172.17.0.2.nip.io.ca.pem
	@kubectl -n ${GITLAB_NAMESPACE} get secret gitlab-gitlab-initial-root-password -ojsonpath='{.data.password}' | base64 --decode ; echo " - your password"

gitlab-push: ## Push to GitLab
	@GIT_SSL_NO_VERIFY=true git push -u minikube --all
