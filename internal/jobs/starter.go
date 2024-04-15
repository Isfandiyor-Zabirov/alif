package jobs

import (
	"github.com/jasonlvhit/gocron"
)

func StartJobs() {
	scheduler := gocron.NewScheduler()

	scheduler.Every(10).Seconds().Do(sendNotifications)

	<-scheduler.Start()
}
