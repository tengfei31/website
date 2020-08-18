module github.com/tengfei31/website

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20200817155316-9781c653f443 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/tengfei31/website/pkg/setting 	=> ./pkg/setting
	github.com/tengfei31/website/conf    	  	=> ./conf
	github.com/tengfei31/website/middleware  	=> ./middleware
	github.com/tengfei31/website/models 	  	=> ./models
	github.com/tengfei31/website/routers 	  	=> ./routers
)

