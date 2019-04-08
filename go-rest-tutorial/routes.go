package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//
// ────────────────────────────────────────────────────────── I ──────────
//   :::::: F U N C T I O N S : :  :   :    :     :        :          :
// ────────────────────────────────────────────────────────────────────
//

func homePage(w http.ResponseWriter, r *http.Request) {
	fpf(w, "welcome to the hompage")
	pl("Endpoint Hit: Homepage")
}

func get(w http.ResponseWriter, r *http.Request) {
	//mock up the articles
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	pl("Endpoint Hit: Homepage")

	json.NewEncoder(w).Encode(articles)
}

func show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fpf(w, "Key: "+key)
}

//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: R O U T E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", get)
	myRouter.HandleFunc("/articles/{id}", show)
	log.Fatal(http.ListenAndServe(":8090", myRouter))
}
