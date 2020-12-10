package main

import "fmt"

/*
* 基本
 */
// func [関数名]([引数定義]) [戻り値型] {
//   [関数本体]
// }

func plus(x, y int) int {
	return x + y
}

/*
* 引数定義
 */
// func plus(x int, y int) int
// の様にそれぞれ型を定義することも可能だが冗長なので省略する

/*
* 戻り値
 */
// 戻り値が無い場合は型の省略をするだけ
func hello() {
	fmt.Println("Hello")
	return
}

/*
* 複数の戻り値
 */
// 複数の戻り値を定義するには「()」で囲む
func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}
