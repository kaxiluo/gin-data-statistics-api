package st

import (
	"gin-api/app/common/request"
	"gin-api/app/common/response"
	"gin-api/app/crontable"
	"gin-api/app/services"
	"github.com/gin-gonic/gin"
	"time"
)

func UserReach(c *gin.Context) {
	var params request.UserReach
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(params, err))
		return
	}
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	params.Date = currentDate

	UserReachService := services.UserReachService()
	isExist := UserReachService.IsExistInRedisBit(params)
	if isExist {
		response.BusinessFail(c, "已存在")
		return
	}

	//meta := models.UserReachMeta{
	//	Appcode: params.Appcode,
	//	Date: params.Date,
	//	UserId:  params.UserId,
	//	Channel: params.Channel,
	//	EventId: params.EventId,
	//}
	//if err := global.App.DB.Create(&meta).Error; err != nil {
	//	response.ServerError(c, err)
	//	return
	//}

	err := UserReachService.PushMetaToQueue(params)
	if err != nil {
		response.ServerError(c, "推送队列失败")
		return
	}

	response.Success(c, nil)
}

func RunQueueConsumer(c *gin.Context) {
	crontable.UserReachQueueConsumer()
}
