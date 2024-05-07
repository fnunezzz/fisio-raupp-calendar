package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	domain "github.com/fnunezzz/fisio-raupp-calendar/internal/domain/appointments"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type CalendarService interface {
	GetNextDayAppointments() ([]domain.Appointment, time.Time, error)
}

type calendarService struct{}

func NewCalendarService() CalendarService {
	return &calendarService{}
}

func (c *calendarService) GetNextDayAppointments() ([]domain.Appointment, time.Time, error) {
	ctx := context.Background()
	tokenService := NewTokenService()
	token, _, err := tokenService.GenerateToken()
	if err != nil {
		return nil, time.Time{}, err
	}
	if token == nil {
		return nil, time.Time{}, errors.New("token is nil")
	}

	credentialsService := NewCredentialsService()
	credentials, err := credentialsService.LoadCredentials()
	if err != nil {
		return nil, time.Time{}, err
	}

	config, err := google.ConfigFromJSON(credentials, calendar.CalendarReadonlyScope)
	if err != nil {
		errMsg := fmt.Sprintf("unable to parse client secret file to config: %v", err)
		return nil, time.Time{}, errors.New(errMsg)
	}

	client := config.Client(context.Background(), token)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		errMsg := fmt.Sprintf("unable to retrieve Calendar client: %v", err)
		return nil, time.Time{}, errors.New(errMsg)
	}
	
	tomorrow := time.Now().AddDate(0, 0, 1)
	if tomorrow.Weekday() == time.Saturday {
		tomorrow = tomorrow.AddDate(0, 0, 2)
	} else if tomorrow.Weekday() == time.Sunday {
		tomorrow = tomorrow.AddDate(0, 0, 1)
	}
	start := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 1, 0, 0, tomorrow.Location()).Format(time.RFC3339)
	end := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 23, 0, 0, 0, tomorrow.Location()).Format(time.RFC3339)


	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(start).TimeMax(end).MaxResults(100).OrderBy("startTime").Do()
	if err != nil {
		errMsg := fmt.Sprintf("unable to retrieve next ten of the user's events: %v", err)
		return nil, time.Time{}, errors.New(errMsg)
	}
	if len(events.Items) == 0 {
		return nil, time.Time{}, errors.New("no upcoming events found")
	} 

	var appointments []domain.Appointment 
	appointments = []domain.Appointment{}

	for _, item := range events.Items {
		date := item.Start.DateTime
		appointment := strings.Split(item.Summary, "-")
		sessions := 0

		name := strings.TrimSpace(appointment[0])
		if len(appointment) >= 2 {
			sessions, err = strconv.Atoi(strings.TrimSpace(appointment[1]))
			if err != nil {
				sessions = 0
			}
		}
		d, err := domain.CreateAppointment(name, date, sessions)
		if err != nil {
			return nil, time.Time{}, err
		}
		appointments = append(appointments, *d)
	}
	return appointments, tomorrow, nil
	
}