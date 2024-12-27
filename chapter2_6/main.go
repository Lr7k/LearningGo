package main

import "fmt"
import "os"
import "log"

func main() {
	var fileName string = "./chapter2_6" + "/main.go"
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	var data []byte = make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(count, string(data))
}
