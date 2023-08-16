package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"onursahin.dev/vuegolith/api"
	"onursahin.dev/vuegolith/routes"
	"onursahin.dev/vuegolith/ui"
	"onursahin.dev/vuegolith/utils"
)

const INSTALLED_PATH = "/usr/local/bin/vuegolith"

func processEnv(wd string) {
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

func main() {
	utils.ClearScreen()
	intro := `🧑‍💻️ Welcome to Vuegolith`
	fmt.Println(intro)

	var isSecure bool
	flag.BoolVar(&isSecure, "secure", true, "Set to true for secure mode")
	flag.Parse()

	wd, err := os.Executable()
	if err != nil {
		panic(err)
	}

	processEnv(wd)

	fmt.Println()
	uploadsPath := utils.GetEnvVariable("UPLOADS_PATH", "/uploads/")
	apiPrefix := utils.GetEnvVariable("API_PREFIX", "/api")
	certFile := utils.GetEnvVariable("CERT_FILE", "/etc/vuegolith/ssl/server.crt")
	keyFile := utils.GetEnvVariable("KEY_FILE", "/etc/vuegolith/ssl/server.key")
	nonSecureURL := utils.GetEnvVariable("NON_SECURE_URL", "http://localhost:")
	nonSecurePort := utils.GetEnvVariable("NON_SECURE_PORT", "8484")
	secureURL := utils.GetEnvVariable("SECURE_URL", "https://vuegolith.local/")
	fmt.Println()

	_, err = utils.CreateVuegolithUploadsDir()
	if err != nil {
		panic(err)
	}

	homeDir, err := utils.GetHomeDir()
	if err != nil {
		panic(err)
	}

	assets, _ := ui.Assets()
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", http.StripPrefix("/", fs))

	fs = api.FileServerWithCors(http.Dir(homeDir + "/" + utils.GetUploadsDirName()))
	http.Handle(uploadsPath, http.StripPrefix(uploadsPath, fs))

	router := routes.SetupRoutes(apiPrefix)
	corsHandler := api.CorsMiddleware(router)
	http.Handle(apiPrefix+"/", corsHandler)

	var servedUrl string
	if isSecure {
		servedUrl = secureURL
	} else {
		servedUrl = nonSecureURL + nonSecurePort
	}

	fmt.Println("✅ Webserver has started and is available on " + servedUrl)
	fmt.Println("ⓘ Exec Directory: " + wd)

	if isSecure {
		err = http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(":"+nonSecurePort, nil)
	}

	if err != nil {
		panic(err)
	}

}
