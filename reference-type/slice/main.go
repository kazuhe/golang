package main

import "fmt"

func main() {
	/*
	* slice
	 */
	// スライスはGoで利用頻度が高いデータ構造で、可変長配列を実現する型と考えても良い

	// make(T, n) ← 要素数と容量がnであるT型のスライスを生成
	// make(T, n, m) ← 要素数がnで容量がmであるT型のスライスを生成

	// int型を格納するスライスであることを宣言
	var s []int
	fmt.Printf("%T\n", s)

	// 要素と容量が10であるint型のスライス
	s2 := make([]int, 10)
	fmt.Println(s2)

	// 「スライスの's2'」と「配列の'a'」の出力は同じだが、配列とスライスは実装上全く異なるもの
	a := [10]int{}
	fmt.Println(a)

	/*
	* 要素への代入と参照
	 */
	// 代入も参照も配列型と同様の操作
	// スライスの要素数を超過した要素へのアクセスはランタイムパニックになる
	a2 := make([]float64, 3)
	fmt.Println(a2)
	a2[0] = 3.14
	fmt.Println(a2)
	a2[1] = 6.28
	fmt.Println(a2)
	fmt.Println(a2[0])

	/*
	* len
	 */
	// スライスは可変長配列なので動的に要素数が変化するが組み込み関数のlenで調べる事ができる
	// 配列でもlenは使える（定義時に要素数が明示されているのであまり使う事はないが）
	s3 := make([]int, 8)
	fmt.Println(len(s3))

	/*
	* cap
	 */
	// スライスの容量を調べる為のcap
	// 要素数5 容量5
	s4 := make([]int, 5)
	fmt.Println(len(s4)) // 5
	fmt.Println(cap(s4)) // 5
	// 要素数7 容量10
	s5 := make([]int, 7, 10)
	fmt.Println(len(s5)) // 7
	fmt.Println(cap(s5)) // 10

	// ====================================================================
	// 容量はあらかじめメモリの領域を確保している数値
	// 元の容量以上に要素数を拡張しようとすると、元のスライスが格納していたデータを
	// 丸ごと新しい領域にコピーし別の領域に移動されるのでコストが高い処理になる
	// 容量を超えた要素の追加の度に容量の数が倍増する為できる限り容量の指定にも気を配る
	// ※ Goのランタイムは閾値によって容量の拡張幅が変わるので一概に倍増ではない
	// ====================================================================

	/*
	* スライスを生成するリテラル
	 */
	// makeを使用せず配列型のリテラルと同様の書き方で生成できる
	// 個別に容量は指定できないので要領は要素数と同一になる
	s6 := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s6)) // 5
	fmt.Println(cap(s6)) // 5
	fmt.Printf("%T\n", s6)

	/*
	* 簡易スライス式
	 */
	// 配列やスライスから新たにスライスを生成することができる
	// [n:m] n: インデックス m: m-1までの要素数
	a3 := [5]int{1, 2, 3, 4, 5}
	s7 := a3[0:2]
	fmt.Println(s7) // [1 2]

	// 簡易スライス式には任意の式を記述できる
	s8 := a3[len(a3)-2:]
	fmt.Println(s8) // [1 2]

	/*
	* 文字列と簡易スライス式
	 */
	// この操作は文字を単位とせず「バイト列」で計算する為、マルチバイト文字が混在する場合は注意
	s9 := "ABCDE"[1:3]
	fmt.Println(s9) // BC

	/*
	* append
	 */
	// 配列とスライスも最も大きな違い「拡張性」
	// 配列は要素数を含めて1つの型で[10]intと[11]intは全く別の型となる
	// スライスは可変長配列を表現するデータで要素数に制限がない

	// appendを使用する際は必ず変数の代入(:= or =)を伴う必要がある
	s10 := []int{1, 2, 3}
	s10 = append(s10, 4, 5)
	fmt.Println(s10) // [1 2 3 4 5]

	// スライスの末尾に「s11...」の特殊な書きかたでスライスを追加
	s11 := []int{6, 7}
	s10 = append(s10, s11...)
	fmt.Println(s10) // [1 2 3 4 5 6 7]

	/*
	* ざっくりまとめ
	 */
	// 配列は「[]」内に値を定義する
	array1 := [3]int{1, 2, 3}
	// sliceは「[]」内に値を定義しない
	slice1 := []int{1, 2, 3}

	// array1 = append(array1, 4, 5) ←配列なのでappendはエラー
	slice1 = append(slice1, 4, 5) // ←sliceなのでOK
	fmt.Println(array1)
	fmt.Println(slice1)

	/*
	* copy
	 */
	// 第1引数のスライス: コピー先
	// 第2引数のスライス: コピー元
	// コピー先のスライスを先頭から塗りつぶす様にコピーされる
	// copy関数はコピーが成功された要素数を整数の形式で返す
	s12 := []int{1, 2, 3, 4, 5}
	s13 := []int{10, 11}
	n := copy(s12, s13)
	fmt.Println(n, s12, s13) // 2 [10 11 3 4 5] [10 11]

	// コピー元のスライスの方が大きい場合でもコピー先の要素数分はコピー可能
	s14 := []int{1, 2, 3, 4, 5}
	s15 := []int{10, 11, 12, 13, 14, 15, 16}
	n2 := copy(s14, s15)
	fmt.Println(n2, s14, s15) // 5 [10 11 12 13 14] [10 11 12 13 14 15 16]

	/*
	* 完全スライス式
	 */
	// 簡易スライス式との違いは容量をコントロールできる点
	a4 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 簡易スライス式の場合は参照していない範囲が自動的に容量になる
	s16 := a4[2:4]        // [3 4]
	fmt.Println(len(s16)) // 2
	fmt.Println(cap(s16)) // 8

	// 完全スライス式の場合は「[]」内の第3引数が容量のmax
	// すなわち、第3引数 - 第3引数が容量となる
	s17 := a4[2:4:6]      // [3 4]
	fmt.Println(len(s17)) // 2
	fmt.Println(cap(s17)) // 4
}