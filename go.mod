module github.com/tengfei31/website

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.2
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/go-delve/delve v1.5.0 // indirect
	github.com/go-ini/ini v1.60.0
	github.com/go-openapi/spec v0.19.9 // indirect
	github.com/go-openapi/swag v0.19.9 // indirect
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli/v2 v2.2.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200824131525-c12d262b63d8 // indirect
	golang.org/x/tools v0.0.0-20200823205832-c024452afbcd // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/tengfei31/website/conf => ./conf
	github.com/tengfei31/website/middleware => ./middleware
	github.com/tengfei31/website/middleware/jwt => ./middleware/jwt
	github.com/tengfei31/website/models => ./models
	github.com/tengfei31/website/pkg/e => ./pkg/e
	github.com/tengfei31/website/pkg/setting => ./pkg/setting
	github.com/tengfei31/website/pkg/util => ./pkg/util
	github.com/tengfei31/website/routers => ./routers
// github.com/tengfei31/website/routers/api => ./routers/api
// github.com/tengfei31/website/routers/api/v1 => ./routers/api/v1
)
