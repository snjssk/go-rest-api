package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/snjssk/go-rest-api/app/models"
	"github.com/snjssk/go-rest-api/config"
)

// Content-Type
// https://stackoverflow.com/questions/51456253/how-to-set-http-responsewriter-content-type-header-globally-for-all-api-endpoint
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func WebServer()  {
	// router
	router := mux.NewRouter()
	router.Use(middleware)

	// /item
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItems).Methods("GET")

	// serve
	log.Println("Listen Server ....")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), router))
}

func getItems(w http.ResponseWriter, r *http.Request){
	// パラメータ
	vars := mux.Vars(r)
	id := vars["id"]
	log.Println(id)

	// データを取得
	item := models.GetItem(id)
	// データを返却
	// json.NewEncoder(w).Encode(item)
	fmt.Fprintf(w, "%s", item)
}