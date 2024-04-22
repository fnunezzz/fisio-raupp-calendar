package domain

type Appointment struct {
	Patient *Patient
	Date    string
}

func CreateAppointment(name string, date string, remainingSessions int) (*Appointment, error) {
	patient, err := CreatePatient(name, remainingSessions)
	if err != nil {
		return nil, err
	}
	return &Appointment{
		Patient: patient,
		Date:    date,
	}, nil
}