package handler

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"
	"os"

	"github.com/joho/godotenv"
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
        ip = r.RemoteAddr
    }
    return ip
}

func getCityInfo(ip string) (ipInfo, error) {
    var info ipInfo

    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        return info, fmt.Errorf("error loading .env file")
    }

    apiToken := os.Getenv("IPINFO_API_TOKEN")
    if apiToken == "" {
        return info, fmt.Errorf("IPINFO_API_TOKEN is not set")
    }

    url := fmt.Sprintf("https://ipinfo.io/%s?token=%s", ip, apiToken)
    resp, err := http.Get(url)
    if err != nil {
        return info, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return info, err
    }
    if err := json.Unmarshal(body, &info); err != nil {
        return info, err
    }
    return info, nil
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
        cityInfo, err := getCityInfo(ip)

        if err != nil || name == "" {
            context.JSON(400, H{
                "message":   "name not found or error getting city info",
                "client_ip": ip,
                "location":  cityInfo.City,
            })
        } else {
            context.JSON(200, H{
                "message":   fmt.Sprintf("Hello, %v! The temperature is 11 degrees Celsius in %v.", name, cityInfo.City),
                "client_ip": ip,
                "location":  cityInfo.City,
            })
        }
    })

    server.Handle(w, r)
}
