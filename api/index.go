package handler

import (
	"fmt"
	. "github.com/tbxark/g4vercel"
	"net"
	"net/http"
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
		cityInfo, err := GetCityInfo(ip)

		if err != nil || name == "" {
			context.JSON(400, H{
				"message":   "name not found or error getting city info",
				"client_ip": ip,
				"location":  cityInfo.City,
			})
		} else {
			context.JSON(200, H{
				"message":   fmt.Sprintf("Hello, %v! The temperature is 11 degrees Celsius in %v.", name, cityInfo),
				"client_ip": ip,
				"location":  cityInfo.City,
			})
		}
	})

	server.Handle(w, r)
}
