package router

import "github.com/gin-gonic/gin"

func LRouterOneGroup(groupOne *gin.RouterGroup) {
	groupOne.POST("/login", Login)
	groupOne.POST("/Register", Register)
	groupOne.POST("/Hello", Hello)
}

func Login(c *gin.Context) {

}

func Register(c *gin.Context) {

}

func Hello(c *gin.Context) {
}
