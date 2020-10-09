package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var Data []Entry
var tFileName string

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Printf("Host: %s Path: %s", r.Host, r.URL.Path)
//	t := template.Must(template.ParseGlob(tFileName))
//	t.ExecuteTemplate(w, tFileName, Data)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFileName))
	myT.ExecuteTemplate(w, tFileName, Data)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("go run %s example.db example.html\n",filepath.Base(args[0]))
		return
	}

	database := args[1]
	tFileName = args[2]

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println("open database", err)
		return
	}

	fmt.Println("create table.")
	_, err = db.Exec("create table numbers(id int primary key ,number int, double int, square int)")
	if err != nil {
		fmt.Println("Create talbe failed",err)
		return
	}

	fmt.Println("Emptying database table.")
	_, err = db.Exec("delete from numbers")
	if err != nil {
		fmt.Println("Empty database falild:", err)
		return
	}

	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("insert into numbers(number, double, square) values(?,?,?)")
	for i := 2; i < 20; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	rows, err := db.Query("select * from numbers")
	if err != nil {
		fmt.Println("select database", err)
		return
	}

	var n, d, s int
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		tmp := Entry{n, d, s}
		Data = append(Data, tmp)
	}

	fmt.Println("Http server runing......")
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
