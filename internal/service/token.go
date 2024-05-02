package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/oauth2"
)

type TokenService interface {
	readToken() (*oauth2.Token, error)
	writeToken(*oauth2.Token) error
	GenerateToken() (*oauth2.Token, error)
	CheckToken() error
	getTokenFromWeb(*oauth2.Config)
	WriteTokenUsingAuthCode(authCode string)
}

type tokenService struct {
	folderName string
	fileName  string
	config *oauth2.Config
}

func NewTokenService() TokenService {
	return &tokenService{
		folderName: FolderName,
		fileName: "token.json",
	}
}

func (t *tokenService) CheckToken() error {
	_, err := os.Stat(t.folderName + "/" + t.fileName)
	if err != nil {
		return err
	}
	return nil
}

func (t *tokenService) GenerateToken() (*oauth2.Token, error) {
	googleAuth := NewGoogleAuthenticationService()
	credentialsService := NewCredentialsService()
	
	credentials, err := credentialsService.LoadCredentials()
	if err != nil {
		return nil, err
	}
	oauth, err := googleAuth.Auth(credentials)

	if err != nil {
		return nil, err
	}
	tok, err := t.readToken()
	// no token on file, so time to make a new one
	if err != nil {
			t.getTokenFromWeb(oauth)
			tok, err = t.readToken()
			if err != nil {
				return nil, err
			}

	}
	return tok, nil

}

func (t *tokenService) writeToken(token *oauth2.Token) error {
	file := t.folderName + "/" + t.fileName
	f, err := os.Create(file)
	if err != nil {
			return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

func (t *tokenService) getTokenFromWeb(config *oauth2.Config) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Abra o seguinte LINK: \n%v\n", authURL)
	
	i := 0
	for {
		if i == 120 {
			log.Fatalf("Unable to read authorization code: %v", errors.New("timeout"))
		}
		err := t.CheckToken()
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
		i++

	}


}

func (t *tokenService) WriteTokenUsingAuthCode(authCode string) {
	googleAuth := NewGoogleAuthenticationService()
	credentialsService := NewCredentialsService()
	
	credentials, err := credentialsService.LoadCredentials()
	if err != nil {
		log.Fatalf("Unable to load credentials: %v", err)
	}
	oauth, err := googleAuth.Auth(credentials)
	if err != nil {
		log.Fatalf("Unable to authenticate: %v", err)
	}
	tok, err := oauth.Exchange(context.TODO(), authCode)
	if err != nil {
			log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	time.Sleep(3 * time.Second) // enough time to http handle do it's job of redirection
	err = t.writeToken(tok)
	if err != nil {
			log.Fatalf("Unable to write token: %v", err)
	}
}

func (t *tokenService) readToken() (*oauth2.Token, error) {
	// Reads the token in the hidden folder
	file := t.folderName + "/" + t.fileName
	f, err := os.Open(file)
	if err != nil {
			return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
