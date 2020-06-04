module hulujia

go 1.12

require (
	cloud.google.com/go/pubsub v1.4.0 // indirect
	dmitri.shuralyov.com/gpu/mtl v0.0.0-20191203043605-d42048ed14fd // indirect
	github.com/BurntSushi/xgb v0.0.0-20200324125942-20f126ea2843 // indirect
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/aliyun/aliyun-oss-go-sdk v2.1.0+incompatible
	github.com/appleboy/gin-jwt v2.5.0+incompatible
	github.com/appleboy/gofight/v2 v2.1.2 // indirect
	github.com/astaxie/beego v1.12.1
	github.com/cncf/udpa/go v0.0.0-20200508205342-3b31d022a144 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/envoyproxy/go-control-plane v0.9.5 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.3.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20200420212212-258d9bec320e // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20200507031123-427632fa3b1c // indirect
	github.com/ianlancetaylor/demangle v0.0.0-20200524003926-2c5affb30a03 // indirect
	github.com/issue9/identicon v1.0.1
	github.com/jinzhu/gorm v1.9.12
	github.com/kr/pretty v0.2.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/mapstructure v1.3.1 // indirect
	github.com/onsi/ginkgo v1.12.2 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/peterh/liner v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.6.0 // indirect
	github.com/rs/zerolog v1.18.0
	github.com/siddontang/goredis v0.0.0-20180423163523-0b4019cbd7b7 // indirect
	github.com/siddontang/ledisdb v0.0.0-20190202134119-8ceb77e66a92 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.7.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/stretchr/testify v1.6.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/unknwon/com v1.0.1
	github.com/yuin/goldmark v1.1.31 // indirect
	github.com/yuin/gopher-lua v0.0.0-20200521060427-6ff375d91eab // indirect
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/exp v0.0.0-20200513190911-00229845015e // indirect
	golang.org/x/image v0.0.0-20200430140353-33d19683fad8 // indirect
	golang.org/x/mobile v0.0.0-20200329125638-4c31acba0007 // indirect
	golang.org/x/net v0.0.0-20200528225125-3c3fba18258b // indirect
	golang.org/x/tools v0.0.0-20200601175630-2caf76543d99 // indirect
	google.golang.org/api v0.26.0 // indirect
	google.golang.org/genproto v0.0.0-20200601130524-0f60399e6634 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/dgrijalva/jwt-go.v3 v3.2.0
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200601152816-913338de1bd2 // indirect
	rsc.io/sampler v1.99.99 // indirect
)

replace (
	github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
	github.com/jinzhu/gorm v1.9.12 => /Users/flaravel/go/src/github.com/jinzhu/gorm
	go.etcd.io/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.4
	golang.org/x/mod => /Users/flaravel/go/src/golang.org/x/mod
)