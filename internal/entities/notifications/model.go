package notifications

type Notification struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Sent  bool   `json:"sent"`
}

type Event struct {
	OrderType  string `json:"orderType"`
	SessionID  string `json:"sessionId"`
	Card       string `json:"card"`
	EventDate  string `json:"eventDate"`
	WebsiteUrl string `json:"websiteUrl"`
}
