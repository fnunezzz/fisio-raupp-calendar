package main

import (
	"fmt"
	"log"

	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)

func main() {

	credentialsService := service.NewCredentialsService()

	credentials, err := credentialsService.LoadCredentials()
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	googleAuth := service.NewGoogleAuthenticationService()
	oauth, err := googleAuth.Auth(credentials)

	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	tokenService := service.NewTokenService()
	tok, err := tokenService.GenerateToken(oauth)
	if err != nil {
		log.Fatalf("Unable to generate token: %v", err)
	}
	fmt.Println(tok)
}