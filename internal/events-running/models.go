package eventsrunning

import "time"

type EventRunning struct {
	IsChanged          *bool     `json:"IsChanged,omitempty" db:"IsChanged"`
	EventID            int       `json:"EventID" db:"EventID"`
	EventModuleID      int       `json:"EventModuleID" db:"EventModuleID"`
	EventTitle         string    `json:"EventTitle" db:"EventTitle"`
	EventStatus        string    `json:"EventStatus" db:"EventStatus"`
	EventType          string    `json:"EventType" db:"EventType"`
	CreatedBy          string    `json:"CreatedBy" db:"CreatedBy"`
	EventManager       string    `json:"EventManager" db:"EventManager"`
	EventArrivalDate   time.Time `json:"EventArrivalDate" db:"EventArrivalDate"`
	EventStartDate     time.Time `json:"EventStartDate" db:"EventStartDate"`
	EventEndDate       time.Time `json:"EventEndDate" db:"EventEndDate"`
	EventDepartureDate time.Time `json:"EventDepartureDate" db:"EventDepartureDate"`
	EventDayNumbers    int       `json:"EventDayNumbers" db:"EventDayNumbers"`
	EventNightNumbers  int       `json:"EventNightNumbers" db:"EventNightNumbers"`
	PersonID           *int      `json:"PersonID" db:"PersonID"`
	PersonName         *string   `json:"PersonName" db:"PersonName"`
	CompanyID          *int      `json:"CompanyID" db:"CompanyID"`
	CompanyName        *string   `json:"CompanyName" db:"CompanyName"`
	MeetingRoomID      *int      `json:"MeetingRoomID,omitempty" db:"MeetingRoomID"`
	MeetingRoomName    *string   `json:"MeetingRoomName,omitempty" db:"MeetingRoomName"`
}

type MeetingRoom struct {
	MeetingRoomID   int            `json:"MeetingRoomID" db:"MeetingRoomID"`
	MeetingRoomName string         `json:"MeetingRoomName" db:"MeetingRoomName"`
	Events          []EventRunning `json:"Events"`
}
