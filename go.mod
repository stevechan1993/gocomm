module github.com/stevechan1993/gocomm

go 1.13

require (
	github.com/Shopify/sarama v1.24.0
	github.com/astaxie/beego v1.10.0
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect

	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gomodule/redigo v1.7.0
	github.com/google/go-cmp v0.3.0
	github.com/gorilla/websocket v1.4.1
	github.com/lib/pq v1.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.11.0 // indirect

	github.com/opentracing-contrib/go-stdlib v0.0.0-20190519235532-cf7a6c988dc9
	github.com/opentracing/opentracing-go v1.1.0
	github.com/sony/sonyflake v1.0.0
	github.com/spf13/viper v1.4.0
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber/jaeger-client-go v2.16.0+incompatible
	github.com/uber/jaeger-lib v2.0.0+incompatible
	github.com/ugorji/go v1.1.7 // indirect
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
