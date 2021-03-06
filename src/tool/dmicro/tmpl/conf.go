package tmpl

var (
	ConfAPP=`run_mode: dev`
	ConfAPP_DEV =`
mysql_url: "test:123456@tcp(127.0.0.1:3306)/dmicro"
redis_url: "127.0.0.1:6379"
redis_auth: "123456"
redis_max_idle: 10

name: "go.micro.api.passport"
version: "latest"
trace_addr: "127.0.0.1:6831"

register_ttl: 30
register_interval: 15

rps_limit: 1024
max_concurrent: 1024	
`


	ConfAPP_PROD =`
mysql_url: "test:123456@tcp(127.0.0.1:3306)/dmicro"
redis_url: "127.0.0.1:6379"
redis_auth: "123456"
redis_max_idle: 10

name: "go.micro.api.passport"
version: "latest"
trace_addr: "127.0.0.1:6831"

register_ttl: 30
register_interval: 15

rps_limit: 1024
max_concurrent: 1024
`
)
