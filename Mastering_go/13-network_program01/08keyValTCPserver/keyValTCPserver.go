package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type info struct {
	ID      string
	Name    string
	SurName string
}

const msg = "Welcome to the key value store!\n"

var Data = make(map[string]info)
var DataFile = "/tmp/dataFile.gob"

func init() {
	file, err := os.Create(DataFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(Data)
	if err != nil {
		fmt.Println(err)
	}
}

func connHandler(conn net.Conn) {
	conn.Write([]byte(msg))

	for {
		readString, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error at Read conn->", err)
			return
		}

		cmd := strings.TrimSpace(readString)
		tokens := strings.Fields(cmd)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "exit":
			err := save()
			if err != nil {
				fmt.Println(err)
			}
			conn.Close()
			return
		case "getData":
			getData(conn)
		case "del":
			if !del(tokens[1]) {
				tips := "Delete operation failed!\n"
				conn.Write([]byte(tips))
			} else {
				tips := "Delete operation successful!\n"
				conn.Write([]byte(tips))
			}
		case "add":
			newData := info{tokens[2], tokens[3], tokens[4]}
			if !add(tokens[1], newData) {
				tips := "Add operation failed!\n"
				conn.Write([]byte(tips))
			} else {
				tips := "add operation successful!\n"
				conn.Write([]byte(tips))
			}
		case "get":
			data := get(tokens[1])
			// there is a data
			if data != nil {
				newData := fmt.Sprintf("%v\n", *data)
				conn.Write([]byte(newData))
			} else {
				tips := "Did not find the key!\n"
				conn.Write([]byte(tips))
			}
		case "edit":
			newData := info{tokens[2], tokens[3], tokens[4]}
			if !edit(tokens[1], newData) {
				tips := "Updata operation failed!\n"
				conn.Write([]byte(tips))
			} else {
				tips := "Updata operation successful!\n"
				conn.Write([]byte(tips))
			}
		default:
			tips := "Unknow command - please try again!\n"
			conn.Write([]byte(tips))
		}
	}
}

func save() error {
	fmt.Println("Saving:", DataFile)
	err := os.Remove(DataFile)
	if err != nil {
		return err
	}

	file, err := os.Create(DataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(Data)
	if err != nil {
		return err
	}

	return nil
}

func load() error {
	fmt.Println("Loading:", DataFile)
	file, err := os.Open(DataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&Data)
	if err != nil {
		return err
	}
	return nil
}

func get(key string) *info {
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
	if get(key) == nil {
		Data[key] = value
		return true
	}
	return false
}

func edit(key string, value info) bool {
	if get(key) != nil {
		Data[key] = value
		return true
	}
	return false
}

func del(key string) bool {
	if get(key) != nil {
		delete(Data, key)
		return true
	}
	return false
}

func getData(conn net.Conn) {
	for key, val := range Data {
		data := fmt.Sprintf("key: %s value: %v\n", key, val)
		conn.Write([]byte(data))
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: go run %s <Port>\n", filepath.Base(args[0]))
		return
	}

	port := ":" + args[1]
	listener, err := net.Listen("tcp4", port)
	if err != nil {
		fmt.Println("Error at listen")
		return
	}
	defer listener.Close()

	err = load()
	if err != nil {
		fmt.Println("Error at load()->",err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error at Accept()->", err)
			return
		}
		go connHandler(conn)
	}
}
