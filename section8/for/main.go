package main

import "fmt"

// for文
// 繰り返し処理

func main() {
	/* i := 0
	for {
		i ++ 
		if i == 3 {
			// for文を強制終了させる意味のbreak
			break
			//for ループの中を抜ける
		}
		fmt.Println("Loop")
	} */

/* 	point := 0
	// point が10になったらforを止める
	for point < 10 {
		fmt.Println(point)
		point ++
	} */
	// 0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9

	/* or i := 0; i < 10; i ++ {
		if i == 3 {
			// ３になったらスキップしてくれ！！
			continue
		}
		if i == 6 {
			// 6になったらfor文を終わらしてくれ！
			break
		}
		fmt.Println(i)
	} */
	// 0
	// 1
	// 2
	// 4
	// 5

	//配列で行う場合は？
	/* arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i ++ {
		fmt.Println(arr[i])
	} */
	// 1
	// 2
	// 3

	// 範囲式のfor
	/* arr := [3]int{1, 2, 3}

	for i, v :=range arr {
		fmt.Println(i, v)
		// i は　index番号
		// v は　要素の値　
	} */
	// 0 1
	// 1 2
	// 2 3
	
	// index番号を表示したくない場合
	/* for _, v :=range arr {
		fmt.Println(v)
	} */
	// 1
	// 2
	// 3

	// 要素の値を表示したくない場合
	/* for i :=range arr {
		fmt.Println(i)
	} */
	// 0
	// 1
	// 2

	// sliceの場合

	/* sl := []string{"Python", "PHP", "GO"}
	for i, v :=range sl {
		fmt.Println(i, v)
	}  */
	// 0 Python
	// 1 PHP
	// 2 GO

	// mapの場合
	// key valueの形式で配列を作成する
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
		// k は　key
		// v は　value
	}
	// keyだけ表示、valueだけ表示も対応可能










}