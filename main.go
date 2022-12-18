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
	dbWhiskey.CreateWhiskeyDb(model.Whiskey{Name: "wsk", AlcoholLevel: 50})
}

func mainURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello... this is %s", r.URL.Path[1:])
}
