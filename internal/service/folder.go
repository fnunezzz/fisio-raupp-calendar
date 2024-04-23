package service

import (
	"errors"
	"os"
	"syscall"
)

var FolderName = ".data"

type folderService struct {
	folderName string
}

type FolderServiceInterface interface {
	CreateFolder() error
}

func NewFolderService() FolderServiceInterface {
	return &folderService{folderName: FolderName}
}

// This functions writes to file in a hidden folder
// It's a basic security measure to avoid exposing the token to the user
// It's VERY basic, but it's better than nothing
// Todo encrypt the folder (?)
func (t *folderService) CreateFolder() error {

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