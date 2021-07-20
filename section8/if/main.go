package main

import "fmt"

// if
// 条件分岐

func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("I don't know")
	}

	// 簡易文つきif文 
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	// 簡易文つきif文の注意点
	x := 0
	if x :=2; true {
		
		fmt.Println(x)
		// if文内では
		// 2
		// 外側ではなく内部の変数が優先される
	}
	fmt.Println(x)
	// 0
	// 外側の変数がないとエラーが出る
}
