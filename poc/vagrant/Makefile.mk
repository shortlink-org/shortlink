# KUBEADM TASKS ========================================================================================================

kubeadm-vagrant-up:
	#kubeadm init --apiserver-cert-extra-sans=10.0.7.101 --pod-network-cidr=10.2.0.0/16 --apiserver-advertise-address=0.0.0.0 --kubernetes-version=v1.14.0
	@ansible-playbook ops/ansible/application/playbooks/playbook.yml -v -i ops/ansible/hosts.ini \
 		--vault-password-file ops/ansible/vault-password.txt \
 		--tags "test"
