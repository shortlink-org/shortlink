module github.com/shortlink-org/shortlink/boundaries/platform/csi

go 1.22.2

require (
	github.com/container-storage-interface/spec v1.9.0
	github.com/golang/glog v1.2.1
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.6.0
	github.com/kubernetes-csi/csi-test/v5 v5.2.0
	github.com/spf13/cobra v1.8.0
	github.com/spf13/viper v1.18.2
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/otel/trace v1.26.0
	golang.org/x/net v0.24.0
	golang.org/x/sync v0.7.0
	google.golang.org/grpc v1.63.2
	google.golang.org/protobuf v1.34.0
	k8s.io/kubernetes v1.30.0
	k8s.io/mount-utils v0.30.0
	k8s.io/utils v0.0.0-20240102154912-e7106e64919e
)

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.3 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20240424215950-a892ee059fd6 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/onsi/ginkgo/v2 v2.17.2 // indirect
	github.com/onsi/gomega v1.33.1 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	go.opentelemetry.io/otel v1.26.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240103183307-be819d1f06fc // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240415180920-8c6c420018be // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/apimachinery v0.30.0 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.30.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.30.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.30.0
	k8s.io/apiserver => k8s.io/apiserver v0.30.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.30.0
	k8s.io/client-go => k8s.io/client-go v0.30.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.30.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.30.0
	k8s.io/code-generator => k8s.io/code-generator v0.30.0
	k8s.io/component-base => k8s.io/component-base v0.30.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.30.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.30.0
	k8s.io/cri-api => k8s.io/cri-api v0.30.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.30.0
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.30.0
	k8s.io/kms => k8s.io/kms v0.30.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.30.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.30.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.30.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.30.0
	k8s.io/kubectl => k8s.io/kubectl v0.30.0
	k8s.io/kubelet => k8s.io/kubelet v0.30.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.30.0
	k8s.io/metrics => k8s.io/metrics v0.30.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.30.0
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.30.0
)
