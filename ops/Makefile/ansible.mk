# ANSIBLE TASKS ========================================================================================================
export ANSIBLE_CONFIG=ops/ansible/ansible.cfg

ansible-dep: ## Install ansible dep
	@ansible-galaxy install -r ops/ansible/requirements.yml

ansible-up: ## Apply ansible playbook
	@ansible-playbook ops/ansible/playbooks/playbook.yml -i ops/ansible/hosts \
		--vault-password-file ops/ansible/vault-password.txt

ansible-conf: ## Edit secret variable
	@ansible-vault edit ops/ansible/playbooks/group_vars/values.yml \
		--vault-password-file ops/ansible/vault-password.txt

ansible-do: ## Test
	@ansible-playbook ops/ansible/playbooks/playbook.yml -v -i ops/ansible/hosts \
 		--vault-password-file ops/ansible/vault-password.txt \
 		--tags "test"

