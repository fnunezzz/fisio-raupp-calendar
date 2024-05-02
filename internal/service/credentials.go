package service

import (
	"errors"
	"os"
)

var (
	ErrCredentialsNotFound = errors.New("credentials file not found")
)

type CredentialsService interface {
	GetCredentials() []byte
	LoadCredentials() ([]byte, error)
}

type configService struct {
	crendetials []byte
}

func NewCredentialsService() CredentialsService {
	return &configService{}
}

func (c *configService) GetCredentials() []byte {
	return c.crendetials
}

func (c *configService) LoadCredentials() ([]byte, error) {
	credentials, err := os.ReadFile(FolderName + "/credentials.json")
	if err != nil {
		return nil, ErrCredentialsNotFound
	}

	c.crendetials = credentials

	return c.crendetials, nil
}