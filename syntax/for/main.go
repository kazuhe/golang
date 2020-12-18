package main

import (
	"fmt"
)

func main() {
	/*
	* for文
	 */
	// ループを記述する為の構文はforのみ

	// ↓の様に何も条件を指定しないと無限ループとなる
	// for {
	// 	fmt.Println("無限ループ")
	// }

	// 条件式を与える一般的なfor文
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Goは使用する予約後を少なく抑えている代償に
	// 各制御構文の記述方法は複数のバリエーションが存在する

}
