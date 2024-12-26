package roomusage

import "time"

type TemplateData struct {

	// Main data:
	SiteName string    `json:"SiteName"`
	DateFrom time.Time `json:"DateFrom"`
	DateTo   time.Time `json:"DateTo"`

	// If this is false, then ChangedSince (will be nil) flags will be off
	HasChangedSince bool      `json:"HasChangedSince"`
	ChangedSince    time.Time `json:"ChangedSince"`

	// // This will flag which section of the template will be rendered so we can use only one template
	// IsMeetingRoomsIncluded bool `json:"IsMeetingRoomsIncluded"`

	// // This will be used if IsMeetingRoomsIncluded is false
	// EventsRunning []EventRunning `json:"Events"`

	// // This will be used if IsMeetingRoomsIncluded is true
	// MeetingRooms []MeetingRoom `json:"MeetingRooms"`
}
