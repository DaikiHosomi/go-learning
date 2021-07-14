package main

import "fmt"

// 論理演算子

func main() {
	fmt.Println(true && false == true)
	// false
	fmt.Println(true && true == true)
	fmt.Println(true && false == false)
	//true

	fmt.Println(true || false == true)
	//true
	fmt.Println(false || false == true)
	//flase

	fmt.Println(!true)
	//false
	fmt.Println(!false)
	//true

	// ！が論理値を反転させる



}