package main

import (
	"fmt"
	"os"
)
//defer

func TestDefer() {
	// deferを使うことで、関数内の最後（終了時）に実行するように処理できる
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	// 複数deferが登録されている場合、最後に定義したdeferから処理される
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// 3
	// 2
	// 1
}

func main () {
	TestDefer()

	// main関数内の最後に実行
	// defer func() {
	// 	fmt.Println("1")
	// 	fmt.Println("2")
	// 	fmt.Println("3")
	// }()

	RunDefer()
    // file作成
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// fileをcloseする
	defer file.Close()

	file.Write([]byte("Hello"))
}