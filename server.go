package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

func requests(w http.ResponseWriter, r *http.Request) {
	_, errorURL := template.ParseFiles(r.URL.Path[1:])
	if errorURL != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

}
func manageRender(w http.ResponseWriter, r *http.Request) {
	tmpl, errorURL := template.ParseFiles(r.URL.Path[1:] + ".html")
	//page, _ := loadPage(r.URL.Path[1:])
	if errorURL != nil {
		http.Error(w, "Invalid request", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
	//fmt.Fprintf(w, string(page.Body))
}
func main() {
	errorEnv := godotenv.Load()
	if errorEnv != nil {
		fmt.Printf("Error Loading the env file: %v\n", errorEnv)
	}

	//mux := http.NewServeMux()
	http.HandleFunc("/req", requests)
	http.HandleFunc("/", manageRender)
	server := &http.Server{
		Addr: ":3113",
		//Handler:      mux,
		ReadTimeout:  10 * time.Since(time.Now()),
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is running on http://localhost:3113")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}
