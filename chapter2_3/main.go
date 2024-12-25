package main

import "fmt"
import "time"

func main() {
	fmt.Println("switchの基礎")
	fmt.Println("switch分で条件分岐をする")
	var sex string
	sex = "male"
	fmt.Println("switchの後にcaseで条件判定")
	switch sex {
	case "male":
		fmt.Println("male")
	case "female":
		fmt.Println("female")
	default:
		fmt.Println("defaultは必須ではない")
		fmt.Println("default")
	}
	fmt.Println("switchで条件分岐をしたが、caseで条件分岐を行う")
	var t time.Time = time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	case 17 < t.Hour():
		fmt.Println("evening")
	}
}
