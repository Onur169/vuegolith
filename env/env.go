package env

import (
	"fmt"
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

func Load() Env {
	return Env{
		UploadsPath:   utils.GetEnvVariable("UPLOADS_PATH", "/uploads/"),
		APIPrefix:     utils.GetEnvVariable("API_PREFIX", "/api"),
		CertFile:      utils.GetEnvVariable("CERT_FILE", "/etc/vuegolith/ssl/server.crt"),
		KeyFile:       utils.GetEnvVariable("KEY_FILE", "/etc/vuegolith/ssl/server.key"),
		NonSecureURL:  utils.GetEnvVariable("NON_SECURE_URL", "http://localhost:"),
		NonSecurePort: utils.GetEnvVariable("NON_SECURE_PORT", "8484"),
		SecureURL:     utils.GetEnvVariable("SECURE_URL", "https://vuegolith.local/"),
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
