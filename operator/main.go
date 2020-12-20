package main

import (
	"fmt"
)

func main() {
	/*
	* 演算子
	 */
	// 算術演算時
	// + - * / % & | ^ &^ << >>

	// 比較演算子
	// == != < <= > >=

	// 論理演算子
	// || &&

	// 単項演算子
	// + - ! ^ * & <-

	/*
	* 算術演算子
	 */
	// 算術演算子は「+ - * /」の四則演算子のみ整数・浮動小数点・複素数を対象にしており、それ以外は整数が対象
	// しかし、特別に「+」だけは文字列の結合に使用できる
	s := "Taro" + " " + "Yamada"
	fmt.Printf("%v\n", s) // → Taro Yamada

	// 商はゼロの方向に切り捨て/切り上げられる
	num1 := 5 / 3
	num2 := -5 / 3
	fmt.Printf("%v %v\n", num1, num2) // → 1 -1

	// 代入を短縮して書ける
	num1 += 5
	fmt.Printf("%v\n", num1) // → 6

	num1 *= 2
	fmt.Printf("%v\n", num1) // → 12

	num1 *= 2
	fmt.Printf("%v\n", num1) // → 12

	/*
	* ビット演算
	 */
	// 「（&）論理積」や「（|）論理和」や「（^）ビットの補数」などビット演算の理解は一旦スキップ
}