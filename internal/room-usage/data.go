package roomusage

import (
	"encoding/json"
	"log"
)

func GetEventDataWithoutMeetingRooms() []Dates {
	jsonData := simulateDatesData()
	events := unmarshalData(jsonData)
	return events
}

func unmarshalData(jsonData []byte) []Dates {
	var events []Dates
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Println("JSON error: ", err)
	}
	return events
}

func simulateDatesData() []byte {
	return []byte(`
		[
			{
				"date": "2021-07-01T00:00:00Z",
				"bookings": [
					{
						"reference": "44444",
						"name": "Example Event in Room 1",
						"dayCurrent": 1,
						"daysTotal": 3,
						"roomName": "Room 1",
						"startTime": null,
						"endTime": null,
						"number": 10,
						"layout": "Boardroom",
						"roomNotes": "This is a note for the room",
						"equipment": [
							{
								"name": "Projector",
								"quantity": 10,
								"notes": "This is a note for the equipment"
							},
							{
								"name": "Laptop",
								"quantity": 10,
								"notes": "This is a note for the equipment"
							},
							{
								"name": "Flipchart",
								"quantity": 10,
								"notes": "This is a note for the equipment"
							}
						],
						"nextUse": null,
						"isChangedSince": false
					},
					{
						"reference": "12345",
						"name": "Example Event in Room 2",
						"dayCurrent": 1,
						"daysTotal": 2,
						"roomName": "Room 2",
						"startTime": "2021-07-01T08:00:00Z",
						"endTime": "2021-07-01T10:00:00Z",
						"number": 5,
						"layout": "Theatre",
						"roomNotes": "Another note for the room",
						"equipment": [
							{
								"name": "Microphone",
								"quantity": 2,
								"notes": "Wireless mic"
							}
						],
						"nextUse": "2021-07-02T08:00:00Z",
						"isChangedSince": true
					}
				]
			},
			{
				"date": "2021-07-02T00:00:00Z",
				"bookings": [
					{
						"reference": "44444",
						"name": "Example Event in Room 1",
						"dayCurrent": 2,
						"daysTotal": 3,
						"roomName": "Room 1",
						"startTime": null,
						"endTime": null,
						"number": 20,
						"layout": "Classroom",
						"roomNotes": "Second-day note",
						"equipment": [
							{
								"name": "Projector",
								"quantity": 5,
								"notes": "Slightly different quantity"
							},
							{
								"name": "Laptop",
								"quantity": 5,
								"notes": "Still needed"
							}
						],
						"nextUse": null,
						"isChangedSince": false
					},
					{
						"reference": "12345",
						"name": "Example Event in Room 2",
						"dayCurrent": 2,
						"daysTotal": 2,
						"roomName": "Room 2",
						"startTime": "2021-07-02T08:00:00Z",
						"endTime": "2021-07-02T10:00:00Z",
						"number": 5,
						"layout": "Theatre",
						"roomNotes": "Second day for Room 2",
						"equipment": [
							{
								"name": "Microphone",
								"quantity": 2,
								"notes": "Wireless mic"
							}
						],
						"nextUse": "2021-07-03T08:00:00Z",
						"isChangedSince": false
					}
				]
			},
			{
				"date": "2021-07-03T00:00:00Z",
				"bookings": [
					{
						"reference": "44444",
						"name": "Example Event in Room 1",
						"dayCurrent": 3,
						"daysTotal": 3,
						"roomName": "Room 1",
						"startTime": "2021-07-03T09:00:00Z",
						"endTime": null,
						"number": 15,
						"layout": "U-Shape",
						"roomNotes": "Third day note",
						"equipment": [
							{
								"name": "Projector",
								"quantity": 1,
								"notes": "Day 3 projector note"
							}
						],
						"nextUse": "2021-07-10T09:00:00Z",
						"isChangedSince": false
					},
					{
						"reference": "77777",
						"name": "Another Example Event",
						"dayCurrent": 1,
						"daysTotal": 1,
						"roomName": "Conference Hall",
						"startTime": "2021-07-03T11:00:00Z",
						"endTime": "2021-07-03T13:00:00Z",
						"number": 50,
						"layout": "Banquet",
						"roomNotes": "Single day event",
						"equipment": [
							{
								"name": "Speaker system",
								"quantity": 2,
								"notes": "For a large hall"
							}
						],
						"nextUse": null,
						"isChangedSince": true
					}
				]
			},
			{
				"date": "2021-07-04T00:00:00Z",
				"bookings": [
					{
						"reference": "99999",
						"name": "All Day Meeting",
						"dayCurrent": 1,
						"daysTotal": 1,
						"roomName": "Room 3",
						"startTime": null,
						"endTime": null,
						"number": 12,
						"layout": "Square",
						"roomNotes": "No notes",
						"equipment": [],
						"nextUse": null,
						"isChangedSince": false
					},
					{
						"reference": "55555",
						"name": "Partial Day Workshop",
						"dayCurrent": 1,
						"daysTotal": 1,
						"roomName": "Room 4",
						"startTime": "2021-07-04T10:30:00Z",
						"endTime": "2021-07-04T12:30:00Z",
						"number": 7,
						"layout": "Theatre",
						"roomNotes": "Focus on new features",
						"equipment": [
							{
								"name": "Whiteboard",
								"quantity": 1,
								"notes": "Large whiteboard"
							}
						],
						"nextUse": "2021-07-05T09:00:00Z",
						"isChangedSince": true
					}
				]
			}
		]
	`)
}
