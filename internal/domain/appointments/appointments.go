package domain

import (
	"fmt"
	"time"
)

type Appointment struct {
	Patient *Patient
	date    time.Time
}

func CreateAppointment(name string, date string, remainingSessions int) (*Appointment, error) {
	patient, err := CreatePatient(name, remainingSessions)
	if err != nil {
		return nil, err
	}
	d, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, err
	}
	return &Appointment{
		Patient: patient,
		date:    d,
	}, nil
}

func (a Appointment) Date() time.Time {
	return a.date
}

func (a Appointment) GetPatientNameAndSessions() string {
	s := fmt.Sprintf("%s - %d", a.Patient.GetName(), a.Patient.GetRemainingSessions())
	return s
}

func (a Appointment) GetTime() string {
	hourString := a.date.Format("15:04")
	return hourString
}