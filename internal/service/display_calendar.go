package service

import (
	"fmt"
	"log"
	"time"
)

type CalendarService interface {
	DisplayCalendar()
}

type calendarService struct{}

func NewCalendarService() CalendarService {
	return &calendarService{}
}

func (c *calendarService) DisplayCalendar() {
	// todo remove depencyy from service
	auth := NewGoogleAuthenticationService()
	calendarService := auth.GoogleAuthentication()

	tomorrow := time.Now().AddDate(0, 0, 1)
	midnight := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 0, 0, 0, tomorrow.Location()).Format(time.RFC3339)
	tomorrowFormatted := tomorrow.Format(time.RFC3339)

	events, err := calendarService.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(midnight).TimeMax(tomorrowFormatted).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
		}
	}
}