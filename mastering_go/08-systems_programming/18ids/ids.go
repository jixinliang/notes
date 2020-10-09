package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	fmt.Println("User ID:", os.Getuid())

	var u *user.User

	u, _ = user.Current()
	fmt.Println("Name:", u.Name)
	fmt.Println("Uid", u.Uid)
	fmt.Println("Gid:", u.Gid)
	fmt.Println("username:", u.Username)

	fmt.Print("Group IDs:")
	groupIds, _ := u.GroupIds()
	for _, val := range groupIds {
		fmt.Print(val, " ")
	}
	fmt.Println()
}
