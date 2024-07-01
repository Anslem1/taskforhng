package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("visitor_name")
	ipAddr := "127.0.0.1"
	query = strings.Trim(query, `"`)

	greeting := fmt.Sprintf("Hello, %v! The temperature is 11 degrees Celsius in Mark.", query)

	fmt.Println(query)

	response := Response{
		ClientIP: ipAddr,
		Location: "New York",
		Greeting: greeting,
	}

	// Marshal response object to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error marshaling JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/api/hello", helloHandler)
	
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
