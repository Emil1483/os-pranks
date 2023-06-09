package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func buildFileHandler(fileName string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

		w.Write(content)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func main() {
	http.HandleFunc("/server.exe", buildFileHandler("server.exe"))
	http.HandleFunc("/client.exe", buildFileHandler("client.exe"))
	http.HandleFunc("/install.bat", buildFileHandler("install.bat"))

	port := getEnv("PORT", "8080")

	fmt.Println("Starting server at port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
