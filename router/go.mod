module gin-essential/router

go 1.14

replace (
	gin-essential/common => ../common
	gin-essential/controller => ../controller
	gin-essential/model => ../model
	gin-essential/response => ../response
	gin-essential/util => ../util
	gin-essential/middleware => ../middleware
	gin-essential/dto => ../dto
    gin-essential/repository => ../repository
    gin-essential/vo => ../vo
)

require (
    gin-essential/controller v0.0.0-00010101000000-000000000000 // indirect
    gin-essential/middleware v0.0.0-00010101000000-000000000000 // indirect
)
