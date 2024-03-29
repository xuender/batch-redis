module github.com/xuender/batch-redis

go 1.12

replace golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190621222207-cc06ce4a13d4

replace golang.org/x/sys => github.com/golang/sys v0.0.0-20190624142023-c5567b49c5d0

replace golang.org/x/net => github.com/golang/net v0.0.0-20190620200207-3b0461eec859

replace cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0

replace golang.org/x/text => github.com/golang/text v0.3.2

replace golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45

replace golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58

replace google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190620144150-6af8c5fc6601

replace golang.org/x/tools => github.com/golang/tools v0.0.0-20190625160430-252024b82959

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1

replace golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522

replace golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4

replace golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422

replace google.golang.org/api => github.com/googleapis/google-api-go-client v0.7.0

replace golang.org/x/image => github.com/golang/image v0.0.0-20190622003408-7e034cad6442

replace golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88

replace google.golang.org/appengine => github.com/golang/appengine v1.6.1

require (
	github.com/go-redis/redis/v7 v7.0.0-beta.4
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/mitchellh/go-homedir v1.1.0
	github.com/onsi/ginkgo v1.10.3 // indirect
	github.com/onsi/gomega v1.7.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
	golang.org/x/crypto v0.0.0-20191117063200-497ca9f6d64f // indirect
	golang.org/x/net v0.0.0-20191118183410-d06c31c94cae // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20191118133127-cf1e2d577169 // indirect
	golang.org/x/tools v0.0.0-20191118222007-07fc4c7f2b98 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
)
