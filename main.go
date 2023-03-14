package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "way_BE/docs"
	h "way_BE/handler"
	s "way_BE/service"
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

	r.HandleFunc("/api/whiskey", h.GetWhiskeys).Methods(http.MethodGet)
	r.HandleFunc("/api/whiskey", h.CreateWhiskey).Methods(http.MethodPost)
	r.HandleFunc("/api/whiskey", h.UpdateWhiskey).Methods(http.MethodPut)
	r.HandleFunc("/api/whiskey/{id:[0-9]+}", h.DeleteWhiskey).Methods(http.MethodDelete)

	//category
	r.HandleFunc("/api/whiskeyCategory", h.GetWhiskeysCategory).Methods(http.MethodGet)
	r.HandleFunc("/api/whiskeyCategory", h.CreateWhiskeyCategory).Methods(http.MethodPost)
	r.HandleFunc("/api/whiskeyCategory", h.UpdateWhiskeyCategory).Methods(http.MethodPut)
	r.HandleFunc("/api/whiskeyCategory/{id:[0-9]+}", h.DeleteWhiskeyCategory).Methods(http.MethodDelete)

	r.PathPrefix("/api/swagger").Handler(httpSwagger.WrapHandler)
	corsHandler := cors.Default().Handler(r)
	err = http.ListenAndServe(":8000", corsHandler)
	if err != nil {
		fmt.Println(err)
	}
}
