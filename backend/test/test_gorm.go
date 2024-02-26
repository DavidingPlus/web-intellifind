package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	g := gin.New()
	g.POST("/login", func(ctx *gin.Context) {
		r := &LoginRequest{}
		ctx.ShouldBind(r)
		fmt.Printf("login-request:%+v\n", r.Username)
	})

	g.Run(":8080")
}
