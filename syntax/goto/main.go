package main

import (
	"fmt"
)

func main() {
	/*
	* goto
	 */
	// 一般的なプログラムの作法としてはできるだけ使用を避けるべきと言われている
	// goto文と「ラベル」:の形式で任意の位置にジャンプすることができる
	fmt.Println("A")
	goto L
	fmt.Println("B") // 処理されない
L:
	fmt.Println("C")

	// forのブロック内や変数定義を飛び越えることコンパイルエラーになる
	// 基本的にgotoは使用しない想定で問題なし
}
