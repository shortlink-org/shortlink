# Loki

Stack:
+ loki
+ grafana
+ promtail
+ fluentd-bit (optional) or docker plugin

### Docker-compose

+ [docker driver](https://grafana.com/docs/loki/latest/clients/docker-driver/)

### Up/Down

```
make logger
make down
```

##### How use is it?

```
services:
    nginx:
        ...
        logging:
          driver: loki
          options:
            loki-url: "http://loki:3100/loki/api/v1/push"
            loki-pipeline-stages: |
              - json:
                  expressions:
                    level: level
              - labels:
                  level:
```

##### LogQL

+ [Docs](https://grafana.com/docs/loki/latest/logql/)
