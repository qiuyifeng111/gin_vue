package main

import (
	"fmt"
	"goforpra/common"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	common.InitDB()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		r.Run(":" + port)
	}

}
func InitConfig() {
	workdir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdir + "\\config")
	fmt.Println(workdir + "\\config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
}
