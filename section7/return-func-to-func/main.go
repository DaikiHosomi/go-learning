package main

import "fmt"

//関数
// 関数を返す関数

func ReturnFunc() func() {
	// 無名関数を作成
	return func() {
		fmt.Println("I am a function")
	}
}

func main() {
	// 中の無名関数を定義
	f := ReturnFunc()
	f()
}