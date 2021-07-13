package main

import "fmt"
// 配列型

func main() {
	// 要素の数が３つのint型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)
	// 明示的定義で、要素の数が３つのstring型かつ値を持たせてみると、、	
	var arr2 [3]string = [3]string{"A", "B"}
	// ３つ目は空文字として表示される
	fmt.Println(arr2)
	// 明示的定義で値を持たせて行うと、’、
	arr3 := [3]int{1,2,3}
	fmt.Println(arr3)
	// 要素数の省略-> 要素の数を自動で決定
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)
	// [2]string

	// 配列の操作
	// 値の取り出し(変数の値のindex番号で指定)
	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	// 配列の値を[]内で計算可能
	fmt.Println(arr2[2-1])

	// 値の更新
	arr2[2] = "C"
	fmt.Println(arr2)

	// var arr5 [4]int
	// arr5 = arr1
	// fmt.Println(arr5)
	// cannot use arr1 (type [3]int) as type [4]int in assignment
	// 中の要素が同じでも要素数が異なる場合は全く別の型として認識されてエラーになる

	//変数の中の要素数を調べる場合
	fmt.Println(len(arr1))
}