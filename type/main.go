package main

import (
	"fmt"
)

func main() {
	/*
	* 論理値型
	 */
	b := true
	b = false
	fmt.Println(b)

	/*
	* 数値型
	 */
	// Goは環境によって32ビット整数を表したり64ビット整数を表すような「実装依存」の
	// 難しさを軽減する為に、明確に定義された多数の数値型が用意されている
	// 符号付き整数型は int8, int16, int32, int64
	var i8 int8
	i8 = 10
	fmt.Println(i8)
	fmt.Printf("型=%T\n", i8)
}
