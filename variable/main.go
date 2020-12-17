package main

import (
	"fmt"
)

/*
* パッケージ変数
 */
// Goは関数の中で定義された変数は「ローカル変数」
// 関数の外で定義された変数は「パッケージ変数」
// packageVariable := 100 ←暗黙的な変数定義は関数外では不可能
var packageVariable = 100

func main() {
	// Goは変数の参照チェックが厳密で、参照されていない変数があるとコンパイルエラーとなる

	// パッケージ変数に値を代入
	packageVariable = packageVariable + 1
	fmt.Println(packageVariable)

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
	i = 2
	// i := 2 ←コンパイルエラー！
	// 「:=」を使った再代入は許されない「:=」は「変数の代入」ではなく「変数の定義」
	fmt.Println(i)

	b := true
	fmt.Println(b)

	f := 3.14
	fmt.Println(f)

	s := "abc"
	fmt.Println(s)

	// 関数の戻り値を変数に代入
	callOne := one()
	fmt.Println(callOne)

	/*
	* varと暗黙的な定義
	 */
	// 変則的だが下記の様に定義する事も可能
	var a = 1
	// だが、:=に比べて冗長なので「a := 1」と記述する方が適切
	fmt.Println(a)

	// しかし、複数の変数を暗黙的に定義する場合は変数定義のブロックを可視化できる為好ましい
	var (
		num = 1
		str = "string"
	)
	// ↑↓ どちらが見やすいか？
	num2 := 1
	str2 := "string"

	fmt.Println(num, str, num2, str2)

	/*
	* まとめ
	 */
	//  基本的には↓の型推論がきく暗黙的な方法で定義する
	muni := "muni"

	// が、あらかじめ変数を定義しておく場合や複数の変数を定義する場合はvarを使って記述する
	var puni string

	var (
		nnn = 2
		sss = "複数定義"
	)

	fmt.Println(muni, puni, nnn, sss)

	/*
	* 定数
	 */
	// ./const.go
	constSample()

	// ファイルが分割されていても定数や変数は相互に参照可能
	// ./one.go の変数「two」を参照
	fmt.Println(two)
}
