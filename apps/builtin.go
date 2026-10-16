package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)


func handleBuiltin(command string, args[]string) bool{
	switch command{
	case "exit":
		handleExit(args)
		return true

	case "echo":
		handleEcho(args)
		return true
	
	default:
		return false
	}
}


func handleExit(args []string){
	if len(args) == 0{
		os.Exit(0)
	}

	code, err := strconv.Atoi(args[0])
	if err != nil{
		fmt.Printf("exit: %s: numeric argument required\n",args[0])
		os.Exit(2)
	}
	os.Exit(code)
}


func handleEcho(args []string){
	fmt.Println(strings.Join(args," "))
}