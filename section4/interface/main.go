package main

import "fmt"

// interface & nill

func main() {
	// あらゆる型と互換性のある特殊な型
	// ex) 数値型や文字列型もinterfaceと互換性がある
	var x interface{}
	fmt.Println(x)
	// 初期値は　<nil>　→　値を何も持っていない（全ての型と互換性あり）
	// 数値型に変換した場合
	x = 1
	fmt.Println(x)
	x =  3.14
	fmt.Println(x)
	// 文字列型に変換した場合
	x = "A"
	fmt.Println(x)
	// 配列型に変換した場合
	x = [3]int{1,2,3}
	fmt.Println(x)

	// 値の更新→OK
	//  データ計算や演算→出来ない
	// x = 2
	// fmt.Println(x + 3)
	// invalid operation: x + 3 (mismatched types interface {} and int)

// 	interfaceは全ての型を汎用的に表す手段

// 	・全ての型と互換性がある
// 	・初期値は　nill
}