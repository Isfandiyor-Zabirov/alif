package notifications

import "testing"

func TestCreateNotification(t *testing.T) {
	lengthBefore := len(notificationList)

	event := Event{
		OrderType:  "Purchase",
		SessionID:  "242123-arr23-ooe34",
		Card:       "4433***0914",
		EventDate:  "2024-03-12 14:03:43.34532 +00:00",
		WebsiteUrl: "https://test.com",
	}

	err := CreateNotification(event)
	if err != nil {
		t.Errorf("Wanted nil error, got %s", err.Error())
	}

	lengthAfterWanted := lengthBefore + 1
	lengthAfterActual := len(notificationList)

	if lengthAfterWanted != lengthAfterActual {
		t.Errorf("Wanted length of %d, got length of %d", lengthAfterWanted, lengthAfterActual)
	}

	event2 := Event{
		OrderType:  "Purchase",
		SessionID:  "242123-arr23-ooe34",
		Card:       "4433***0914",
		EventDate:  "2024-03-12 14:03:43.34532 +00:00",
		WebsiteUrl: "https://test.com",
	}

	err = CreateNotification(event2)
	if err != nil {
		t.Errorf("Testing when the same event received again. Wanted nil error, got %s", err.Error())
	}

	lengthAfterWanted = lengthBefore + 1
	lengthAfterActual = len(notificationList)

	if lengthAfterWanted != lengthAfterActual {
		t.Errorf("Testing when the same event received again. Wanted length of %d, got length of %d", lengthAfterWanted, lengthAfterActual)
	}

	event3 := Event{
		OrderType:  "TopUp",
		SessionID:  "242123-arr23-ooe34",
		Card:       "4433***0914",
		EventDate:  "2024-03-12 14:03:43.34532",
		WebsiteUrl: "https://test.com",
	}

	err = CreateNotification(event3)
	if err == nil {
		t.Errorf("Testing wrong date format. Wanted error message, got %v", err)
	} else {
		wantedErrorMessage := "wrong order type"
		if err.Error() != wantedErrorMessage {
			t.Errorf("Testing wrong date format. Wanted error %s, got %s", wantedErrorMessage, err.Error())
		}
	}
}
