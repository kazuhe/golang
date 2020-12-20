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

}
