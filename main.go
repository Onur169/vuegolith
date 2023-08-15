package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"onursahin.dev/vuegolith/api"
	"onursahin.dev/vuegolith/routes"
	"onursahin.dev/vuegolith/ui"
	"onursahin.dev/vuegolith/utils"
)

const INSTALLED_PATH = "/usr/local/bin/vuegolith"

func main() {
	var isSecure bool
	flag.BoolVar(&isSecure, "secure", true, "Set to true for secure mode")
	flag.Parse()

	intro := `🧑‍💻️ Welcome to Vuegolith`

	wd, err := os.Executable()
	if err != nil {
		panic(err)
	}

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
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	router := routes.SetupRoutes()

	corsHandler := api.CorsMiddleware(router)
	http.Handle(routes.ApiPrefix + "/", corsHandler)

	certFile := "/etc/vuegolith/ssl/server.crt"
	keyFile := "/etc/vuegolith/ssl/server.key"
	port := "8484"
	securePort := "443"

	var servedUrl string
	if isSecure {
		servedUrl = "https://vuegolith.local/"
	} else {
		servedUrl = "http://localhost:" + port
	}

	fmt.Println(intro)
	fmt.Println("Webserver has started and is available on " + servedUrl)
	fmt.Print("")
	fmt.Println("Exec Directory: " + wd)

	if isSecure {
		err = http.ListenAndServeTLS(":" + securePort, certFile, keyFile, nil)
	} else {
		err = http.ListenAndServe(":" + port, nil)
	}

	if err != nil {
		panic(err)
	}

}
