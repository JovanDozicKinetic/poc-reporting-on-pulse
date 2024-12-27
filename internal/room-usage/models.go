package roomusage

import "time"

type TemplateData struct {

	// Main data:
	AreaName string    `json:"areaName"`
	DateFrom time.Time `json:"dateFrom"`
	DateTo   time.Time `json:"dateTo"`

	// If this is false, then ChangedSince (will be nil) flags will be off
	HasChangedSince bool      `json:"hasChangedSince"`
	ChangedSince    time.Time `json:"changedSince"`

	IsSeparated bool `json:"isSeparated"`

	Dates          *[]Date          `json:"dates"`          // This is nil if IsSeparated is true
	DatesSeparated *[]DateSeparated `json:"datesSeparated"` // This is nil if IsSeparated is false
}

type Date struct {
	Date     time.Time `json:"date"`
	Bookings []Booking `json:"bookings"`
}

type DateSeparated struct {
	Date time.Time `json:"date"`

	NewEventsSameDayTakeDown *[]Booking `json:"newEventsSameDayTakeDown"`
	NewEvents                *[]Booking `json:"newEvents"`
	RunningEvents            *[]Booking `json:"runningEvents"`
	TakeDownEvents           *[]Booking `json:"takeDownEvents"`
}

type Booking struct {
	Reference string `json:"reference"`
	Name      string `json:"name"`

	DayCurrent int `json:"dayCurrent"` // Day X of Y
	DaysTotal  int `json:"daysTotal"`  // Y

	RoomName string `json:"roomName"`

	// If both are nil, then we have: `All day`
	// If StartTime is nil, then we have: `Booked Until: xx:xx`
	// If EndTime is nil, then we have: `Booked From: xx:xx`
	// If neither is nil, then we have: `Booked: xx:xx - xx:xx`
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`

	Number int `json:"number"`

	Layout string `json:"layout"`

	RoomNotes string `json:"roomNotes"`

	Equipment []Equipment `json:"equipment"`

	NextUse *time.Time `json:"nextUse"`

	IsChangedSince *bool `json:"changedSince"`
}

type Equipment struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Notes    string `json:"notes"`
}
