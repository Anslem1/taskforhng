package handler

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/ipinfo/go/ipinfo"
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
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

func getCityInfo(ip string) (string, error) {
	var info ipInfo

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return "", fmt.Errorf("error loading .env file")
	}

	apiToken := os.Getenv("IPINFO_API_TOKEN")
	if apiToken == "" {
		return "", fmt.Errorf("IPINFO_API_TOKEN is not set")
	}

	client := ipinfo.NewClient(http.DefaultClient)
	netIp := net.ParseIP(ip)
	ipInfoResp, err := client.GetCity(netIp)
	if err != nil {
		return "", err
	}

	info.City = ipInfoResp 
	return ipInfoResp , nil
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

		fmt.Println(cityInfo, "cityy", "errorrrr:", err)
		if err != nil || name == "" {
			context.JSON(400, H{
				"message":   "name not found or error getting city info",
				"client_ip": ip,
				"location":  cityInfo,
			})
		} else {
			context.JSON(200, H{
				"message":   fmt.Sprintf("Hello, %v! The temperature is 11 degrees Celsius in %v.", name, cityInfo),
				"client_ip": ip,
				"location":  cityInfo,
			})
		}
	})

	server.Handle(w, r)
}