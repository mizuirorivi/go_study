package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func testMySQLConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:t3st@tcp(127.0.0.1:3306)/mydb")
	defer db.Close()

	if err != nil {
		t.Fatal("[!]", err)
	}
	t.Log("[+]", "MySQL connection established")
	stmtOut, err := db.Prepare("SELECT * from user")
	if err != nil {
		t.Fatal(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtOut.Close()
}
