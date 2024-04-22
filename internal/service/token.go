package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"

	"golang.org/x/oauth2"
)

type TokenService interface {
	writeToFile() error
	readToken() (*oauth2.Token, error)
	writeToken(*oauth2.Token) error
	GenerateToken(*oauth2.Config) (*oauth2.Token, error)
	getTokenFromWeb(*oauth2.Config) *oauth2.Token
}

type tokenService struct {
	folderName string
	fileName  string
}

func NewTokenService() TokenService {
	return &tokenService{
		folderName: ".data",
		fileName: "token.json",
	}
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

// This functions writes to file in a hidden folder
// It's a basic security measure to avoid exposing the token to the user
// It's VERY basic, but it's better than nothing
// Todo encrypt the token (?)
func (t *tokenService) writeToFile() error {

    // Create the hidden folder
    err := os.Mkdir(t.folderName, 0755) // 0755 sets the folder permissions
    if err != nil {
		if !errors.Is(err, os.ErrExist) {
			return err
		}
    }

	// Convert folderName to UTF-16 encoded pointer
	folderNamePtr, err := syscall.UTF16PtrFromString(t.folderName)
	if err != nil {
		return err
	}

	// Get the file attributes
	attrs, err := syscall.GetFileAttributes(folderNamePtr)
	if err != nil {
		return err
	}

	// Set the hidden attribute
	err = syscall.SetFileAttributes(folderNamePtr, attrs|syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}
	
	return nil
}