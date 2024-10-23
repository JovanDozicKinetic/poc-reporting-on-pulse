package eventsrunning

import (
	"encoding/json"
	"log"
)

func GetData() []EventRunning {

	// jsonData := simulateEventsRunningEndpointCallWithoutIsChanged()
	jsonData := simulateEventsRunningEndpointCallWithIsChanged()

	events := unmarshalEventsRunningData(jsonData)

	return events
}

func simulateEventsRunningEndpointCallWithoutIsChanged() []byte {
	return []byte(`
		[
			{
				"eventID": 101,
				"eventModuleID": 1,
				"eventTitle": "Event one",
				"eventStatus": "Confirmed",
				"eventType": "Residential Conference",
				"createdBy": "Creation User",
				"eventManager": "Lisa Rafter",
				"eventArrivalDate": "2024-05-01T09:00:00Z",
				"eventStartDate": "2024-05-01T09:00:00Z",
				"eventEndDate": "2024-05-01T15:30:00Z",
				"eventDepartureDate": "2024-05-01T15:30:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 1,
				"personName": "Lisa Rafter",
				"companyID": 21,
				"companyName": "Company one"
			},
			{
				"eventID": 104,
				"eventModuleID": 4,
				"eventTitle": "Event four",
				"eventStatus": "Confirmed",
				"eventType": "Residential Conference",
				"createdBy": "Creation User",
				"eventManager": "Penelope Pitstop",
				"eventArrivalDate": "2024-07-01T10:00:00Z",
				"eventStartDate": "2024-07-01T10:00:00Z",
				"eventEndDate": "2024-07-01T16:00:00Z",
				"eventDepartureDate": "2024-07-01T16:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 3,
				"personName": "Penelope Pitstop",
				"companyID": 22,
				"companyName": "Company two"
			},
			{
				"eventID": 105,
				"eventModuleID": 5,
				"eventTitle": "Event five",
				"eventStatus": "Provisional",
				"eventType": "Catering Event",
				"createdBy": "Creation User",
				"eventManager": "Robert Magnussen",
				"eventArrivalDate": "2024-08-10T14:00:00Z",
				"eventStartDate": "2024-08-10T14:00:00Z",
				"eventEndDate": "2024-08-10T17:00:00Z",
				"eventDepartureDate": "2024-08-10T17:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 4,
				"personName": "Steven French",
				"companyID": 23,
				"companyName": "Company three"
			},
			{
				"eventID": 106,
				"eventModuleID": 6,
				"eventTitle": "Event six",
				"eventStatus": "Provisional",
				"eventType": "Day Meeting",
				"createdBy": "Creation User",
				"eventManager": "Penelope Pitstop",
				"eventArrivalDate": "2024-04-15T09:00:00Z",
				"eventStartDate": "2024-04-15T09:00:00Z",
				"eventEndDate": "2024-04-15T12:00:00Z",
				"eventDepartureDate": "2024-04-15T12:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 4,
				"personName": "Steven French",
				"companyID": 23,
				"companyName": "Company three"
			},
			{
				"eventID": 1016,
				"eventModuleID": 16,
				"eventTitle": "Event sixteen",
				"eventStatus": "Confirmed",
				"eventType": "Day Meeting",
				"createdBy": "Creation User",
				"eventManager": "Steven French",
				"eventArrivalDate": "2024-04-15T09:00:00Z",
				"eventStartDate": "2024-04-15T09:00:00Z",
				"eventEndDate": "2024-04-15T12:00:00Z",
				"eventDepartureDate": "2024-04-15T12:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": null,
				"personName": null,
				"companyID": 24,
				"companyName": "Company four"
			},
			{
				"eventID": 10106,
				"eventModuleID": 106,
				"eventTitle": "Event one o six",
				"eventStatus": "Provisional",
				"eventType": "Day Meeting",
				"createdBy": "Creation User",
				"eventManager": "Lisa Rafter",
				"eventArrivalDate": "2024-04-15T09:00:00Z",
				"eventStartDate": "2024-04-15T09:00:00Z",
				"eventEndDate": "2024-04-15T12:00:00Z",
				"eventDepartureDate": "2024-04-15T12:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": null,
				"personName": null,
				"companyID": 25,
				"companyName": "Company five"
			}
		]
	`)
}

