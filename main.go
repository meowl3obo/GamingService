package main

import (
	"fmt"

	"gaming-service/controller"

	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("not found env file")
	}

	r := gin.Default()               //1. 註冊一個路由器
	r.RedirectFixedPath = true       //   自動修正url 允許大小寫
	controller.NewController(r).Router() //2. 建立新的Router
	r.Run("0.0.0.0:1234")            //3. 執行
}