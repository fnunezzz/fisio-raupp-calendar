package main

import (
	"fmt"
	"log"

	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)

func main() {
	tokenService := service.NewTokenService()
	tok, err := tokenService.GenerateToken()
	if err != nil {
		log.Fatalf("Unable to generate token: %v", err)
	}
	fmt.Println(tok)
}