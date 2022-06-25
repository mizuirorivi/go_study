package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	Get(where, value string) (*User, error)
	Set(data User) error
	Update(sessionid, data User) error
	// Delete(sessionid string) error
}

type User struct {
	SessionId string
	Name      string
	Password  string
	Life      int
}

func init() {
	db, err := sql.Open("sqlite3", "./user.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
	create table if not exists user(sessionid text not null primary key, name text, password text, life integer);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func Get(where, value string) (*User, error) {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		return nil, err
	}

	prestr := "select * from user where " + where + "= ? limit 1"
	stmt, err := db.Prepare(prestr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(value).Scan(&user.SessionId, &user.Name, &user.Password, &user.Life)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (user *User) Set() error {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into user(sessionid, name, password, life) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.SessionId, user.Name, user.Password, user.Life)
	if err != nil {
		tx.Rollback()
		log.Fatalf("insert error: %v", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Update(sessinid string, updated *User) error {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("update user set(sessionid, name, password, life) where sessionid = ? values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(sessinid, updated.SessionId, updated.Name, updated.Password, updated.Life)
	if err != nil {
		log.Fatal(err)
	}
	a, err := res.RowsAffected()
	if err == nil {
		log.Fatal(err)
	}

	fmt.Println(a)

	return nil
}

// func Delete(sessionid string) error {
// 	db, err := sql.Open("sqlite3", "./user.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stmt, err := db.Prepare("delete from user where sessionid = ?")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	res, err := stmt.Exec(sessionid)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	a, err := res.RowsAffected()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(a)

// 	return nil
// }
