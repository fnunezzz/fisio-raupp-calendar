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
	DisplayCalendar() ([]domain.Appointment ,error)
}

type calendarService struct{}

func NewCalendarService() CalendarService {
	return &calendarService{}
}

func (c *calendarService) DisplayCalendar() ([]domain.Appointment, error) {
	ctx := context.Background()
	tokenService := NewTokenService()
	token, _, err := tokenService.GenerateToken()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("token is nil")
	}

	credentialsService := NewCredentialsService()
	credentials, err := credentialsService.LoadCredentials()
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(credentials, calendar.CalendarReadonlyScope)
	if err != nil {
		errMsg := fmt.Sprintf("unable to parse client secret file to config: %v", err)
		return nil, errors.New(errMsg)
	}

	client := config.Client(context.Background(), token)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		errMsg := fmt.Sprintf("unable to retrieve Calendar client: %v", err)
		return nil, errors.New(errMsg)
	}
	
	tomorrow := time.Now().AddDate(0, 0, 1)
	start := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 0, 1, 0, 0, tomorrow.Location()).Format(time.RFC3339)
	end := time.Date(tomorrow.Year(), tomorrow.Month(), tomorrow.Day(), 23, 0, 0, 0, tomorrow.Location()).Format(time.RFC3339)


	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(start).TimeMax(end).MaxResults(100).OrderBy("startTime").Do()
	if err != nil {
		errMsg := fmt.Sprintf("unable to retrieve next ten of the user's events: %v", err)
		return nil, errors.New(errMsg)
	}
	if len(events.Items) == 0 {
		return nil, errors.New("no upcoming events found")
	} 

	var appointments []domain.Appointment 
	appointments = []domain.Appointment{}

	for _, item := range events.Items {
		date := item.Start.DateTime
		appointment := strings.Split(item.Summary, "-")
		if len(appointment) != 2 {
			continue
		}
		name := strings.TrimSpace(appointment[0])
		sessions, err := strconv.Atoi(strings.TrimSpace(appointment[1]))
		if err != nil {
			sessions = 0
		}
		d, err := domain.CreateAppointment(name, date, sessions)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, *d)
	}
	return appointments, nil
	
}