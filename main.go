package main

import (
	"fmt"
	"net/http"

	"onursahin.dev/vuegolith/ui"
)

func main() {
	fmt.Println("Hallo Welt!")
	assets, _ := ui.Assets()

	// Use the file system to serve static files
	fs := http.FileServer(http.FS(assets))
	http.Handle("/", http.StripPrefix("/", fs))

	port := ":8080"
    println("Server läuft auf http://localhost" + port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        panic(err)
    }
}
