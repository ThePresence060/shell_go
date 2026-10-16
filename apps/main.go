package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){
	reader := bufio.NewReader(os.Stdin)
	for {fmt.Print("$")
		input, err := reader.ReadString('\n')

	if err!= nil {
		break
	}

	input = strings.TrimSpace(input)
	parts := strings.Fields(input)

	if len(parts) == 0{
		continue
	}

	command := parts[0]
	args := parts[1:]

	if handled := handleBuiltin(command,args);handled{
		continue
	}
	executeExternal(command,args)
	

	}

}

func checkNilMethd(err error){
	if err!= nil{
		panic(err)
	}
}