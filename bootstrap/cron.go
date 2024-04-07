package bootstrap

import (
	"fmt"
	"gin-api/app/crontable"
	"gin-api/global"
	"github.com/robfig/cron/v3"
)

func InitializeCron() {
	global.App.Cron = cron.New(cron.WithSeconds())

	go func() {
		global.App.Cron.AddFunc("0 */5 * * * *", crontable.UserReachQueueConsumer)
		global.App.Cron.Start()
		fmt.Println("Cron Start")
		defer global.App.Cron.Stop()
		select {}
	}()
}
