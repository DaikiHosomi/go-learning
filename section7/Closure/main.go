package main

import "fmt"
//関数
//クロージャー

func Later() func(string) string{
	var store string
	// クロージャーが実行される限り、storeは破棄されない
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// 実行のタイミングで、storeの変数が初期化されない
// 環境とセットになった関数うとしてのクロージャーとしての性質は非常に大切




func main() {
	f := Later() 
	fmt.Println(f("Hello"))
	fmt.Println(f("my"))
	fmt.Println(f("name"))
	fmt.Println(f("is"))
	fmt.Println(f("Golang"))

	//
	// Hello
	// my
	// name
	// is

}