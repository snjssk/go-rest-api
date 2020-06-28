package models

import (
	"log"
)

type Items struct {
	Id	  int `json:id`
	Name  string `json:name`
	Price int `json:price`
}

//func NewItems(id int, name string, price int) *Items {
//	fmt.Println("hello")
//	return &Items{
//		id,
//		name,
//		price,
//	}
//}

func GetItem(id string) []Items {
	log.Println("id:", id)

	cmd := `SELECT * FROM items`
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()

	var items []Items
	for rows.Next() {
		var i Items
		err := rows.Scan(&i.Id, &i.Name, &i.Price)
		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, i)
	}
	return items
}