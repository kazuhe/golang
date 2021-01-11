package main

import (
	"fmt"
)

func main() {
	/*
	* 構造体
	 */
	// 「構造体（struct）」とは複数の任意の型の値を1つにまとめたもの
	// Goでは複雑なデータ構造を1つの「型」として取り扱うことができる構造体の定義は重要な位置を占める

	/*
	* type
	 */
	// typeは既に定義されている型をもとに新しい型を定義するための機能
	// 基本型のint型にMyIntのようなエイリアスを定義する
	type MyInt int
	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1) // 5
	fmt.Println(n2) // 7

	// メリットはmap[string][2]float64のような複雑な型にエイリアスを定義し見通しがよくなる
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)
	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)
	fmt.Println(pair) // [1 2]
	fmt.Println(strs) // [Apple Banana Cherry]
	fmt.Println(amap) // map[Tokyo:[35.689488 139.691706]]
	fmt.Println(ich)  // 0xc00008c060

	// 関数型にエイリアスを定義することもできる
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n) // 30
}
