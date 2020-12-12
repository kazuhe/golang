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

// エラー処理確認の為の仮func
func doError() (string, error) {
	msg := "エラー発生"
	return "実行", fmt.Errorf("%s", msg)
}

/*
* 戻り値を表す変数
 */
// 戻り値xはreturnで返していないが、0の値を持った変数xが返される
func doSomething() (x, y int) {
	y = 5
	return // → 0 5
}

// 上記の関数は↓の関数を短縮している
func doSomething2() (int, int) {
	var x, y int
	y = 5
	return x, y // → 0 5
}

// つまり、戻り値に変数を持たせた場合は関数内で変数を定義しなくても
// returnで初期値を返すことができる

/*
* 関数を返す関数
 */
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}
