package main

// 変数
import "fmt"

// i5 := 500
var i5 int = 500

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}



func main() {
	//明示的な定義
	// var 変数名　型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "hello go"
	fmt.Println(s)

	var t,f bool = true, false
	fmt.Println(t,f)

	//  異なる型でまとめる場合
	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	//変数の型をだけを指定する方法
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	// 変数に値を代入
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 変数上書き
	i = 150
	fmt.Println(i)

	// 暗黙的な定義
	// 変数名 :=値
	i4 :=400
	fmt.Println(i4)
	// 上書き
	i4 = 450
	fmt.Println(i4)

	// i4 :=500
	// fmt.Println(i4)

	// i4 = "Hello"
	// fmt.Println(i4)

	fmt.Println(i5)

	outer()

	// fmt.Println(s4)

	var s5 string = "Not Use"
	fmt.Println(s5)
}