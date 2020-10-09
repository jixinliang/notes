package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
)

func genBytes(n int64) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func genPasswd(length int64) (string, error) {
	bytes, err := genBytes(length)
	return base64.URLEncoding.EncodeToString(bytes), err
}

func main() {
	var length int64 = 8
	args := os.Args

	switch len(args) {
	case 2:
		length, _ = strconv.ParseInt(args[1], 10, 64)
		if length <= 0 {
			length = 8
		}
	default:
		fmt.Println("using default values.")
	}

	passwd, err := genPasswd(length)
	if err != nil {
		fmt.Println("Generate password failed", err)
		return
	}
	fmt.Println("Password:", passwd[0:length])

}
