package main

import "fmt"

func executeExternal(command string,args[]string){
	fmt.Printf("$%s: command not found\n",command)
}