package main

import (
	"fmt"
)

/*
* defer
 */
// 関数終了時に実行される式を登録できる

func runDefer() {
	defer fmt.Println("defer!") // ↓の"done"の後に表示される
	fmt.Println("done")
}

func main() {
	runDefer()

	// リソースの開放処理で活躍
	// ファイルをオープンできたら関数の終了時に確実にそのファイルがクローズされるようにdefer文を利用

	// file, err := osOpen("/path/to/file")
	// if err != nill {
	// 	// fileのオープンに失敗したらreturn
	// 	return
	// }
	// // ファイルのクローズ処理を登録
	// defer file.Close()
}
