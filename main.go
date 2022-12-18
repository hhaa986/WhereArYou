package main

import (
	"fmt"
	"net/http"
	"wayBack/config"
	"wayBack/db"
	"wayBack/model"
)

func main() {
	//toml Setting
	conf := config.SetToml()

	//db
	dbCon := db.GetDbConnection(conf)
	dbWhiskey := db.NewWhiskeyDb(dbCon)
	dbWhiskey.CreateWhiskeyCDb(model.WCategory{
		Name: "싱글몰트?",
		//Whiskeys: []model.Whiskey{},
	})
}

func mainURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello... this is %s", r.URL.Path[1:])
}
