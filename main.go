package main

import (
	"github.com/gin-gonic/gin"
	"github.com/humyPro/golang/logindemo/controller/user"
	"net/http"
)

var uc user.UserController

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./resources/templates/*")
	router.Static("/static", "./resources/static")
	router.StaticFile("/favicon.ico", "./resources/static/favicon.ico")
	router.Static("/html", "./resources/templates")
	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/user", uc.Register)
	router.POST("/user/login", uc.Login)
	group := router.Group("/")
	group.Use(uc.Authorize)
	{
		group.GET("index", func(context *gin.Context) {
			context.HTML(http.StatusOK, "index.html", nil)
		})
		group.GET("user/logout", uc.Logout)
	}
	router.Run(":8081")
}
