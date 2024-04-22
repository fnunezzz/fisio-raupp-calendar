package domain

import "errors"

var ErrInvalidName = errors.New("patient name cannot be empty")

type Patient struct {
	Name              string
	RemainingSessions int
}

func CreatePatient(name string, remainingSessions int) (*Patient, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &Patient{
		Name:              name,
		RemainingSessions: remainingSessions,
	}, nil

}