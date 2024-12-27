package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// ログファイル作成
func LoggingSetting(logFile string) {
	// 書き込むログを作成
	file, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// ログを画面出力
	multiLogFile := io.MultiWriter(os.Stdout, file)
	// ログ出力時の設定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// 出力先の設定
	log.SetOutput(multiLogFile)
}

func main() {
	log.Println("logging start")
	log.Printf("%T %v", "sample", "sample")
	LoggingSetting("test.log")
	_, err := os.Open("aaaaaaa")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("log.Fatalfでエラーとフォーマット。でもそれ以降は実行されない")
	log.Fatalf("error %v", "error")
	log.Fatalln("error")
	fmt.Println("この行は実行されない")
}
