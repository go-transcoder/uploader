package route

import (
	"github.com/go-transcoder/uploader/videos/controllers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/upload", controllers.Post).Methods("POST")

	return r
}
