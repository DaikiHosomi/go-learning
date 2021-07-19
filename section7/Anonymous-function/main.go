package main

//無名関数

// 本質的には従来の関数と変わらない
// しかし、簡易的に試したい場合などの用途によって使えると言える
import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i :=f(1, 2)
	fmt.Println(i)

	i2 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i2)
}
