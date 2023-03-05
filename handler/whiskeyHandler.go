package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"way_BE/model"
	s "way_BE/service"
)

// GetWhiskeys
// @Summary This is simple Summary.
// @Description This is detail Description.
// @Accept  json
// @Produce  json
// @Router /whiskey [get]
// @Success 200 {object} string
func GetWhiskeys(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	cid := r.URL.Query().Get("cid")
	uId, _ := strconv.Atoi(id)
	uCid, _ := strconv.Atoi(cid)
	fmt.Printf("name: %s, id: %d, cid: %d\n", name, uId, uCid)
	var rWhiskeys []model.Whiskey
	if len(name) > 0 {
		rWhiskeys = s.Service.GetWhiskeyByName(name)
	} else if len(cid) > 0 {
		rWhiskeys = s.Service.GetWhiskeyByCategory(uint(uCid))
	} else if len(id) > 0 {
		wh := s.Service.GetWhiskeyById(uint(uId))
		fmt.Print(wh)
	} else {
		rWhiskeys = s.Service.GetAllWhiskey()
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rWhiskeys)
}

func CreateWhiskey(w http.ResponseWriter, r *http.Request) {
	var wsk model.WhiskeyRequest
	json.NewDecoder(r.Body).Decode(&wsk)
	fmt.Println(wsk)
	rWhiskeys := s.Service.CreateWhiskey(model.Whiskey{
		Name:   wsk.WName,
		ABV:    wsk.AlcoholLevel,
		Origin: wsk.Origin,
		CID:    wsk.CID,
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rWhiskeys)
}

func UpdateWhiskey(w http.ResponseWriter, r *http.Request) {
	var wsk model.WhiskeyUpdateRequest
	json.NewDecoder(r.Body).Decode(&wsk)
	fmt.Println(wsk)
	rWhiskeys := s.Service.UpdateWhiskey(model.Whiskey{
		Model: gorm.Model{
			ID: wsk.Id,
		},
		Name:   wsk.WName,
		ABV:    wsk.AlcoholLevel,
		Origin: wsk.Origin,
		CID:    wsk.CID,
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rWhiskeys)
}

func DeleteWhiskey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	err = s.Service.DeleteWhiskey(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}

// GetWhiskeysCategory category
func GetWhiskeysCategory(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	uId, _ := strconv.Atoi(id)
	fmt.Printf("name: %s, id: %d\n", name, uId)
	var wCategory []model.WCategory
	if len(name) > 0 {
		wCategory = s.Service.GetWCategoryByName(name)
	} else if len(id) > 0 {
		wh := s.Service.GetWCategory(uint(uId))
		fmt.Print(wh)
	} else {
		wCategory = s.Service.GetAllWCategory()
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(wCategory)
}

func CreateWhiskeyCategory(w http.ResponseWriter, r *http.Request) {
	var wcr model.WCategoryRequest
	json.NewDecoder(r.Body).Decode(&wcr)
	fmt.Println(wcr)
	wCategory := s.Service.CreateWCategory(model.WCategory{
		Name: wcr.WName,
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wCategory)
}

func UpdateWhiskeyCategory(w http.ResponseWriter, r *http.Request) {
	var wsur model.WhiskeyUpdateRequest
	json.NewDecoder(r.Body).Decode(&wsur)
	fmt.Println(wsur)
	wCategory := s.Service.UpdateWCategory(model.WCategory{
		Model: gorm.Model{
			ID: wsur.Id,
		},
		Name: wsur.WName,
	})
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(wCategory)
}

func DeleteWhiskeyCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	err = s.Service.DeleteWCategory(uint(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
