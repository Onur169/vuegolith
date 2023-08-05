package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"onursahin.dev/vuegolith/api"
	"onursahin.dev/vuegolith/endpoints"
	"onursahin.dev/vuegolith/ui"
	"onursahin.dev/vuegolith/utils"

	"github.com/gorilla/mux"
)

const INSTALLED_PATH = "/usr/local/bin/vuegolith"

func main() {
	var isSecure bool
	flag.BoolVar(&isSecure, "secure", true, "Set to true for secure mode")
	flag.Parse()

	intro := `🧑‍💻️ Welcome to Vuegolith`

	// wd, err := os.Getwd()
	wd, err := os.Executable()
	if err != nil {
		panic(err)
	}

	_, err = utils.CreateVuegolithUploadsDir()
	if err != nil {
		panic(err)
	}

	assets, _ := ui.Assets()

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", http.StripPrefix("/", fs))

	// Create a new Router from Gorilla Mux
	router := mux.NewRouter()

	// Define API endpoints with Gorilla Mux
	router.HandleFunc("/api/log", endpoints.HandleLogGet).Methods("GET")
	router.HandleFunc("/api/log", endpoints.HandleLogPost).Methods("POST")
	router.HandleFunc("/api/uploads", endpoints.HandleUpload).Methods("POST")
	router.HandleFunc("/api/uploads", endpoints.HandleListUploads).Methods("GET")
	router.HandleFunc("/api/uploads", endpoints.HandleDeleteUpload).Methods("DELETE") 

	// Determine home dir
	homeDir, err := utils.GetHomeDir()
	if err != nil {
		panic(err)
	}

	// // Serve files from the home directory under the "/uploads/" endpoint
	fs = api.FileServerWithCors(http.Dir(homeDir + "/" + utils.GetUploadsDirName()))
	http.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	corsHandler := api.CorsMiddleware(router)
	http.Handle("/api/", corsHandler)

	// Pfade zu deinen selbstsignierten Zertifikatsdateien
	// port := "8484"

	//err = http.ListenAndServe(":"+port, nil)
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
