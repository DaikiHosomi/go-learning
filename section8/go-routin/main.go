package main

import (
	"fmt"
	"time"
)

// go goroutin
// 並行処理の機能
// 並行処理される新しい処理の流れをrun timeに追加する
// go文を使うことで、並行処理の実装ができる

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// go を追加するだけで、sub() とmain()を同時に並行して処理させることができる
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}