func simulateEventsRunningEndpointCallWithIsChanged() []byte {
	return []byte(`
		[
			{
				"isChanged": true,
				"eventID": 101,
				"eventModuleID": 1,
				"eventTitle": "Event one",
				"eventStatus": "Confirmed",
				"eventType": "Residential Conference",
				"createdBy": "Creation User",
				"eventManager": "Lisa Rafter",
				"eventArrivalDate": "2024-05-01T09:00:00Z",
				"eventStartDate": "2024-05-01T09:00:00Z",
				"eventEndDate": "2024-05-01T15:30:00Z",
				"eventDepartureDate": "2024-05-01T15:30:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 1,
				"personName": "Lisa Rafter",
				"companyID": 21,
				"companyName": "Company one"
			},
			{
				"isChanged": false,
				"eventID": 785,
				"eventModuleID": 4,
				"eventTitle": "Event four",
				"eventStatus": "Confirmed",
				"eventType": "Residential Conference",
				"createdBy": "Creation User",
				"eventManager": "Penelope Pitstop",
				"eventArrivalDate": "2024-07-01T10:00:00Z",
				"eventStartDate": "2024-07-01T10:00:00Z",
				"eventEndDate": "2024-07-01T16:00:00Z",
				"eventDepartureDate": "2024-07-01T16:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 3,
				"personName": "Penelope Pitstop",
				"companyID": 22,
				"companyName": "Company two"
			},
			{
				"isChanged": true,
				"eventID": 55668,
				"eventModuleID": 5,
				"eventTitle": "Event five but the title is like too long to fit in one row, no - three rows or even four rows I'm typing this and don't know how it's going to look like event",
				"eventStatus": "Provisional",
				"eventType": "Catering Event",
				"createdBy": "Creation User",
				"eventManager": "Robert Magnussen",
				"eventArrivalDate": "2024-08-10T14:00:00Z",
				"eventStartDate": "2024-08-10T14:00:00Z",
				"eventEndDate": "2024-08-10T17:00:00Z",
				"eventDepartureDate": "2024-08-10T17:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 4,
				"personName": "Steven French",
				"companyID": 23,
				"companyName": "Company three"
			},
			{
				"isChanged": false,
				"eventID": 12345,
				"eventModuleID": 6,
				"eventTitle": "Event six",
				"eventStatus": "Provisional",
				"eventType": "Day Meeting",
				"createdBy": "Creation User",
				"eventManager": "Penelope Pitstop",
				"eventArrivalDate": "2024-04-15T09:00:00Z",
				"eventStartDate": "2024-04-15T09:00:00Z",
				"eventEndDate": "2024-04-15T12:00:00Z",
				"eventDepartureDate": "2024-04-15T12:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": 4,
				"personName": "Steven French",
				"companyID": 23,
				"companyName": "Company three"
			},
			{
				"isChanged": true,
				"eventID": 44554,
				"eventModuleID": 16,
				"eventTitle": "Event sixteen",
				"eventStatus": "Confirmed",
				"eventType": "Day Meeting",
				"createdBy": "Creation User",
				"eventManager": "Steven French",
				"eventArrivalDate": "2024-04-15T09:00:00Z",
				"eventStartDate": "2024-04-15T09:00:00Z",
				"eventEndDate": "2024-04-15T12:00:00Z",
				"eventDepartureDate": "2024-04-15T12:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": null,
				"personName": null,
				"companyID": 24,
				"companyName": "Company four"
			},
			{
				"isChanged": false,
				"eventID": 10106,
				"eventModuleID": 106,
				"eventTitle": "Event one o six",
				"eventStatus": "Provisional",
				"eventType": "Day Meeting",
				"createdBy": "Creation User",
				"eventManager": "Lisa Rafter",
				"eventArrivalDate": "2024-04-15T09:00:00Z",
				"eventStartDate": "2024-04-15T09:00:00Z",
				"eventEndDate": "2024-04-15T12:00:00Z",
				"eventDepartureDate": "2024-04-15T12:00:00Z",
				"eventDayNumbers": 0,
				"eventNightNumbers": 0,
				"personID": null,
				"personName": null,
				"companyID": 25,
				"companyName": "Company five"
			}
		]
	`)
}

func unmarshalEventsRunningData(jsonData []byte) []EventRunning {
	var events []EventRunning
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	return events
}
