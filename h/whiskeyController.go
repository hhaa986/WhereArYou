package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wayBack/model"
)

// GetWhiskeysHandler godoc
// @Summary get whiskey
// @Schemes
// @Description give me a whiskey
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} model.CreateWRequest{Name :  }
// @Router /whiskey/{name} [get]
func GetWhiskeysHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.CreateWRequest{
		WName:        "",
		AlcoholLevel: 0,
		Origin:       "",
		CID:          0,
	})
}

//func WhiskeyController(router *mux.Router) error {
//	err := serv.Service.InitService()
//	if err != nil {
//		return err
//	}
//	// GET 전체 위스키 반환
//	router.HandleFunc("/whiskey", func(w http.ResponseWriter, r *http.Request) {
//		whiskeys := serv.Service.GetAllWhiskey()
//		Response(w, whiskeys, http.StatusOK, nil)
//	}).Methods("GET")
//
//	//// GET name 위스키 반환
//	//router.HandleFunc("/whiskey/{name}", func(w http.ResponseWriter, r *http.Request) {
//	//	name :=
//	//	whiskeys := serv.Service.GetWhiskey(name)
//	//	Response(w, whiskeys, http.StatusOK, nil)
//	//}).Methods("GET")
//
//	// CREATE
//	router.HandleFunc("/whiskeyss", func(w http.ResponseWriter, r *http.Request) {
//		var cwr model.CreateWRequest
//		json.NewDecoder(r.Body).Decode(&cwr)
//		whiskeys := serv.Service.CreateWhiskey(cwr)
//		Response(w, whiskeys, http.StatusOK, nil)
//	}).Methods("POST")
//	return nil
//}
