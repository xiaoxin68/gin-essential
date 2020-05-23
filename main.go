package main

import (
	"gin-essential/common"
	"gin-essential/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = router.CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run("localhost:" + port))
	}
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
