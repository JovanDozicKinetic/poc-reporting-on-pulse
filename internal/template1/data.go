package template1

import (
	"encoding/json"
	"log"
)

func GetData() ([]EventData, []CateringType) {

	jsonData := simulateEventsEndpointCall()
	events := unmarshalEventData(jsonData)

	jsonData = simulateCateringTypesEndpointCall()
	cateringTypes := unmarshalCateringTypeData(jsonData)

	return events, cateringTypes
}

func simulateEventsEndpointCall() []byte {
	return []byte(`
	[
		{
			"cateringType": "Mid Day",
			"confirmed": 65,
			"eventDate": "2024-02-05T00:00:00Z",
			"eventId": 14,
			"provisional": 36,
			"total": 101
		},
		{
			"cateringType": "PM",
			"confirmed": 86,
			"eventDate": "2024-02-09T00:00:00Z",
			"eventId": 15,
			"provisional": 3,
			"total": 89
		},
		{
			"cateringType": "AM",
			"confirmed": 37,
			"eventDate": "2024-03-05T00:00:00Z",
			"eventId": 9,
			"provisional": 26,
			"total": 63
		},
		{
			"cateringType": "AM",
			"confirmed": 85,
			"eventDate": "2024-03-19T00:00:00Z",
			"eventId": 10,
			"provisional": 4,
			"total": 89
		},
		{
			"cateringType": "PM",
			"confirmed": 53,
			"eventDate": "2024-04-17T00:00:00Z",
			"eventId": 3,
			"provisional": 2,
			"total": 55
		},
		{
			"cateringType": "Early Morning",
			"confirmed": 31,
			"eventDate": "2024-04-28T00:00:00Z",
			"eventId": 1,
			"provisional": 32,
			"total": 63
		},
		{
			"cateringType": "Mid Day",
			"confirmed": 66,
			"eventDate": "2024-05-08T00:00:00Z",
			"eventId": 13,
			"provisional": 20,
			"total": 86
		},
		{
			"cateringType": "Mid Day",
			"confirmed": 25,
			"eventDate": "2024-05-21T00:00:00Z",
			"eventId": 6,
			"provisional": 30,
			"total": 55
		},
		{
			"cateringType": "Mid Day",
			"confirmed": 6,
			"eventDate": "2024-06-11T00:00:00Z",
			"eventId": 11,
			"provisional": 22,
			"total": 28
		},
		{
			"cateringType": "AM",
			"confirmed": 7,
			"eventDate": "2024-07-05T00:00:00Z",
			"eventId": 5,
			"provisional": 24,
			"total": 31
		},
		{
			"cateringType": "Evening",
			"confirmed": 32,
			"eventDate": "2024-08-04T00:00:00Z",
			"eventId": 7,
			"provisional": 38,
			"total": 70
		},
		{
			"cateringType": "AM",
			"confirmed": 81,
			"eventDate": "2024-08-15T00:00:00Z",
			"eventId": 16,
			"provisional": 39,
			"total": 120
		},
		{
			"cateringType": "Early Morning",
			"confirmed": 53,
			"eventDate": "2024-09-03T00:00:00Z",
			"eventId": 4,
			"provisional": 45,
			"total": 98
		},
		{
			"cateringType": "PM",
			"confirmed": 57,
			"eventDate": "2024-10-09T00:00:00Z",
			"eventId": 2,
			"provisional": 38,
			"total": 95
		},
		{
			"cateringType": "Early Morning",
			"confirmed": 95,
			"eventDate": "2024-10-21T00:00:00Z",
			"eventId": 8,
			"provisional": 2,
			"total": 97
		},
		{
			"cateringType": "PM",
			"confirmed": 89,
			"eventDate": "2024-12-10T00:00:00Z",
			"eventId": 12,
			"provisional": 49,
			"total": 138
		}
	]`)
}

func simulateCateringTypesEndpointCall() []byte {
	return []byte(`
	[
		{
			"catering_type": "Early Morning",
			"id": 101
		},
		{
			"catering_type": "AM",
			"id": 203
		},
		{
			"catering_type": "Mid Day",
			"id": 345
		},
		{
			"catering_type": "PM",
			"id": 4851
		},
		{
			"catering_type": "Evening",
			"id": 545
		}
	]`)
}

func unmarshalEventData(jsonData []byte) []EventData {
	var events []EventData
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	return events
}

func unmarshalCateringTypeData(jsonData []byte) []CateringType {
	var events []CateringType
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	return events
}
