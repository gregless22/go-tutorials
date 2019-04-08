package main

import (
	"encoding/json"
	"log"
	"net/http"
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

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	//mock up the articles
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	pl("Endpoint Hit: Homepage")

	json.NewEncoder(w).Encode(articles)
}

//
// ──────────────────────────────────────────────────── I ──────────
//   :::::: R O U T E S : :  :   :    :     :        :          :
// ──────────────────────────────────────────────────────────────
//
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllArticles)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
