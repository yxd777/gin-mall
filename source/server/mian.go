package main

import (
	"fmt"
	serviceAccount "server/service/account/account"
	serviceOrder2 "server/service/account/order"
	"server/service/server"
)
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("hello world")
	//sd
	g := gin.New()
	g.GET("/login", server.Login)
	account := g.Group("/account")
	{
		order := account.Group("/order")
		order.POST("/list", serviceOrder2.List)
		order.POST("/get", serviceOrder2.Get)
	}
	account.POST("/profile", serviceAccount.Profile)
	g.Run(":9292")
}
