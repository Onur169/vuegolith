package env

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"onursahin.dev/vuegolith/utils"
)

type Env struct {
	UploadsPath   string
	APIPrefix     string
	CertFile      string
	KeyFile       string
	NonSecureURL  string
	NonSecurePort string
	SecureURL     string
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

func Load() Env {
	return Env{
		UploadsPath:   GetEnvVariable("UPLOADS_PATH", "/uploads/"),
		APIPrefix:     GetEnvVariable("API_PREFIX", "/api"),
		CertFile:      GetEnvVariable("CERT_FILE", "/etc/vuegolith/ssl/server.crt"),
		KeyFile:       GetEnvVariable("KEY_FILE", "/etc/vuegolith/ssl/server.key"),
		NonSecureURL:  GetEnvVariable("NON_SECURE_URL", "http://localhost:"),
		NonSecurePort: GetEnvVariable("NON_SECURE_PORT", "8484"),
		SecureURL:     GetEnvVariable("SECURE_URL", "https://vuegolith.local/"),
	}
}

func Init(wd string) {
	err := godotenv.Load(filepath.Dir(wd) + "/vuegolith.env")

	if err != nil {
		fmt.Println("❌ vuegolith.env not loaded")
		err = godotenv.Load(filepath.Dir(wd) + "/.env")

		if err != nil {
			fmt.Println("❌ Standard .env not loaded")
			err = utils.CreateEmptyVuegolithEnvFile(filepath.Dir(wd))
			if err != nil {
				fmt.Println("❌ Could not create empty vuegolith.env", err)
			}
			fmt.Println("✅ Create empty vuegolith.env")
			return
		}

		fmt.Println("✅ Environment variables loaded (.env)")
		return

	}

	fmt.Println("✅ Environment variables loaded (vuegolith.env)")
}
