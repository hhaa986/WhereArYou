package main

import (
	"github.com/gorilla/mux"
	"net/http"
	cont "wayBack/controller"
)

func main() {
	//toml Setting
	//conf := config.SetToml()

	r := mux.NewRouter()
	err := cont.WhiskeyController(r)

	if err != nil {
		panic("서버 실행 실패")
	}
	http.ListenAndServe(":8000", r)
}
