package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	url = "https://ifconfig.co/json"
)

func routerpage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<a href=\"/api\">Route to API</a>")
}

func api(w http.ResponseWriter, r *http.Request) {
	req, err := http.Get(url)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	ipStruct := struct {
		IP       string `json:"ip"`
		Location string `json:"country"`
		City     string `json:"city"`
		TimeZone string `json:"time_zone"`
	}{}

	err = json.NewDecoder(req.Body).Decode(&ipStruct)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	fmt.Fprintf(w, "Your IP address is: %s\n", ipStruct.IP)
	fmt.Fprintf(w, "Your location is: %s\n", ipStruct.Location)
	fmt.Fprintf(w, "Your city is: %s\n", ipStruct.City)
	fmt.Fprintf(w, "Your timezone is: %s\n", ipStruct.TimeZone)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api", api)
	mux.HandleFunc("/", routerpage)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Printf("Server listening on :8080\n")
	fmt.Println("http://localhost:8080")
	server.ListenAndServe()
}
