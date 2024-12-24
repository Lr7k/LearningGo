package main

import "fmt"

func main() {
	fmt.Println("if文")
	var num int
	num = 4
	if num%2 == 0 {
		fmt.Println("偶数")
	} else {
		fmt.Println("奇数")
	}
	fmt.Println("if文 else if")

	num = 3
	if num%2 == 0 {
		fmt.Println("偶数")
	} else if num%3 == 0 {
		fmt.Println("3で割り切れます")
	} else {
		fmt.Println("その他")
	}

	fmt.Println("複数条件")
	var x int
	var y int
	x, y = 10, 10
	if x == 10 && y == 10 {
		fmt.Println("&&です")
	}
	y = 5
	if x == 10 || y == 10 {
		fmt.Println("||です")
	}
	fmt.Println("if文の条件式で変数宣言")
	num = 2
	var res string = by2(num)
	if res == "OK" {
		fmt.Println("2で割れる")
	} else {
		fmt.Println("2で割れない")
	}
	fmt.Println("みにくいからつかわねー")
	if res2 := by2(num); res2 == "OK" {
		fmt.Println("2で割れる")
	} else {
		fmt.Println("2で割れない")
	}
}

func by2(num int) string {
	if num%2 == 0 {
		return "OK"
	} else {
		return "NG"
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
