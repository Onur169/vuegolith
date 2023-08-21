package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
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

	_, err = os.Stat(uploadsDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(uploadsDir, 0755)
		if err != nil {
			return "", err
		}
	}

	return uploadsDir, nil
}

func IsPathSafe(path string) bool {
	return !strings.Contains(path, "../")
}

func GetEnvVariable(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		fmt.Println("❌ Env " + key + " has not loaded. Use default value.")
		return defaultValue
	}
	fmt.Println("✅ Env " + key + " has loaded")
	return value
}

func ClearScreen() {
	clearCmd := ""
	switch runtime.GOOS {
	case "linux", "darwin":
		clearCmd = "clear"
	case "windows":
		clearCmd = "cls"
	default:
		fmt.Println("Unsupporteds Operating System")
		return
	}

	cmd := exec.Command(clearCmd)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func CreateEmptyVuegolithEnvFile(path string) error {
	filename := "vuegolith.env"

	file, err := os.Create(path + "/" + filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("")
	if err != nil {
		return err
	}

	return nil
}
