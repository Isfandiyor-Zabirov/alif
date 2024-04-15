package jobs

import (
	"alif/internal/entities/notifications"
	"fmt"
	"log"
)

func sendNotifications() {
	log.Println("Starting send notification job")
	list := notifications.GetUnsentNotifications()

	for _, notification := range list {
		fmt.Println(notification)

		err := notifications.ChangeNotificationStatusByID(notification.ID)
		if err != nil {
			return
		}
	}

	log.Println("Finished send notification job")
}
