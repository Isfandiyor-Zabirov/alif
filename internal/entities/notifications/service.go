package notifications

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	PURCHASE = "Purchase"
	VERIFY   = "CardVerify"
	OTP      = "SendOtp"
)

func CreateNotification(event Event) error {

	var exists bool
	for _, existing := range notificationList {
		if existing.ID == event.SessionID {
			exists = true
			break
		}
	}

	if !exists {
		dateFormat, err := time.Parse("2006-01-02 15:04:05 +15:04", event.EventDate)
		if err != nil {
			log.Println("CreateNotification func parse time error:", err.Error())
			return errors.New("Something went wrong ")
		}

		formattedDate := dateFormat.Format("02.01.2006 15:04:05")
		notification := Notification{
			ID:    event.SessionID,
			Title: getNotificationTitle(event.OrderType, event.Card),
			Body:  getNotificationBody(event.OrderType, event.WebsiteUrl, formattedDate),
			Sent:  false,
		}
		return createNotification(notification)
	}

	return nil
}

func GetUnsentNotifications() []Notification {
	return getUnsentNotifications()
}

func ChangeNotificationStatusByID(id string) error {
	return changeNotificationStatusByID(id)
}

func GetAllNotifications() []*Notification {
	return getAllNotifications()
}

func getNotificationTitle(orderType, card string) string {
	switch orderType {
	case PURCHASE:
		return fmt.Sprintf("New purchase with card %s", card)
	case VERIFY:
		return fmt.Sprintf("Verification code for card %s", card)
	case OTP:
		return fmt.Sprintf("Otp code for card %s", card)
	default:
		return ""
	}
}

func getNotificationBody(orderType, url, date string) string {

	switch orderType {
	case PURCHASE:
		return fmt.Sprintf("New purchase made from %s at %s", url, date)
	case VERIFY:
		return fmt.Sprintf("Please verify your cards operation at %s via %s", date, url)
	case OTP:
		minValue := 1000
		maxValue := 9999
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(maxValue-minValue) + minValue
		return fmt.Sprintf("Your otp code is %d. Please enter otp code to confirm your operation at %s via %s. Don't tell code to anybody", number, date, url)
	default:
		return ""
	}
}
