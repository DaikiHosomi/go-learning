package main

import "fmt"
// 定数

// 頭文字を大文字にすることで他のパッケージから呼び出せる
// 頭文字を小文字にするとこのパッケージからのみ呼び出せる
const Pi = 3.14

// 複数の定数を定義する場合
const (
	URL = "https//xxx.co.jp"
	SiteName = "test"
)

// 値の省略
// B　と　C　は何も定義していないが　Aの値が入る
// E　と　F　は何も定義していないが　Dの値が入る
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 定数な最大値がない
// inth型の最大値に＋１すると、、
// var Big int = 9223372036854775807 + 1
// 変数だとエラーが出るoverflows
const Big  = 9223372036854775807 + 1


const (
	c0 = iota
	// iotaとは連続する整数の連番を生成
	c1 
	c2
)

func main() {
	fmt.Println(Pi)


	// Pi = 3
	// fmt.Println(Pi)
	// 上書きできないためコメントアウト
	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(Big - 1)
	fmt.Println(c0, c1, c2)
}