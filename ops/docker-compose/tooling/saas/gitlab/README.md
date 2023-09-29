## GitLab

### References

- [Run GitLab in docker-compose](https://docs.gitlab.com/ee/install/docker.html)
  - [Backup](https://docs.gitlab.com/ee/install/docker.html#back-up-gitlab)
- [Configuration](https://docs.gitlab.com/omnibus/settings/configuration.html)
  - [`gitlab.rb` template](https://gitlab.com/gitlab-org/omnibus-gitlab/blob/master/files/gitlab-config-template/gitlab.rb.template)

### Backups by Cron

```bash
0 5 * * * docker exec -t gitlab gitlab-rake gitlab:backup:create
0 6 * * * docker exec -t gitlab /bin/sh -c 'umask 0077; tar cfz /var/opt/gitlab/backups/$(date "+etc-gitlab-\%s.tgz") -C /etc/gitlab'
```

### Add runner

```bash
docker exec -it gitlab-runner gitlab-runner register --url "http://gitlab:10180" --clone-url "http://gitlab:10180"
```

### Reconfigure

```bash
docker exec -it gitlab gitlab-ctl reconfigure
docker exec -it gitlab gitlab-ctl restart
docker exec -it gitlab gitlab-ctl status
```

### Troubleshooting

```bash
# Problem with socket
rm ./data/gitlab/data/gitlab-rails/sockets/gitlab.socket
docker restart gitlab
```
