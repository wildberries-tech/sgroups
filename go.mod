module github.com/wildberries-tech/sgroups/v2

go 1.22

require (
	connectrpc.com/connect v1.16.2
	github.com/H-BF/corlib v1.2.2-dev
	github.com/ahmetb/go-linq/v3 v3.2.0
	github.com/c-robinson/iplib v1.0.8
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/go-openapi/spec v0.21.0
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/google/nftables v0.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/hashicorp/go-memdb v1.3.4
	github.com/hashicorp/terraform-plugin-framework v1.9.0
	github.com/hashicorp/terraform-plugin-framework-validators v0.12.0
	github.com/hashicorp/terraform-plugin-go v0.23.0
	github.com/hashicorp/terraform-plugin-log v0.9.0
	github.com/hashicorp/terraform-plugin-testing v1.8.0
	github.com/jackc/pgx/v5 v5.6.0
	github.com/jackc/puddle/v2 v2.2.1
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.19.1
	github.com/satori/go.uuid v1.2.0
	github.com/stretchr/testify v1.9.0
	github.com/vishvananda/netns v0.0.4
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/sys v0.22.0
	golang.org/x/text v0.16.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240708141625-4ad9e859172b
	google.golang.org/grpc v1.65.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/google/nftables v0.3.0 => github.com/H-BF/nftables v0.3.0-dev

replace github.com/vishvananda/netlink v1.3.0 => github.com/H-BF/netlink v1.3.0-dev

//TODO: need upgrade to actual version
replace go.opentelemetry.io/otel v1.24.0 => go.opentelemetry.io/otel v1.9.0

//TODO: need upgrade to actual version
replace go.opentelemetry.io/otel/trace v1.24.0 => go.opentelemetry.io/otel/trace v1.9.0

require (
	github.com/ProtonMail/go-crypto v1.1.0-alpha.3-proton // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudflare/circl v1.3.9 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/fatih/color v1.17.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-checkpoint v0.5.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-cty v1.4.1-0.20200414143053-d3edf31b6320 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.6.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/hc-install v0.7.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.21.0 // indirect
	github.com/hashicorp/logutils v1.0.0 // indirect
	github.com/hashicorp/terraform-exec v0.21.0 // indirect
	github.com/hashicorp/terraform-json v0.22.1 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.34.0 // indirect
	github.com/hashicorp/terraform-registry-address v0.2.3 // indirect
	github.com/hashicorp/terraform-svchost v0.1.1 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/josharian/native v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mdlayher/netlink v1.7.2 // indirect
	github.com/mdlayher/socket v0.5.1 // indirect
	github.com/miekg/dns v1.1.61 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/sagikazarmark/locafero v0.6.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.19.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/vishvananda/netlink v1.3.0 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/zclconf/go-cty v1.14.4 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/exp v0.0.0-20240707233637-46b078467d37 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/tools v0.23.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240708141625-4ad9e859172b // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
