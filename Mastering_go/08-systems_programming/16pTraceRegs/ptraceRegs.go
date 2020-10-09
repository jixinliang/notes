package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	args := os.Args

	var pTraceRegs syscall.PtraceRegs

	cmd := exec.Command(args[1], args[2:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Ptrace: true}

	err := cmd.Start()
	if err != nil {
		fmt.Println("Start:",err)
		return
	}

	err = cmd.Wait()
	fmt.Printf("State: %v\n", err)

	pid := cmd.Process.Pid

	err = syscall.PtraceGetRegs(pid, &pTraceRegs)
	if err != nil {
		fmt.Println("PtraceGetRegs:", err)
		return
	}

	fmt.Printf("Registers: %#v\n", pTraceRegs)
	fmt.Printf("R15: %d, Gs: %d\n", pTraceRegs.R15, pTraceRegs.Gs)

	time.Sleep(2 * time.Second)
}
