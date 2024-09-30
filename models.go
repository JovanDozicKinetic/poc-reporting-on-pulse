package main

import "time"

type EventData struct {
	EventID      int       `json:"eventId"`
	EventDate    time.Time `json:"eventDate"`
	CateringType string    `json:"cateringType"`
	Confirmed    int       `json:"confirmed"`
	Provisional  int       `json:"provisional"`
	Total        int       `json:"total"`
}
