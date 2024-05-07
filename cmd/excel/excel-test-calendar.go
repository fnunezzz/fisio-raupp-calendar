package main

import (
	"fmt"
	"log"

	"github.com/fnunezzz/fisio-raupp-calendar/internal/service"
)


func main() {
	xlsxService := service.NewXlsxService()

	calendarService := service.NewCalendarService()
	p, t, err := calendarService.GetNextDayAppointments()
	if err != nil {
		log.Fatalf("Unable to display calendar: %v", err)
	}

	var dtos []service.Input
	dtos = []service.Input{}
	for _, v := range p {
		fmt.Println(v.Patient.GetName(), v.GetTime())
		dto := service.Input{
			Text: v.GetPatientNameAndSessions(),
			Time: v.GetTime(),
		}
		dtos = append(dtos, dto)
	}

	
	xlsxService.GenerateXlsxReport(dtos, t)
	
}