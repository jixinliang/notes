package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	filename := "/dev/random"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open file failed", err)
		return
	}
	defer file.Close()

	var seed int64
	binary.Read(file, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
