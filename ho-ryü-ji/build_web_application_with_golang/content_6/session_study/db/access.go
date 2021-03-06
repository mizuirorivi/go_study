package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB interface {
	Get(sessionid string) (User, error)
	Set(data User) error
	Delete(sessionid string) error
}

type User struct {
	SessionId string
	Name      string
	Password  string
}

func init() {
	db, err := sql.Open("sqlite3", "./user.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	sqlStmt := `
	create table if not exists user(sessionid text not null primary key, name text, password text);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}

func Get(sessionid string) (*User, error) {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		return nil, err
	}
	stmt, err := db.Prepare("select * from user where sessionid = ? limit 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user User
	err = stmt.QueryRow(sessionid).Scan(&user.SessionId, &user.Name, &user.Password)
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
	stmt, err := tx.Prepare("insert into user(sessionid, name, password) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.SessionId, user.Name, user.Password)
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
