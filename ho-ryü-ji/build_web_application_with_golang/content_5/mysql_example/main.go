package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:t3st@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS mydb")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE IF EXISTS animals")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE animals (
		name VARCHAR(255),
		sex VARCHAR(255),
		age INT,
		species VARCHAR(255),
		datetime DATETIME
	)`)
	if err != nil {
		panic(err)
	}

	dbt, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
		return
	}
	// now := time.Now().In(dbt)
	// fmt.Println(now)
	_, err = db.Exec("INSERT INTO animals (name, sex, age, species, datetime) VALUES (?, ?, ?, ?, ?)", "chiro", "female", 14, "dog", time.Now().In(dbt))
	if err != nil {
		panic(err)
	}

	rows, verr := db.Query("SELECT * FROM animals")
	if verr != nil {
		panic(verr)
	}

	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			panic(err)
		}

		values := make([]interface{}, len(columns))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}

		err = rows.Scan(scanArgs...)
		if err == nil {
			for i, value := range values {
				switch value.(type) {
				default:
					fmt.Printf("%s :: %s :: %s\n", columns[i], reflect.TypeOf(value), value)
				}
			}
		} else {
			panic(err)
		}
	}
}
