package main

import "fmt"

// switch文
// 型スイッチ

// まず型アサーション→動的に変数の型をチェックするgoの重要な機能の１つ
// 全ての型と互換性のあるinterface型を使用すると様々な型の引数得られる
func anything (a interface{}) {
	// fmt.Println(a)

	switch v :=a.(type) {
	case string: 
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
	// aaa!?
	// 10001

}

func main () {
	anything("aaa")
	anything(1)
	// aaa
	// 1

	var x interface{} = 3
	// 型アサーションの実施
	// xをint型で復元
	// i := x.(int)
	// fmt.Println(i + 2)
	i , isInt := x.(int)
	fmt.Println(i , isInt)
	// 3 true


	// f := x.(float64)
	// xはint型のため、float型での復元はできない
	// run time　エラーがでる
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)
	// 0 false
	// run time　エラーが起きない→強制終了させないで実行可能
	// 変換に失敗するとfalseで帰ってくる

	// 型アサーションと条件分岐を使ってinterface型で引数をわたして、様々な課題に対応した処理を分岐して行う
	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		// isIntがtrueであれば
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		// isStringがtrueであれば
		fmt.Println(s, "s is String")
	} else {
		fmt.Println("I don't know")
	}
	// 3 x is Int


	// 型によるスイッチ文→型アサーションと条件分岐を組み合わせる
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}
	// int
	//if文よりもコンパクトに条件分岐を行える。

	// もし値も表示したい場合は
	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool") 
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
	// 3 int





}