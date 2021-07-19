package main

//関数
//関数を引数にとる関数

import "fmt"

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I am a function")
	})

}
