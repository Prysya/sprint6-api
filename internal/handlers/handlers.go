package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	maxUploadSize = 10 << 20 // 10MB
	uploadDir     = "temp"
)

func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	htmlFile, err := os.Open("index.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer htmlFile.Close()

	if _, err = io.Copy(w, htmlFile); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(maxUploadSize)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading file content", http.StatusInternalServerError)
		return
	}

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	convertedString := service.ConvertMorseOrText(string(fileData))

	originalExt := filepath.Ext(handler.Filename)
	fileName := time.Now().UTC().Format("02.01.06_15.04.05") + originalExt

	tempDir, err := os.OpenRoot(uploadDir)
	if err != nil {
		fmt.Println(err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer tempDir.Close()

	dst, err := tempDir.Create(fileName)
	if err != nil {
		http.Error(w, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.Write([]byte(convertedString))
	if err != nil {
		http.Error(w, "Error updating file data", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(convertedString))
}
