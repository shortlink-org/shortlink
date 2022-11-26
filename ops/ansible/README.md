### Ansible

#### Install

```
ansible-galaxy install -r requirements.yml
ansible-playbook playbooks/playbook.yml
```

#### Vagrant

```
cd ops/vagrant
vagrant up

cd ops/ansible
ansible-playbook playbooks/playbook.yml
```

#### DNS/HTTP

+ `ui-next.shortlink.vagrant:8081`
