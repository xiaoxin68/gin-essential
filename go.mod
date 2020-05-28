module gin-essential

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.12
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/viper v1.7.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
)

replace (
	gin-essential/common => ./common
	gin-essential/controller => ./controller
	gin-essential/dto => ./dto
	gin-essential/middleware => ./middleware
	gin-essential/model => ./model
	gin-essential/repository => ./repository
	gin-essential/response => ./response
	gin-essential/router => ./router
	gin-essential/util => ./util
	gin-essential/vo => ./vo
)
