module github.com/batazor/shortlink

go 1.16

require (
	github.com/DataDog/zstd v1.4.4 // indirect
	github.com/HdrHistogram/hdrhistogram-go v0.9.0 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/PuerkitoBio/goquery v1.6.1
	github.com/Shopify/sarama v1.29.0
	github.com/Unleash/unleash-client-go/v3 v3.1.1
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/container-storage-interface/spec v1.4.0
	github.com/containerd/continuity v0.0.0-20200107194136-26c1120b8d41 // indirect
	github.com/dgraph-io/badger/v3 v3.2011.1
	github.com/dgraph-io/dgo/v2 v2.2.0
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/getsentry/sentry-go v0.10.0
	github.com/go-chi/chi/v5 v5.0.3
	github.com/go-chi/cors v1.2.0
	github.com/go-chi/render v1.0.1
	github.com/go-kit/kit v0.10.0
	github.com/go-redis/redis/v8 v8.8.2
	github.com/go-sql-driver/mysql v1.6.0
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.0.0-rc1
	github.com/gocql/gocql v0.0.0-20210504150947-558dfae50b5d
	github.com/golang-migrate/migrate/v4 v4.14.2-0.20201125065321-a53e6fc42574
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0 // indirect
	github.com/google/wire v0.5.0
	github.com/gorilla/mux v1.8.0
	github.com/graph-gophers/graphql-go v1.1.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.4.0
	github.com/heptiolabs/healthcheck v0.0.0-20180807145615-6ff867650f40
	github.com/jackc/pgx/v4 v4.11.0
	github.com/jmoiron/sqlx v1.3.3
	github.com/johejo/golang-migrate-extra v0.0.0-20210217013041-51a992e50d16
	github.com/kubernetes-csi/csi-test/v4 v4.2.0
	github.com/lib/pq v1.10.1
	github.com/luna-duclos/instrumentedsql v1.1.3
	github.com/luna-duclos/instrumentedsql/opentracing v0.0.0-20201103091713-40d03108b6f4
	github.com/markbates/pkger v0.17.1
	github.com/mattn/go-sqlite3 v1.14.7
	github.com/nats-io/nats.go v1.11.0
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing-contrib/go-stdlib v1.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/ory/dockertest/v3 v3.6.5
	github.com/ory/kratos-client-go v0.5.5-alpha.4
	github.com/pborman/uuid v1.2.1
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/common v0.20.0 // indirect
	github.com/pterm/pterm v0.12.13
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/scylladb/gocqlx/v2 v2.4.0
	github.com/sirupsen/logrus v1.8.1
	github.com/smartystreets/assertions v1.0.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/uber/jaeger-client-go v2.28.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.mongodb.org/mongo-driver v1.5.2
	go.uber.org/atomic v1.7.0
	go.uber.org/automaxprocs v1.4.0
	go.uber.org/goleak v1.1.10
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/net v0.0.0-20210505214959-0714010a04ed
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/api v0.46.0
	google.golang.org/genproto v0.0.0-20210506142907-4a47615972c2
	google.golang.org/grpc v1.37.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/rethinkdb/rethinkdb-go.v6 v6.2.1
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
	k8s.io/kubernetes v1.21.0
	k8s.io/mount-utils v0.21.0
	k8s.io/utils v0.0.0-20210305010621-2afb4311ab10
)

replace (
	k8s.io/api => k8s.io/api v0.20.0-beta.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.0-beta.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.0-beta.0
	k8s.io/apiserver => k8s.io/apiserver v0.20.0-beta.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.0-beta.0
	k8s.io/client-go => k8s.io/client-go v0.20.0-beta.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.0-beta.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.0-beta.0
	k8s.io/code-generator => k8s.io/code-generator v0.20.0-beta.0
	k8s.io/component-base => k8s.io/component-base v0.20.0-beta.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.0-beta.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.0-beta.0
	k8s.io/cri-api => k8s.io/cri-api v0.20.0-beta.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.0-beta.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.0-beta.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.0-beta.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.0-beta.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.0-beta.0
	k8s.io/kubectl => k8s.io/kubectl v0.20.0-beta.0
	k8s.io/kubelet => k8s.io/kubelet v0.20.0-beta.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.0-beta.0
	k8s.io/metrics => k8s.io/metrics v0.20.0-beta.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.0-beta.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.0-beta.0
)
