package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func func_call() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimRight(input,"\n")
		if input == "exit 0" {
			os.Exit(0)
		}
			
	}
}

func main(){
	func_call()
}