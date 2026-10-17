package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var builtin = map[string] bool{
	"echo" : true,
	"exit" : true,
	"type" : true,
}

func handleBuiltin(command string, args[]string) bool{
	switch command{
	case "exit":
		handleExit(args)
		return true

	case "echo":
		handleEcho(args)
		return true
	case "type":
		if len(args) == 0{
			return true
		}else{
			target := args[0]

			if _,ok := builtin[target] ; ok{
				fmt.Printf("%s is a shell builtin\n",target)
				return true
			}else{
				fmt.Printf("%s: command not found\n",target)
				return true
			}
		}

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