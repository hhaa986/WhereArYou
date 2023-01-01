package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	_ "wayBack/docs"
	h "wayBack/handler"
	s "wayBack/service"
)

// @title Way API
// @version 0.0.1
// @description Whiskeyyyyyydddyy
// @host localhost:8000
// @BasePath /

func main() {
	r := mux.NewRouter()
	err := s.Service.InitService()
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/whiskey", h.GetWhiskeys).Methods(http.MethodGet)
	r.HandleFunc("/whiskey", h.CreateWhiskey).Methods(http.MethodPost)
	r.HandleFunc("/whiskey", h.UpdateWhiskey).Methods(http.MethodPut)
	r.HandleFunc("/whiskey/{id:[0-9]+}", h.DeleteWhiskey).Methods(http.MethodDelete)

	//category
	r.HandleFunc("/whiskeyCategory", h.GetWhiskeysCategory).Methods(http.MethodGet)
	r.HandleFunc("/whiskeyCategory", h.CreateWhiskeyCategory).Methods(http.MethodPost)
	r.HandleFunc("/whiskeyCategory", h.UpdateWhiskeyCategory).Methods(http.MethodPut)
	r.HandleFunc("/whiskeyCategory/{id:[0-9]+}", h.DeleteWhiskeyCategory).Methods(http.MethodDelete)

	corsHandler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8000", corsHandler)
	if err != nil {
		fmt.Println(err)
	}
}
