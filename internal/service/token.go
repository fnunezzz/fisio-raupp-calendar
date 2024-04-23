package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)

type TokenService interface {
	readToken() (*oauth2.Token, error)
	writeToken(*oauth2.Token) error
	GenerateToken(*oauth2.Config) (*oauth2.Token, error)
	CheckToken() error
	getTokenFromWeb(*oauth2.Config) *oauth2.Token
}

type tokenService struct {
	folderName string
	fileName  string
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

func (t *tokenService) GenerateToken(config *oauth2.Config) (*oauth2.Token, error) {
	tok, err := t.readToken()
	// no token on file, so time to make a new one
	// TODO display message to user
	if err != nil {
			tok = t.getTokenFromWeb(config)
			err := t.writeToken(tok)
			if (err != nil) {
				return nil, err
			}
			return tok, nil
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

func (t *tokenService) getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
			"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
			log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
			log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
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

