package utils

import (
	"os"
	"path/filepath"
)

func GetUploadsDirName() string {
	return "vuegolith-uploads"
}

func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func CreateVuegolithUploadsDir() (string, error) {
	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	uploadsDir := filepath.Join(homeDir, GetUploadsDirName())

	// Überprüfe, ob der Ordner bereits existiert
	_, err = os.Stat(uploadsDir)
	if os.IsNotExist(err) {
		// Ordner existiert nicht, also erstellen
		err = os.Mkdir(uploadsDir, 0755)
		if err != nil {
			return "", err
		}
	}

	return uploadsDir, nil
}
