package router

import "github.com/gin-gonic/gin"

func LRouterTwoGroup(groupTwo *gin.RouterGroup) {
	groupTwo.POST("/login", GetCar)
}

func GetCar(c *gin.Context) {

}
