package main

import (
	"database/sql"
	"fmt"
	// _ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Insert(table string) {
	var name, sex, species string
	var birthday, age int64

	if table == "Animals" {
		db, err := sql.Open("sqlite3", "./mydb.db")
		checkErr(err)
		defer db.Close()

		prepareString := "INSERT INTO" + table + "(Name, Sex, Birthday, Age, Species) values(?, ?, ?, ?, ?)"
		stmt, err := db.Prepare(prepareString)
		checkErr(err)

		fmt.Println("Enter Values")
		fmt.Scan(&name, &sex, &birthday, &age, &species)
		res, err := stmt.Exec(name, sex, birthday, age, species)
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)
	} else {
		fmt.Println("It does not exist.")
	}
}

func Update(table string) {
	db, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)
	defer db.Close()

	var row, where, strValue1, strValue2 string
	var intValue1, intValue2 int64
	var res *sql.Result

	fmt.Println("Enter a rowId and a whereId to set.")
	fmt.Scan(row)
	fmt.Scan(where)
	prepareString := "update" + table + "set" + row + "=? where" + where + "=?"
	stmt, err := db.Prepare(prepareString)
	checkErr(err)

	fmt.Println("Enter a rowValue and a whereValue to set.")

	if row == "Name" || row == "Sex" || row == "Species" {
		if where == "Name" || where == "Sex" || where == "Species" {
			fmt.Scan(&strValue1, &strValue2)
			res, err = stmt.Exec(strValue1, strValue2)
			checkErr(err)
		} else if where == "Birthday" || where == "Age" {
			fmt.Scan(&strValue1, &intValue2)
			res, err = stmt.Exec(strValue1, intValue2)
			checkErr(err)
		}
	} else if row == "Birthday" || row == "Age" {
		if where == "Name" || where == "Sex" || where == "Species" {
			fmt.Scan(&intValue1, &strValue2)
			res, err = stmt.Exec(intValue1, strValue2)
			checkErr(err)
		} else if where == "Birthday" || where == "Age" {
			fmt.Scan(&intValue1, &intValue2)
			res, err = stmt.Exec(intValue1, intValue2)
			checkErr(err)
		}
	} //else {
	// 	return errors.New("Error")
	// }

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func Select(table string) {
	db, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)
	defer db.Close()

	var row string
	fmt.Println("Enter * or a row name you want to see.")
	fmt.Scan(&row)
	queryString := "SELECT" + row + "FROM" + table + ""
	rows, err := db.Query(queryString)
	checkErr(err)

	for rows.Next() {
		var Name, Sex, Species string
		var Birthday, Age int64

		err = rows.Scan(&Name, &Sex, &Birthday, &Age, &Species)
		checkErr(err)

		fmt.Println(Name)
		fmt.Println(Sex)
		fmt.Println(Birthday)
		fmt.Println(Age)
		fmt.Println(Species)
	}
}

func Delete(table string) {
	db, err := sql.Open("sqlite3", "./mydb.db")
	checkErr(err)
	defer db.Close()

	var where string
	fmt.Println("Enter a row name where you want to delete.")
	prepareString := "delete from" + table + "where" + where + "=?"
	stmt, err := db.Prepare(prepareString)
	checkErr(err)

	res, err := stmt.Exec()
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func main() {
	var order, table string

	for {
		fmt.Println("Enter a table name and an order. (INSERT, UPDATE, SELECT or DELETE)")
		fmt.Scan(&order, &table)

		if table == "Animals" {
			if order == "INSERT" {
				Insert(table)
			} else if order == "UPDATE" {
				Update(table)
			} else if order == "SELECT" {
				Select(table)
			} else if order == "DELETE" {
				Delete(table)
			} else {
				fmt.Println("Such an order doesn't exist.")
			}
		} else {
			fmt.Println("Such a table doesn't exist.")
		}
	}
}
