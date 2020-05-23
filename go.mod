module gin-essential

go 1.14

require (
	gin-essential/common v0.0.0-00010101000000-000000000000
	gin-essential/dto v0.0.0-00010101000000-000000000000 // indirect
	gin-essential/router v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
	github.com/spf13/viper v1.7.0
)

replace (
	gin-essential/common => ./common
	gin-essential/controller => ./controller
	gin-essential/dto => ./dto
	gin-essential/model => ./model
	gin-essential/response => ./response
	gin-essential/router => ./router
	gin-essential/util => ./util
	gin-essential/middleware => ./middleware
	gin-essential/vo => ./vo
	gin-essential/repository => ./repository
)
