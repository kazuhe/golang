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

	/*
	* break
	 */
	// breakでループを中断させる
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}

	/*
	* 条件式のみ与えたfor
	 */
	// whileと同等
	foo := 30
	for foo < 35 {
		fmt.Println(foo)
		foo++
	}

	/*
	* continue
	 */
	// ループ内でブロックの残処理をスキップして次のループ処理へ継続させる
	for i := 50; i < 55; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		i++
	}

	/*
	* 範囲節によるfor
	 */
	// 予約語rangeと任意の色を組み合わせて定義する
	// 配列以外にもスライスやマップや文字列などにも適用できる
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	for i, s := range fruits {
		// iは文字列配列のインデックス
		// sはインデックスに対応した文字列の値
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	// Goは使用する予約後を少なく抑えている代償に
	// 各制御構文の記述方法は複数のバリエーションが存在する

}
