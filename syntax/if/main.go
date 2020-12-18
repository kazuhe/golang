package main

import (
	"fmt"
)

// 例文用 エラーを返す関数
func doError() (string, error) {
	msg := "エラー発生"
	return "実行", fmt.Errorf("%s", msg)
}

func main() {
	/*
	* if文
	 */
	// ifの条件式は必ず論理値を返す必要がある
	// 1文の場合に{}を省略することは不可
	x := 1
	if x == 1 {
		fmt.Println("xは1である")
	}

	/*
	* 簡易文付きif
	 */
	// ifの中でのみ使用可能な局所的な変数を作る事ができる
	// 式や代入分を持たない簡易式を条件式の前に置き;で区切る
	if x, y := 5, 10; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}
	// この場合、簡易式で定義された変数xはif分の外にある変数xとは名称は同一でも別の変数となる
	// if文の中で外の変数xには一切アクセスできない
	// また、外からも一切アクセスできない

	// 例えばエラー処理の際に簡易文付きifを実務的に使える
	if _, err := doError(); err != nil {
		// 関数doErrorがエラーありと返した場合の処理
		fmt.Println(err)
	}
	// errはif文の中でのみ有効になるので外側のスコープに影響を与えない
}
