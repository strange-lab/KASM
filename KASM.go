package main

import (
    "fmt"
    //"io/ioutil"
	"os"
)


func main() {
    fmt.Println("hello world")
	allArgs := os.Args
	if(len(allArgs)>1){
	  fileName := allArgs[1]
	  fmt.Println(fileName)
	}
	
}
