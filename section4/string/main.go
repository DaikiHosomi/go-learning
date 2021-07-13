package main

import "fmt"
//型
//文字列型

func main() {
	var s string = "hello golang"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	var s1 string = "300"
	fmt.Println(s1)
	fmt.Printf("%T\n", s1)

	fmt.Println(`test
	test
			test`)

	fmt.Println("\"")
	fmt.Println(`"`)

	//byte型
	fmt.Println(s[0])
	
	fmt.Println(string(s[0]))

}