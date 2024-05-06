package domain

import "errors"

var ErrInvalidName = errors.New("patient name cannot be empty")

type Patient struct {
	name              string
	remainingSessions int
}

func CreatePatient(name string, remainingSessions int) (*Patient, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	return &Patient{
		name:              name,
		remainingSessions: remainingSessions,
	}, nil

}

func (p Patient) GetName() string {
	return p.name
}

func (p Patient) GetRemainingSessions() int {
	return p.remainingSessions
}