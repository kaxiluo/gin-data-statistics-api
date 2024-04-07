package crontable

import (
	"fmt"
	"gin-api/app/services"
	"time"
)

func UserReachQueueConsumer() {
	fmt.Println(time.Now(), "Cron[UserReachQueueConsumer] running...")
	count := services.UserReachService().FlushQueueToMysql(1000)
	fmt.Println(time.Now(), "Cron[UserReachQueueConsumer] insert count", count)
}
