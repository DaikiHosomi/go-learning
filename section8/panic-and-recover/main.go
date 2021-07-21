package main 

import "fmt"


// panic recover
// 強制的にプログラムを終了させる強力なもの
// 一般的には、エラーハンドリングで対応すべし
func main() {
	defer func() {
		// recoverはpanicによって発生したrun time errorを復活させるもの（defferとセットして使う）
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("run time error")
	fmt.Println("start")
}
// run time error


