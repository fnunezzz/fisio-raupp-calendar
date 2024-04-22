package service

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)


type Oauth interface {
	Auth([]byte) (*oauth2.Config, error)
}

type googleAuthenticationService struct{}

func NewGoogleAuthenticationService() Oauth {
	return &googleAuthenticationService{}
}

func (g *googleAuthenticationService) Auth(credentials []byte) (*oauth2.Config, error) {
	// ctx := context.Background()
	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(credentials, calendar.CalendarReadonlyScope)
	if err != nil {
			return nil, err
	}
	return config, nil
	// client := config.Client(context.Background(), tok) // token service

	// calendarService, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	// if err != nil {
	// 		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	// }

	// return calendarService
}