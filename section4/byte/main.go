package main

import "fmt"


//型
//byte(uint8型)
func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	// 文字列に変換
	fmt.Println(string(byteA))

	c := []byte("HI")
	// byteのスライス配列に変換
	fmt.Println(c)
	// 文字列に変換
	fmt.Println(string(c))
	
}