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
}
