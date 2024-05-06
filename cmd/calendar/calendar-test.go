package main

import (
	"log"

	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)

func main() {
	calendarService := service.NewCalendarService()
	p, err := calendarService.DisplayCalendar()
	if err != nil {
		log.Fatalf("Unable to display calendar: %v", err)
	}

	for _, v := range p {
		log.Printf("Name: %v", v.Patient.GetName())
		log.Printf("RemaingingSessions: %v", v.Patient.GetRemainingSessions())
		log.Printf("Date: %v", v.Date())
		log.Printf("\n")
	}
}