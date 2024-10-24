package template1

import "time"

type EventData struct {
	EventID      int       `json:"eventId"`
	EventDate    time.Time `json:"eventDate"`
	CateringType string    `json:"cateringType"`
	Confirmed    int       `json:"confirmed"`
	Provisional  int       `json:"provisional"`
	Total        int       `json:"total"`
}

type CateringType struct {
	ID   int    `json:"id"`
	Name string `json:"catering_type"`
}
