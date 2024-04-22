package domain

import "errors"

var (
	ErrEmailNotConfigured       = errors.New("email not configured")
	ErrTokenNotConfigured = errors.New("token not configured")
)

type user struct {
	Email       string
	Token string
}

type User interface {
	ConfigureToken(string)
	ConfigureEmail(string)
	ValidateUser() error
}

func NewConfigService() User {
	return &user{}
}

func (c *user) ConfigureEmail(email string) {
	c.Email = email
}

func (c *user) ConfigureToken(token string) {
	c.Token = token
}

func (c *user) ValidateUser() error {
	if c.Email == "" {
		return ErrEmailNotConfigured
	}
	if c.Token == "" {
		return ErrTokenNotConfigured
	}
	return nil
}
