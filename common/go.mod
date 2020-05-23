module gin-essential/common

go 1.14

require (
	gin-essential/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/spf13/viper v1.7.0 // indirect
)

replace (
    gin-essential/model => ../model
)
