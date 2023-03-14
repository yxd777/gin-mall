package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	v1 := g.Group("api/v1")
	fmt.Println(v1)
}
