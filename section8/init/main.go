package main

import "fmt"

// init
// goのパッケージの初期化を行う

// パッケージ変更に伴った際にint内に処理を追加することで初期化処理を行える
func init() {
	fmt.Println("init")
}

// init が複数ある場合は上から順に読み込まれる
// しかし、２つ作成するメリットはあまりなく、一つのint関数にまとめてそこに処理を追加していくことが多い
func init() {
	fmt.Println("init2")
}
// main関数よりも先に読み込まれる
// main関数の中で読み込む必要がない


func main() {
	fmt.Println("Main")

}
// init
// init2
// Main