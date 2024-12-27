package main

import "fmt"

func add1(x int) {
	x += 1
	fmt.Println("関数add1内")
	fmt.Println(x)
}

func add2(x *int) {
	*x += 1
}

func main() {
	fmt.Println("ポイント操作")
	var num int = 10
	fmt.Println(num)
	fmt.Println("変数numのポインタを取得")
	fmt.Println(&num)
	var p *int = &num
	fmt.Println(p)
	fmt.Println("ポインタの中身を表示")
	fmt.Println(*p)
	fmt.Println("&をつけた変数に*をつけても同じ")
	fmt.Println(*&num)
	fmt.Println("関数内で値を増やせるかテスト")
	var x int = 1
	add1(x)
	fmt.Println("main関数の中")
	fmt.Println(x)
	fmt.Println("main関数の中では変わってない")
	fmt.Println("関数に渡してmain関数でも変更されているようにする")
	add2(&x)
	fmt.Println(x)
}
