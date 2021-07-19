package main

// 関数
// ジェネレーターの実装

// ジェネレーターとは、何かのルールに従って連続した値を返し続ける仕組みのこと
// Goではクロージャーを応用することでジェネレーターのような振る舞いを実装できる

import "fmt"

func integers() func() int {
	i := 0
	return func() int {
		i ++
		return i
	}
}

func main() {
	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	// 1
	// 2
	// 3
	// 4

	otherints := integers()
	fmt.Println(otherints())
	fmt.Println(otherints())
	fmt.Println(otherints())
	// 1
	// 2
	// 3	

}