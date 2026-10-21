package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func executeExternal(command string,args[]string){
	fmt.Printf("$%s: command not found\n",command)
}

func findInPath(command string) string{
	pathEnv := os.Getenv("PATH")
	if pathEnv == ""{
		return ""
	}

	dirs := filepath.SplitList(pathEnv)

	for _, dir:= range dirs{
		fullPath := filepath.Join(dir,command)

		info,err := os.Stat(fullPath)

		if err!= nil{
			continue
		}

		if !info.IsDir() && info.Mode().Perm()&0111 != 0{
			return fullPath
		}
	}
	return ""

}
