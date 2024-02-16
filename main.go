package main

import (
	"fmt"
	"github.com/go-transcoder/uploader/route"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Printf("Uploader App running on port :%s\n", port)

	// get the router
	router := route.GetRouter()
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		fmt.Printf("error starting server: %v", err)
	}
}
