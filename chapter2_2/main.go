package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	fmt.Println("for文の基礎")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("continueの利用")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("cotinueで飛ばします")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("breakの利用")
	for i := 0; i < 10; i++ {
		if i > 5 {
			fmt.Println("breakで抜けます")
			break
		}
		fmt.Println(i)
	}
	fmt.Println("rangeの利用")
	var languages []string = []string{"python", "java", "go"}
	for i := 0; i < len(languages); i++ {
		fmt.Println(languages[i])
	}
	for i, v := range languages {
		fmt.Println(i, v)
	}
	fmt.Println("rangeのindexを使わない場合")
	for _, v := range languages {
		fmt.Println(v)
	}
	fmt.Println("mapとrangeの組み合わせ")
	var m map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
