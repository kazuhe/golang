package main

import (
	"fmt"
)

/*
* 定数
 */
// パッケージに定義された定数
const X = 1

func constSample() {
	// 関数内でも定義可能
	const (
		Y = 2
		Z = 3
	)
	fmt.Printf("%v %v %v\n", X, Y, Z)

	/*
	* 値の省略
	 */
	// 省略する場合初めに定義された定数の値がそのまま以降の定数に割り当てられる
	// 途中で値を定義した場合は暗黙的に値が切り替わる
	// ※全ての値の省略は不可
	const (
		X1 = 1
		Y1 // 1
		Z1 // 1
		S1 = "text"
		S2 // text
	)
	fmt.Printf("%v %v %v %v %v\n", X1, Y1, Z1, S1, S2)
}
