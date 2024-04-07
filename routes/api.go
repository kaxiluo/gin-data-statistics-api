package routes

import (
	"gin-api/app/controllers/demo"
	"gin-api/app/controllers/st"
	"github.com/gin-gonic/gin"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/demo1", demo.Test1)

	router.POST("/st/user_reach", st.UserReach)
	router.POST("/st/user_reach/RunQueueConsumer", st.RunQueueConsumer)
}
