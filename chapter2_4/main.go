package main

import "fmt"
import "os"

func sample() {
	fmt.Println("sample start")
	defer fmt.Println("defer sample")
	fmt.Println("sample end")
}

func main() {
	fmt.Println("deferについて")
	defer fmt.Println("defer start")
	fmt.Println("sample")
	sample()
	fmt.Println("ファイルをcloseするときにdeferを使う")
	fmt.Println("var使って型明示しながら一行に変数複数宣言できないやん")
	file, _ := os.Open("./main.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
