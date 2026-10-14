package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func func_call() {
	reader := bufio.NewReader(os.Stdin)

	fun := make(map[string]string)
	fun["echo"] = "echo"

	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		
		
		if len(args) == 0{
			continue
			
		}
		command := args[0]

		switch command{
		case "exit":
			os.Exit(0)
		case "echo":
			for _,val := range args[1:]{
					strings.Join(args[1:]," ")
					fmt.Print(val," ")
			}
			fmt.Print("\n")
		default:
			fmt.Printf("%s: command not found!!\n",command)
		}
			

	}
}

func main(){
	func_call()
}