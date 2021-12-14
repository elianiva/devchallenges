package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Response struct {
	Message  string `json:"message"`
	Filename string `json:"file_url,omitempty"`
}

// writeJson is a wrapper for `w.Write()` to write JSON response
func writeJson(w http.ResponseWriter, v Response) {
	resp, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*") // prevent CORS

	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, Response{Message: "Maximum file size is 10MB"})
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	id, err := gonanoid.New(8)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileExt := path.Ext(handler.Filename)
	fileName := id + fileExt
	f, err := os.OpenFile("static/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		writeJson(w, Response{Message: "Server failed to open file."})
		return
	}
	defer f.Close()
	io.Copy(f, file)

	w.Header().Set("Content-Type", "application/json")
	writeJson(w, Response{
		Message:  "File has been successfully uploaded",
		Filename: fileName,
	})
}
