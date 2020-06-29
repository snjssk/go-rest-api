package models

import (
	"encoding/json"
	"log"
)

type Items struct {
	Id	  int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func GetItem(id string) []string {
	log.Println("id:", id)
	log.Printf("%T", id)

	// idで絞り込み
	var cmd string
	if id == "" {
		cmd = `SELECT * FROM items`
	} else {
		cmd = `SELECT * FROM items WHERE id = ?`
	}

	// クエリーを発行
	rows, _ := DbConnection.Query(cmd, id)
	defer rows.Close()

	// 結果をjson形式でリテラルで指定した名前にして返却する
	var result []string
	for rows.Next() {
		var i Items
		err := rows.Scan(&i.Id, &i.Name, &i.Price)
		if err != nil {
			log.Fatalln(err)
		}
		v, _ := json.Marshal(i)
		result = append(result, string(v))
	}
	return result
}