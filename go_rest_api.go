// package go_rest_api
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/snjssk/go-rest-api/app/models"
	"log"
	"net/http"

	"github.com/snjssk/go-rest-api/config"
)

func getItems(w http.ResponseWriter, r *http.Request){
	//title := r.URL.Path
	// fmt.Fprintf(w, "<h1>%s</h1>", title)
	//fmt.Println(models.GetItem)
	//json.NewEncoder(w).Encode(models.GetItem)

	//params := mux.Vars(r)
	//log.Println(params) // map[id:1]

	vars := mux.Vars(r)
	id := vars["id"]
	log.Println(id)

	item := models.GetItem(id)
	json.NewEncoder(w).Encode(item)
}

func main()  {
	// データの取得
	// ioutilはバイト文字で返ってくるため、stringでキャストする
	//resp, _ := http.Get("https://www.google.co.jp/")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	// URL が正しいものかの判定
	//base, _ := url.Parse("http://example.com/aaa")
	//reference, _ := url.Parse("/test?a=b")
	//endpoint := base.ResolveReference(reference)
	//fmt.Println(endpoint)
	fmt.Printf("%T %v\n", config.Config.Port, config.Config.Port)

	// http.ListenAndServeの前に登録してあげる
	// /item
	// http.HandleFunc("/item", itemHandler)

	// エラーが返ってくるため、log.Fatal()を利用する
	// log.Fatal(http.ListenAndServe(":18080", nil))

	// router
	router := mux.NewRouter()
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItems).Methods("GET")

	log.Println("Listen Server ....")
	log.Fatal(http.ListenAndServe(":18080", router))
}
