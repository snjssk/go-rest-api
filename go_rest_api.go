// package go_rest_api
package main

import (
	"log"
	"net/http"
)

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

	// エラーが返ってくるため、log.Fatal()を利用する
	log.Fatal(http.ListenAndServe(":18080", nil))

}
