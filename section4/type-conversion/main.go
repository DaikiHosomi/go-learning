package main

import (
	"fmt"
	"strconv"
)
// 型変換

func main (){
	var i int = 1
	// int型の1をfloat型に型変換
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i)
	fmt.Printf("fl64 = %T\n", fl64)
	// // float型をint型に戻す型変換
	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2)
	fmt.Println(i2)


	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3)
	fmt.Println(i3)
	// int型に変更したことで、小数点以下切り捨てされる
	// 文字列型を数値型に変換
	// Atoiは文字列型を数値型に変換する関数
	// 関数で帰ってくる値は２つある
	// １つ目はiという変数
	// ２つ目の＿　２つ目を破棄する
	// ２つ目を_にすることで使わなくてもエラー出ない
	var s string = "100"
	fmt.Printf("s = %T\n", s)
	i, err := strconv.Atoi(s)
	fmt.Println(s)
	fmt.Println(err)
	fmt.Printf("i = %T\n", i)

	// // 数理型から文字列型に変換
	// var i2 int = 200
	// // Atoiは数値型を文字列型に変換する関数
	// s2 := strconv.Itoa(i2)
	// fmt.Println(s2)
	// fmt.Printf("s2 = %T\n", s2)

	// 文字列からbyte配列に変換
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	// // byte型から文字列型に変換
	h2 := string(b)
	fmt.Println(h2)
}