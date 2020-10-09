package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	pid, _, _ := syscall.Syscall(39, 0, 0, 0)
	fmt.Println("My pid:", pid)
	uid, _, _ := syscall.Syscall(24, 0, 0, 0)
	fmt.Println("My uid:", uid)

	msg := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	syscall.Write(1, msg)

	cmd := "/bin/ls"
	env := os.Environ()
	syscall.Exec(cmd, []string{"ls", "-a", "-x", "-l"}, env)
}
