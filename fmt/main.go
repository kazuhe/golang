package main

import (
	"fmt"
)

func main() {
	/*
	* fmt.Println
	 */
	// 文字列の最後に改行を付加して標準出力（テキストを単純に1行にで出力するのに使用）
	fmt.Println("Hello Golang!")
	fmt.Println("Hello", "Golang!") // 複数の引数の場合は各々スペースで区切られて出力

	/*
	* fmt.Printf
	 */
	// 書式指定子を含んだ文字列とそれに続く可変長の引数を与えて文字列を標準出力
	fmt.Printf("数値=%d\n", 5) // %dが書式と埋め込む位置を表す

	// 数値用の書式
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)

	// %vは様々な型のデータを埋め込む
	fmt.Printf("数値=%v 文字列=%v 配列=%v\n", 5, "Golang", [...]int{1, 2, 3})

	// %#vはGoのリテラル表現でデータを埋め込む
	fmt.Printf("数値=%#v 文字列=%#v 配列=%#v\n", 5, "Golang", [...]int{1, 2, 3})

	// %Tはデータの型情報を埋め込む
	fmt.Printf("数値=%T 文字列=%T 配列=%T\n", 5, "Golang", [...]int{1, 2, 3})

	/*
	* print, println
	 */
	// 引数として与えられた文字列を「標準エラー出力」へ出力
	print("Hello Golang!")   // 末尾改行無し
	println("Hello Golang!") // 末尾改行有り
}
