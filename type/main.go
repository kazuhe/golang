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

	// 符号なし整数型 uint8(byte), uint16, uint32, uint64
	var u8 uint8
	u8 = 10
	fmt.Printf("型=%T\n", u8)

	// 実装に依存する整数型 int, uint, uintptr
	n := 9223372036854775807 // 符号付き64ビット整数で表現可能な最大数
	// ↑コレを32ビット実装のGoでコンパイルしようとするとエラーが発生する
	fmt.Printf("型=%T\n", n)

	/*
	* 整数の型変換（キャスト）
	 */
	// Goは暗黙的な型変換は許容されないが、明示的な型変換は可能
	num := 17 // 整数リテラルを暗黙的に定義した場合はint型に定まる
	fmt.Printf("型=%T\n", num)

	// numをuint32型へキャスト
	u32 := uint32(num)
	fmt.Printf("型=%T\n", u32)

	/*
	* 整数の型変換による問題とラップアラウンド
	 */
	// byte型は0~255までの整数を表現できるが、それを超えて代入しようとするとエラーが発生する
	// bb := byte(256) →エラー

	// しかし、int型を経由した後
	nn := 256
	bb := byte(nn)
	fmt.Println(bb)
}
