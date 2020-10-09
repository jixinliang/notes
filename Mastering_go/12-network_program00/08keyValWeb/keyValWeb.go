package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type info struct {
	ID      string
	Name    string
	SurName string
}

var Data = make(map[string]info)
var DataFile = "/tmp/dataFile.gob"

func init() {
	file, err := os.Create(DataFile)
	if err != nil {
		fmt.Println("Error at Create:", err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(Data)
	if err != nil {
		fmt.Println("Error at Encode:", err)
	}
}

func saveFile() error {
	fmt.Println("Saving:", DataFile)

	err := os.Remove(DataFile)
	if err != nil {
		fmt.Println("Error at Remove:", err)
		return err
	}

	file, err := os.Create(DataFile)
	if err != nil {
		fmt.Println("Error at Create:", err)
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(Data)
	if err != nil {
		fmt.Println("Error at Encode:", err)
		return err
	}

	return nil
}

func loadFile() error {
	fmt.Println("Loading:", DataFile)

	file, err := os.Open(DataFile)
	if err != nil {
		fmt.Println("Error at Open:", err)
		return err
	}

	decoder := gob.NewDecoder(file)
	decoder.Decode(&Data)

	return nil
}

func lookUp(key string) *info {
	_, ok := Data[key]
	if ok {
		data := Data[key]
		return &data
	}
	return nil
}

func add(key string, value info) bool {
	if key == "" {
		return false
	}

	if lookUp(key) == nil {
		Data[key] = value
		return true
	}
	return false
}

func del(key string) bool {
	if lookUp(key) != nil {
		delete(Data, key)
		return true
	}
	return false
}

func edit(key string, value info) bool {
	Data[key] = value
	return true
}

func cat() {
	for key, val := range Data {
		fmt.Printf("key: %s value: %v\n", key, val)
	}
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving:", r.Host, "For", r.URL.Path)
	templ := template.Must(template.ParseGlob("index.gohtml"))
	templ.ExecuteTemplate(w, "index.gohtml", nil)
}

func listAllHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the all contents of the KV store!")
	fmt.Println("Listing the contents of the KV store!")
	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">List all elements!</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\">Edit an element!</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\">Insert new element!</a>")

	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")
	for key, val := range Data {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", key, val)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}

func editEleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of the KV store!")
	// must() make sure that the template file provided contains no errors
	templ := template.Must(template.ParseFiles("update.gohtml"))

	if r.Method != http.MethodPost {
		templ.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	value := info{
		r.FormValue("id"),
		r.FormValue("name"),
		r.FormValue("surname"),
	}

	if !edit(key, value) {
		fmt.Println("Update operation failed!")
	} else {
		err := saveFile()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	templ.Execute(w, struct{ Success bool }{true})
}

func insertEleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the KV store!")
	templ := template.Must(template.ParseFiles("insert.gohtml"))

	if r.Method != http.MethodPost {
		templ.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	value := info{
		r.FormValue("id"),
		r.FormValue("name"),
		r.FormValue("surname"),
	}

	if !add(key, value) {
		fmt.Println("Add peration failed!")
	} else {
		err := saveFile()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	templ.Execute(w, struct{ Success bool }{true})
}


func main() {

	err := loadFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	port := ":8080"
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = ":" + args[1]
		fmt.Println("Using port:", port)
	}

	fmt.Println("Http runing...")
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/list", listAllHandler)
	http.HandleFunc("/edit", editEleHandler)
	http.HandleFunc("/insert", insertEleHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
