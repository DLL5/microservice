package main

import (
	"awesomeProject/responseWraper"
	"awesomeProject/router"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// InitProgram 初始化项目
func InitProgram() {
	responseWraper.InitErrorCodeToExplain()
	//InitDataBase()
	//InitRedis()
}

// Begin 开始运行项目
func Begin() {
	Engine := gin.Default()
	gin.New()
	Engine.Use()
	v1 := Engine.Group("/v1")
	{
		// POST 请求方法路由集合
		v1.POST("/ldl/", func(c *gin.Context) {
			fmt.Println("Hello world!")
			//fmt.Println(responseWraper.Non)
			id, _ := strconv.Atoi(c.PostForm("id"))
			type Ldl struct {
				ID     int
				Name   string
				Gender int
			}
			ldl := Ldl{ID: id, Name: "ldl", Gender: 1}
			if ldl.ID > 0 {
				responseWraper.ResponseJson(c, ldl)
			} else {
				responseWraper.ResponseJsonErrCode(c, 1001)
			}
		})

		// GET 请求方法路由集合
		{
			v1.GET("/Hello", Hello)
		}
	}
	router.LRouterGroup(Engine)
	fmt.Println("The awesomeProject running at http://127.0.0.1:8080/v1/Hello ")
	log.Fatal(Engine.Run(":8080"))
}

// Hello 返回hello word字符串
func Hello(c *gin.Context) {
	responseWraper.ResponseJson(c, "Hello world! ")
}
