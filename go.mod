module github.com/keptn/go-utils

go 1.17

require (
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/benbjohnson/clock v1.3.0
	github.com/cloudevents/sdk-go/observability/opentelemetry/v2 v2.0.0-20211001212819-74757a691209
	github.com/cloudevents/sdk-go/v2 v2.10.0
	github.com/google/uuid v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/nats-io/nats-server/v2 v2.8.4
	github.com/nats-io/nats.go v1.16.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.32.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.2.0
	go.opentelemetry.io/otel/sdk v1.2.0
	go.opentelemetry.io/otel/trace v1.7.0
	golang.org/x/oauth2 v0.0.0-20220608161450-d0670ef3b1eb
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.11
	k8s.io/apimachinery v0.22.11
	k8s.io/client-go v0.22.11
)

require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/evanphx/json-patch v4.12.0+incompatible // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.14.4 // indirect
	github.com/minio/highwayhash v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/jwt/v2 v2.2.1-0.20220330180145-442af02fd36a // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.15.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.2.0 // indirect
	go.opentelemetry.io/otel/metric v0.30.0 // indirect
	go.opentelemetry.io/proto/otlp v0.10.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.19.0 // indirect
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd // indirect
	golang.org/x/lint v0.0.0-20210508222113-6edffad5e616 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220209214540-3681064d5158 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220210224613-90d013bbcef8 // indirect
	golang.org/x/tools v0.1.10-0.20220218145154-897bd77cd717 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368 // indirect
	google.golang.org/grpc v1.43.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	k8s.io/kube-openapi v0.0.0-20211109043538-20434351676c // indirect
	k8s.io/utils v0.0.0-20220210201930-3a6ce19ff2f9 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace golang.org/x/net => golang.org/x/net v0.0.0-20211205041911-012df41ee64c
