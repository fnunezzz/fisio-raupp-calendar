package main

import (
	"log"
	"time"

	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)

func main() {
	clientService := service.NewClientService()
	go clientService.StartClient()
	defer clientService.StopClient()
	tokenService := service.NewTokenService()
	_, s, err := tokenService.GenerateToken()
	if err != nil {
		log.Fatalf("Unable to generate token: %v", err)
	}

	if s != "" {
		log.Printf("Please go to the following link in your browser then type the authorization code: %v", s)
		for i := 0; i < 60; i++ {
			err := tokenService.CheckToken()
			if err == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
	}



}