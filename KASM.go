package main

import (
    "fmt"
	"bufio"
	"log"
    "os"
)


func main() {
    fmt.Println("hello world")
	allArgs := os.Args
	if(len(allArgs)>1){
		fileName := allArgs[1]
		lines, err := readLines(fileName)
		if err != nil {
			log.Fatalf("readLines: %s", err)
		}
		for i, line := range lines {
			fmt.Println(i, line)
		}
	}
	
}


// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
	curLine:= scanner.Text()
	if(curLine!=""){
		lines = append(lines, curLine)
	}
    
  }
  return lines, scanner.Err()
}


