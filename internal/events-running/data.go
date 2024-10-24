package eventsrunning

import (
	"encoding/json"
	"log"
)

func GetEventDataWithoutMeetingRooms(includeIsChanged bool) []EventRunning {
	var jsonData []byte
	if includeIsChanged {
		jsonData = simulateEventsRunningEndpointCallWithIsChanged()
	} else {
		jsonData = simulateEventsRunningEndpointCallWithoutIsChanged()
	}
	events := unmarshalEventsRunningData(jsonData)
	return events
}

func GetEventDataWithMeetingRooms(includeIsChanged bool) []MeetingRoom {
	var jsonData []byte
	if includeIsChanged {
		jsonData = simulateEventsRunningWithMeetingRoomsWithIsChanged()
	} else {
		jsonData = simulateEventsRunningWithMeetingRoomsWithoutIsChanged()
	}
	meetingRooms := unmarshalMeetingRoomsData(jsonData)
	return meetingRooms

}

func unmarshalEventsRunningData(jsonData []byte) []EventRunning {
	var events []EventRunning
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	return events
}

func unmarshalMeetingRoomsData(jsonData []byte) []MeetingRoom {
	var meetingRooms []MeetingRoom
	err := json.Unmarshal(jsonData, &meetingRooms)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	return meetingRooms
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

func simulateEventsRunningWithMeetingRoomsWithoutIsChanged() []byte {
	return []byte(`
		[
			{
				"meetingRoomID": -1,
				"meetingRoomName": "No meeting room",
				"events": [
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
			},
			{
				"meetingRoomID": 2,
				"meetingRoomName": "Room two",
				"events": [
					{
						"eventID": 102,
						"eventModuleID": 2,
						"eventTitle": "Event two",
						"eventStatus": "Provisional",
						"eventType": "Day Meeting",
						"createdBy": "User For Creation",
						"eventManager": "Lisa Rafter",
						"eventArrivalDate": "2024-06-10T11:00:00Z",
						"eventStartDate": "2024-06-10T11:00:00Z",
						"eventEndDate": "2024-06-10T18:30:00Z",
						"eventDepartureDate": "2024-06-10T18:30:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 2,
						"personName": "Robert Magnussen",
						"companyID": 22,
						"companyName": "Company two",
						"meetingRoomID": 2,
						"meetingRoomName": "Room two"
					},
					{
						"eventID": 107,
						"eventModuleID": 7,
						"eventTitle": "Event seven",
						"eventStatus": "Confirmed",
						"eventType": "Residential Conference",
						"createdBy": "User For Creation",
						"eventManager": "Steven French",
						"eventArrivalDate": "2024-09-01T11:00:00Z",
						"eventStartDate": "2024-09-01T11:00:00Z",
						"eventEndDate": "2024-09-01T17:00:00Z",
						"eventDepartureDate": "2024-09-01T17:00:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 5,
						"personName": "Marko Markovich",
						"companyID": 24,
						"companyName": "Company four",
						"meetingRoomID": 2,
						"meetingRoomName": "Room two"
					}
				]
			},
			{
				"meetingRoomID": 3,
				"meetingRoomName": "Room three",
				"events": [
					{
						"eventID": 103,
						"eventModuleID": 3,
						"eventTitle": "Event three",
						"eventStatus": "Confirmed",
						"eventType": "Residential Conference",
						"createdBy": "User For Creation",
						"eventManager": "Robert Magnussen",
						"eventArrivalDate": "2024-03-15T08:00:00Z",
						"eventStartDate": "2024-03-15T08:00:00Z",
						"eventEndDate": "2024-03-15T12:30:00Z",
						"eventDepartureDate": "2024-03-15T12:30:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 1,
						"personName": "Lisa Rafter",
						"companyID": 21,
						"companyName": "Company one",
						"meetingRoomID": 3,
						"meetingRoomName": "Room three"
					}
				]
			},
			{
				"meetingRoomID": 6,
				"meetingRoomName": "Room six",
				"events": [
					{
						"eventID": 102,
						"eventModuleID": 2,
						"eventTitle": "Event two",
						"eventStatus": "Provisional",
						"eventType": "Day Meeting",
						"createdBy": "User For Creation",
						"eventManager": "Lisa Rafter",
						"eventArrivalDate": "2024-06-10T11:00:00Z",
						"eventStartDate": "2024-06-10T11:00:00Z",
						"eventEndDate": "2024-06-10T18:30:00Z",
						"eventDepartureDate": "2024-06-10T18:30:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 2,
						"personName": "Robert Magnussen",
						"companyID": 22,
						"companyName": "Company two",
						"meetingRoomID": 6,
						"meetingRoomName": "Room six"
					}
				]
			}
		]
	`)
}

func simulateEventsRunningWithMeetingRoomsWithIsChanged() []byte {
	return []byte(`
		[
			{
				"meetingRoomID": 2,
				"meetingRoomName": "Room two",
				"events": [
					{
						"isChanged": true,
						"eventID": 102,
						"eventModuleID": 2,
						"eventTitle": "Event two",
						"eventStatus": "Provisional",
						"eventType": "Day Meeting",
						"createdBy": "User For Creation",
						"eventManager": "Lisa Rafter",
						"eventArrivalDate": "2024-06-10T11:00:00Z",
						"eventStartDate": "2024-06-10T11:00:00Z",
						"eventEndDate": "2024-06-10T18:30:00Z",
						"eventDepartureDate": "2024-06-10T18:30:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 2,
						"personName": "Robert Magnussen",
						"companyID": 22,
						"companyName": "Company two",
						"meetingRoomID": 2,
						"meetingRoomName": "Room two"
					},
					{
						"isChanged": false,
						"eventID": 107,
						"eventModuleID": 7,
						"eventTitle": "Event seven",
						"eventStatus": "Confirmed",
						"eventType": "Residential Conference",
						"createdBy": "User For Creation",
						"eventManager": "Steven French",
						"eventArrivalDate": "2024-09-01T11:00:00Z",
						"eventStartDate": "2024-09-01T11:00:00Z",
						"eventEndDate": "2024-09-01T17:00:00Z",
						"eventDepartureDate": "2024-09-01T17:00:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 5,
						"personName": "Marko Markovich",
						"companyID": 24,
						"companyName": "Company four",
						"meetingRoomID": 2,
						"meetingRoomName": "Room two"
					}
				]
			},
			{
				"meetingRoomID": 3,
				"meetingRoomName": "Room three",
				"events": [
					{
						"isChanged": false,
						"eventID": 103,
						"eventModuleID": 3,
						"eventTitle": "Event three",
						"eventStatus": "Confirmed",
						"eventType": "Residential Conference",
						"createdBy": "User For Creation",
						"eventManager": "Robert Magnussen",
						"eventArrivalDate": "2024-03-15T08:00:00Z",
						"eventStartDate": "2024-03-15T08:00:00Z",
						"eventEndDate": "2024-03-15T12:30:00Z",
						"eventDepartureDate": "2024-03-15T12:30:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 1,
						"personName": "Lisa Rafter",
						"companyID": 21,
						"companyName": "Company one",
						"meetingRoomID": 3,
						"meetingRoomName": "Room three"
					}
				]
			},
			{
				"meetingRoomID": 6,
				"meetingRoomName": "Room six",
				"events": [
					{
						"isChanged": true,
						"eventID": 102,
						"eventModuleID": 2,
						"eventTitle": "Event two",
						"eventStatus": "Provisional",
						"eventType": "Day Meeting",
						"createdBy": "User For Creation",
						"eventManager": "Lisa Rafter",
						"eventArrivalDate": "2024-06-10T11:00:00Z",
						"eventStartDate": "2024-06-10T11:00:00Z",
						"eventEndDate": "2024-06-10T18:30:00Z",
						"eventDepartureDate": "2024-06-10T18:30:00Z",
						"eventDayNumbers": 0,
						"eventNightNumbers": 0,
						"personID": 2,
						"personName": "Robert Magnussen",
						"companyID": 22,
						"companyName": "Company two",
						"meetingRoomID": 6,
						"meetingRoomName": "Room six"
					}
				]
			}
		]
	`)
}
