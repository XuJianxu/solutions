module github.com/pingcap/test-infra/tools/sql-crc32

go 1.17

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/pingcap/log v1.1.1-0.20221015072633-39906604fb81
	github.com/pingcap/test-infra/caselib v0.0.0-20210916073340-bd791f09a546
	github.com/spf13/cobra v1.5.0
	go.uber.org/zap v1.23.0
)

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Thearas/ozzo-validation/v4 v4.3.2 // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/benbjohnson/clock v1.3.0 // indirect
	github.com/creasty/defaults v1.6.0 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pingcap/errors v0.11.5-0.20220729040631-518f63d66278 // indirect
	github.com/pingcap/kvproto v0.0.0-20221026112947-f8d61344b172 // indirect
	github.com/pingcap/test-infra v0.0.0-20221111082454-122099c8016a // indirect
	github.com/pingcap/test-infra/sdk v0.0.0-20211105074920-61e24194a590 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.13.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect
	golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220719170305-83ca9fad585f // indirect
	google.golang.org/grpc v1.48.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/apimachinery v0.23.5 // indirect
	k8s.io/klog/v2 v2.30.0 // indirect
	k8s.io/utils v0.0.0-20211116205334-6203023598ed // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
)

replace github.com/pingcap/test-infra/caselib => ../../caselib

replace github.com/pingcap/test-infra/sdk => ../../sdk

replace github.com/pingcap/test-infra => ../../