package main

import (
	"fmt"
)

func main() {
	/*
	* switch
	 */
	//  「式によるswitch文」と「型によるswitch文」にの2つの処理がある

	/*
	 * 式によるswitch文
	 */
	// switchに与えた式の値とcase節に列挙された値が演算子==によって真になる場合にcase節の処理が実行される
	// どのcase節にも分岐しなければdefault節が実行される
	// case節に複数の値を与えることができる

	// 「switch n := 3; n {」 ← 簡易文を置くことも可能
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("unknown")
	}

	// Goのswitchはbreakがなく一致したcase節でswitch文は終了するが、fallthroughを置くことで次の節へ処理が継続される
	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)

	// 式を伴うcase ほぼif文の置き換えの様な文法でbool型を返す任意の式を書ける
	// switchに式を与えない場合「switch true {}」と内部的になっている
	// caseに定数と式を混在して与えると型の不一致によるエラーが発生
	num := 5
	switch {
	case num > 0 && num < 3:
		fmt.Println("0 < num < 3")
	case num > 3 && num < 6:
		fmt.Println("3 < num < 6")
	}

	/*
	 * 型によるswitch文
	 */
	// 型アサーションと分岐を組み合わせた処理を手軽に書ける
	x := interface{}(7)
	switch x.(type) {
	case bool:
		fmt.Println("bool")
	case int, uint:
		fmt.Println("integer or unsigned interger")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("don't know")
	}

	// 値も必要であれば変数に代入する
	// 一致しないcase節は実行されないので、
	// 型アサーションに失敗してランタイムエラーになる事がない
	switch v := x.(type) {
	case bool:
		fmt.Println("bool:", v)
	// case int, uint: ← 1つの型に定まらないのでエラー
	case int:
		fmt.Println(v * v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("don't know:", v)
	}
	// ↑の様に変数に値を格納する場合はcase節で複数の型を列挙するのは控えた方が良い
}
