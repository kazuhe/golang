package main

import (
	"fmt"
)

func main() {
	// 変数は大きく分けて「値型」「参照型」「ポインタ型」の3種類

	/*
	* 明示的な定義
	 */
	// varを使用する場合は変数名と型を明示的に指定する必要がある
	var n int

	// 同じ型であれば複数の変数まとめて定義可能
	var x, y, z int

	// ()で囲むことで異なる型の変数をまとめて定義可能
	var (
		nyan, nyon int
		name       string
	)

	// 値を代入 型が正しい限り再代入の制限は無い
	n = 5

	// 複数の値を一括代入
	x, y, z = 1, 2, 3

	fmt.Println(n)
	fmt.Println(x, y, z)
	fmt.Println(nyan, nyon, name)

	/*
	* 暗黙的な定義
	 */
	//  暗黙的な変数定義の場合は型推論によって型の定義が必要無い
	i := 1

	fmt.Println(i)
}
