package main

import (
	"fmt"
	"image-uploader/handlers"
	"log"
	"net/http"
	"os"
)

const PORT = "8000"

func main() {
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/upload", handlers.UploadImage)

	fmt.Println("Listening on PORT: " + PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatalf("Failed to start server. Reason: %v", err)
		os.Exit(1)
	}
}
