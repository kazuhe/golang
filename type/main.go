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
	bb := byte(nn) // エラーにならずに「0」になる
	fmt.Println(bb)
	// 2進数で考えた時に、8ビットで表現できない数値に到達してしまい桁が1つ繰り上がる
	// 255 = 0 1111 1111
	// 256 = 1 0000 0000
	// ↑これをオーバーフロー（桁あふれ）という
	// オーバーフローが発生した場合は演算結果をラップアラウンドさせ今回の例では「0」と表現される

	/*
	* ラップアラウンドへの対策
	 */
	// uint32方で表現できる整数の最大値は43億弱
	ui1 := uint32(400000000)
	ui2 := uint32(4000000000)
	sum := ui1 + ui2
	fmt.Printf("%d + %d = %d\n", ui1, ui2, sum) // → 400000000 + 4000000000 = 105032704
	// ↑ラウンドアップされてしまう

	// ./check_overflow.goで定義した関数でオーバーフローのチェックをかける
	// この様な手法でラップアラウンドの問題を回避する
	checkOverflow(ui1, ui2)
	ui3 := uint32(300000000)
	checkOverflow(ui1, ui3)

	/*
	* 浮動少数点型
	 */
	// 浮動小数点リテラルを使用した暗黙的な変数定義はfloat64型に定まる
	f64 := 1.0 // float64
	fmt.Printf("型=%T\n", f64)
	// f32型の値を得る場合は明示的な型変換が必要
	f32 := float32(1.0) // float32
	fmt.Printf("型=%T\n", f32)

	// 整数への型変換を行うとゼロの方向に切り捨てられる
	f := 3.14
	f2 := int(f) // 3
	fmt.Println(f2)
	// Goでは自然にfoloat64が選択される様にデザインされているので基本的にfloat32は使用しない
	// なお現代のCPUでは演算コストを意識する必要がない為、精度が高い（表現できる数値幅が大きい）方が良い

	/*
	* rune型
	 */
	// Unicodeコードポイントを表す特殊な整数型
	// int32の別名として定義されている為32ビット符号付き整数と何ら異なるところは無い
	r := '松'              // 「'(シングルクォート)」で囲む
	fmt.Printf("%v\n", r) // → 26494

	/*
	* 文字列型
	 */
	s1 := "Goの文字列"
	s2 := "\n"
	s3 := ""
	s4 := "\u65e5本\U00008a9e" // → 日本語と表示される バックスラッシュを使用した特殊な文字の表現も可能
	fmt.Printf("%v, %v, %v, %v\n", s1, s2, s3, s4)

	// RAW文字列リテラル
	// 「`(バックスラッシュ)」で囲むと改行を保持して複数行に渡る文字列を書ける
	s5 := `
	Goの
	RAW文字列
	リテラル
	\n
	`
	// また、↑の様に「\n」など改行コードはRAW文字列リテラルだとそのまま「\n」と表示される
	// 「raw(生)」という意味
	fmt.Printf("%v\n", s5)

	/*
	* 配列型
	 */
	// ./array.go
	arrayTest()

	/*
	* interface{}とnil
	 */
	// ./interface.go
	interfaceTest()

	/*
	* 型アサーション
	 */
	// ./assertion.go
	// interface{}を引数とした全ての型と互換性がある関数
	anything(100)
	anything("ABC")
	anything(3.14)

	// 「x.(T)」を使って型アサーションする
	// i, isInt := x.(int)
	// ↑ 1番目の変数にはxの値が代入され、2番目の変数にはbool値(アサーションが成功すればtrue)が代入される
	// この型アサーションを使ってif分で型判定を行う
	x := interface{}("型アサーション")
	if x == nil {
		fmt.Println("x is nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Printf("x is integer: %d\n", i)
	} else if s, isString := x.(string); isString {
		fmt.Printf("x is string: %s\n", s)
	}

	// 「x.(T)」で1つの変数に代入する場合はxの値のみが代入される
	// ここで、型アサーションに失敗した場合にはランタイムエラーが発生するので、
	// 前述した変数を2つ与えた型アサーションをおこなう事が多い
	// foo := x.(int) ←xはint型にはなりえないので、エラーとなる
	foo := x.(string)
	fmt.Println(foo)
}
