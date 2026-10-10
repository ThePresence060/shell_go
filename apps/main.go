package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Print("$: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Print(input)
}