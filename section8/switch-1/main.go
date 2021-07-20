package main

import "fmt"

// switch
// 式switch

func main() {
	/* n := 5
	switch n {
	// caseは行くでも任意で作成可能
	// switchで与えられた式の型とcaseの型が違うとエラーが発生する
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	// ケース分のどれにも当てはまらない場合はdefaltを実行する
	// defaultは１つまで
	default:
		fmt.Println("I don't know")
	} */
   
	/* switch n := 2; n {
	// switch 文内のみ利用可能な変数としてnを定義する方法
	case 1,2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}
	 */

	/* n := 6
	switch {
		// swith文内のcaseでnを使う場合は、switch文で与える式nを省略可能
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	} */
	// 3 < n < 7

	/* switch n := 2; n {
	case 1,2:
		fmt.Println("1 or 2")
	case 3,4:
		fmt.Println("3 or 4")
	case n > 3 && n < 6:
		fmt.Println("3 < n < 6")
	default:
		fmt.Println("I don't know")
	} */
	// 列挙型の式と演算子型の式を混在させた場合は以下のようなエラーがでる
	// ./main.go:49:2: cannot use n > 3 && n < 6 (type bool) as type int
}