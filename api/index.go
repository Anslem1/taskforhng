package handler

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	. "github.com/tbxark/g4vercel"
)

type ipInfo struct {
	City    string `json:"city"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello go from hng !!!!",
		})
	})

	server.GET("/api/hello", func(context *Context) {
		name := context.Query("visitor_name")
		ip := getClientIP(r)
		name = strings.Trim(name, `"`)
		context.JSON(200, H{
			"message":   fmt.Sprintf("Hello, %v! The temperature is 11 degrees Celsius in New York!", name),
			"client_ip": ip,
			"location":  "New York",
		})

	})

	server.Handle(w, r)
}
