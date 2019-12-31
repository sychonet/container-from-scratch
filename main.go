package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command("proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//Namespacing the hostname
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//Cloneflags are parameters that will be passed on clone syscall. Clone is what actually creates a new process.
		//CLONE_NEWUTS is the namespace
		//UTS : Unix Timesharing System
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
