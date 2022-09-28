package router

import "github.com/gin-gonic/gin"

func LRouterGroup(engine *gin.Engine) {
	{
		groupOne := engine.Group("/one")
		groupOne.Use()
		LRouterOneGroup(groupOne)

	}

	{
		groupTwo := engine.Group("/two")
		LRouterTwoGroup(groupTwo)
	}

}
