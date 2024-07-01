package handler

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	query := r.URL.Query().Get("visitor_name")
	ipAddr := "127.0.0.1"
	query = strings.Trim(query, `"`)

	greeting := fmt.Sprintf("Hello, %v! The temperature is 11 degrees Celsius in Mark.", query)

	server.GET(`/api/hello?visitor_name="Mark"`, func(context *Context) {
		context.JSON(200, H{
			"client_ip": ipAddr,
			"greeting":  greeting,
			"location":  "New York",
		})
	})
	server.Handle(w, r)
}
