package main

import (
	"fmt"
)

func interfaceTest() {
	/*
	* interface{}とnil
	 */
	// interface{}型は「{}」も含めて1つの型
	// これはGoにおけるあらゆる型と互換性のある特殊な型
	// int型の場合は初期値は0になるがinterface{}型はnilとなる
	// Goのnilは「具体定期な値を持っていない」状態を表す
	var x interface{}
	fmt.Printf("%#v\n", x) // → <nil>
	fmt.Printf("%#T\n", x) // → <nil>

	// ありとあらゆる型の代入が可能
	x = 1
	fmt.Printf("%#v\n", x) // → 1
	fmt.Printf("%#T\n", x) // → int

	x = "文字"
	fmt.Printf("%#v\n", x) // → "文字"
	fmt.Printf("%#T\n", x) // → string

	x = [...]uint8{1, 2, 3}
	fmt.Printf("%#v\n", x) // → [3]uint8{0x1, 0x2, 0x3}
	fmt.Printf("%#T\n", x) // → [3]uint8

	// 一旦変数に格納されてしまうと整数の同士の加算などデータ型特有の演算もできなくなる
	// interface{}型はあくまでも"全ての型を汎用的に表す手段"で何らかの演算対象として利用は不可能
	var y, z interface{}
	y, z = 1, 2
	// n := y + z → エラー 演算はできない
	fmt.Printf("%#v %#v\n", y, z)

	/*
	* まとめ
	 */
	// ・すべての型を代入することができる
	// ・初期値としてnilという特殊な型を持つ
}
