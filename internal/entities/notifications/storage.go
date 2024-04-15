package notifications

var notificationList []*Notification

func createNotification(notification Notification) error {
	notificationList = append(notificationList, &notification)
	return nil
}

func getUnsentNotifications() []Notification {
	var list []Notification

	for _, notification := range notificationList {
		if !notification.Sent {
			list = append(list, *notification)
		}
	}

	return list
}

func changeNotificationStatusByID(id string) error {
	for i, notification := range notificationList {
		if notification.ID == id {
			notificationList[i].Sent = true
		}
	}

	return nil
}

func getAllNotifications() []*Notification {
	return notificationList
}
