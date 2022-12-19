package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	serv "wayBack/service"
)

type CommonResponse struct {
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
	Error  interface{} `json:"error"`
}

func Response(w http.ResponseWriter, data interface{}, status int, err error) {
	var res CommonResponse

	if status == http.StatusOK {
		res.Data = data
		res.Status = status
	} else {
		res.Status = status
		res.Error = err.Error()
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}

func WhiskeyController(router *mux.Router) error {
	err := serv.Service.InitService()
	if err != nil {
		return err
	}
	// GET 전체 위스키 반환
	router.HandleFunc("/whiskey", func(w http.ResponseWriter, r *http.Request) {
		whiskeys := serv.Service.GetAllWhiskey()
		Response(w, whiskeys, http.StatusOK, nil)
	}).Methods("GET")
	return nil
}
