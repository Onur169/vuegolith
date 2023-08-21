package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"onursahin.dev/vuegolith/api/middleware"
	"onursahin.dev/vuegolith/env"
	"onursahin.dev/vuegolith/routes"
	"onursahin.dev/vuegolith/ui"
	"onursahin.dev/vuegolith/utils"
)

const INSTALLED_PATH = "/usr/local/bin/vuegolith"

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

	env.Init(wd)
	fmt.Println()
	e := env.Load()
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

	uploadFs := middleware.FileServerWithCors(http.Dir(homeDir + "/" + utils.GetUploadsDirName()))
	http.Handle(e.UploadsPath, http.StripPrefix(e.UploadsPath, uploadFs))

	router := routes.SetupRoutes(e.APIPrefix)
	ch := middleware.RouterWithCors(router)
	http.Handle(e.APIPrefix+"/", ch)

	var servedUrl string
	if isSecure {
		servedUrl = e.SecureURL
	} else {
		servedUrl = e.NonSecureURL + e.NonSecurePort
	}

	fmt.Println("✅ Webserver has started and is available on " + servedUrl)
	fmt.Println("ⓘ Exec Directory: " + wd)

	if isSecure {
		err = http.ListenAndServeTLS(":443", e.CertFile, e.KeyFile, nil)
	} else {
		err = http.ListenAndServe(":"+e.NonSecurePort, nil)
	}

	if err != nil {
		panic(err)
	}

}
