module gin-essential/middleware

go 1.14

require (
	gin-essential/common v0.0.0-00010101000000-000000000000 // indirect
	gin-essential/response v0.0.0-00010101000000-000000000000 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
)

replace (
	gin-essential/common => ../common
	gin-essential/model => ../model
	gin-essential/response => ../response
)
