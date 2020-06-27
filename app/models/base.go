package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// DbConnectionはグローバル
var DbConnection *sql.DB

func init() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	// create
	cmd := `CREATE TABLE IF NOT EXISTS items(
				id INT,
				name STRING,
				price INT
			)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// insert
	//cmd = `INSERT INTO items (id, name, price) VALUES (?, ?, ?)`
	//_, err = DbConnection.Exec(cmd, 3, "item3", 300)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// update
	//cmd = `UPDATE items SET price = ? WHERE id = ?`
	//_, err = DbConnection.Exec(cmd, 400, 3)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// multi select
	cmd = `SELECT * FROM items`
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
	for _, i := range items {
		fmt.Println(i.Id, i.Name, i.Price)
	}

}
