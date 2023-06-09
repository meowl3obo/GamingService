package main

import (
	"fmt"
	"os"
	"path/filepath"

	config "gaming-service/config"
	"gaming-service/controller"
	_ "gaming-service/docs"
	lolProvider "gaming-service/provider/lol"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ENV string

func init() {
	loadEnv()
	lolVersions := lolProvider.GetVersions()
	supportLangs := lolProvider.GetLangs()
	os.Setenv("LOL_VERSION", lolVersions[0])
	config.SetLangList(supportLangs)
	config.SetVersionList(lolVersions)
}

// @title Gaming Service Gin swagger
// @version 1.0
// @description Gin swagger

// @host localhost:8080
// schemes http
func main() {
	r := gin.Default()                   //1. 註冊一個路由器
	r.RedirectFixedPath = true           //   自動修正url 允許大小寫
	controller.NewController(r).Router() //2. 建立新的Router
	r.Run("0.0.0.0:8080")                //3. 執行
}

func loadEnv() {
	envPath := ".env"

	if ENV != "" {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		envPath, err = filepath.EvalSymlinks(exPath)
		if err != nil {
			panic(err)
		}
		envPath = fmt.Sprintf("%v/.env.%v", envPath, ENV)
	}

	err := godotenv.Load(envPath)

	if err != nil {
		fmt.Println("env load fail: ", err)
	}
}
