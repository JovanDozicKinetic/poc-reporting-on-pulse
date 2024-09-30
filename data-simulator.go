package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"sort"
	"time"
)

func SimulateEndpointCall() []byte {
	var events []map[string]interface{}

	cateringTypes := []string{"Early Morning", "AM", "Mid Day", "PM", "Evening"}

	for i := 1; i <= 16; i++ {
		confirmed := rand.Intn(100)
		provisional := rand.Intn(50)
		event := map[string]interface{}{
			"cateringType": cateringTypes[rand.Intn(5)],
			"confirmed":    confirmed,
			"eventDate":    time.Date(2024, time.Month(rand.Intn(12)+1), rand.Intn(28)+1, 0, 0, 0, 0, time.UTC).Format("2006-01-02T15:04:05Z"),
			"eventId":      i,
			"provisional":  provisional,
			"total":        confirmed + provisional,
		}
		events = append(events, event)
	}

	sort.Slice(events, func(i, j int) bool {
		dateI, _ := time.Parse("2006-01-02T15:04:05Z", events[i]["eventDate"].(string))
		dateJ, _ := time.Parse("2006-01-02T15:04:05Z", events[j]["eventDate"].(string))

		if dateI.Equal(dateJ) {
			return events[i]["cateringType"].(string) < events[j]["cateringType"].(string)
		}
		return dateI.Before(dateJ)
	})

	jsonData, err := json.Marshal(events)
	if err != nil {
		log.Fatalln("Error marshaling JSON:", err)
	}

	return jsonData
}

func UnmarshalEventData(jsonData []byte) []EventData {
	var events []EventData
	err := json.Unmarshal(jsonData, &events)
	if err != nil {
		log.Fatalln("JSON error: ", err)
	}
	log.Println("JSON data converted")
	return events
}
