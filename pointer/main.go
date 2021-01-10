package main

import (
	"fmt"
)

func main() {
	/*
	*  ポインタ
	 */
	// ポインタとは「値型（value type）」に分類されるデータ構造（基本型や参照型や構造体など）のメモリ上のアドレスと型の情報
	// Goではポインタを使ってデータ構造を間接的に参照・操作できる
	// （Goのポインタを"複雑に"使わずとも大きな支障はきたさない、存在意義は"Cとの互換性の為"との意見も）

	/*
	*  ポインタの定義
	 */
	// ポインタ型は「*int」の様にポインタを使って参照・操作する型の前に「*」を置くことで定義できる
	// また、定義のみを行ったポインタ型の変数の初期値はnilになり、参照型と同じ様に振舞う
	var p *int
	fmt.Println(p == nil) // true

	// ↓の様にint型のポインタのポインタを参照・操作する為のポインタ型も定義できるが、Goではこの様な複雑な定義が必要になることは無い
	// var p ***int

	// 参照型にもポインタ型を定義できるが、参照型自身が型の仕組みにポインタを使った参照を含んでいる
	// 参照型へのポインタが必要になるケースは相当なレアケースに限られる
	var (
		s  *[]string
		m  *map[int]rune
		ch *chan int
	)
	fmt.Println(s, m, ch)

	/*
	*  アドレス演算子とデリファレンス
	 */
	// 演算子「&」はアドレス演算子と呼ばれ、これを使って任意の型からそのポイント型を生成することができる
	var i int
	p2 := &i
	fmt.Printf("%T\n", p2) // *int
	p3 := &p2
	fmt.Printf("%T\n", p3) // **int

	// 「*」をポインタ型の変数の前に置くことでポインタが指し示すデータのデリファレンスができる
	// デリファレンスとは、ポインタ型が保持するメモリ上のアドレスを経由して、データ本体を参照する為の仕組み
	i = 5
	fmt.Println(*p2) // 5

	// ポインタ変数を書き換えることで間接的に参照先の値を変更することができる
	*p2 = 10
	fmt.Println(i) // 10

	// ポインタの性質を利用すれば関数の引数へ値型の参照渡しが実現できる
	// 本来int型のような値型を渡す場合は関数呼び出しの際に引数のコピーが生成されてしまい、同じ領域の値を共有することはできない
	// しかし、ポインタ型を介すれば1つのメモリ上の値を共有することができる
	i2 := 1
	increment(&i2)
	increment(&i2)
	increment(&i2)
	fmt.Println(i2) // 4

	// ポインタを介さない例では当然i3自体の値に変化はない
	i3 := 1
	thatWayIncrement(i3)
	thatWayIncrement(i3)
	thatWayIncrement(i3)
	fmt.Println(i3) // 1

	// 任意の配列型に対してもポインタ型を定義できる
	p4 := &[3]int{1, 2, 3}
	pow(p4)
	fmt.Println(p4) // &[1 4 9]

	// ポインタ型の変数がnilである場合にデリファレンスを実行するとランタイムパニックが発生する
	// var p5 *int
	// fmt.Println(*p5) → ランタイムパニック

	/*
	* 配列へのポインタ型
	 */
	// ↓の(*p)[0]のような書き方はC由来
	p5 := &[3]int{1, 2, 3}
	fmt.Println((*p5)[0]) // 1

	// Goではp6[i]がポインタ型であることをコンパイラが検知して自動的に(*p6)[i]の形式へ置き換えている
	// この仕組みで簡潔に記述することができる
	a := [3]string{"Apple", "Banana", "Cherry"}
	p6 := &a           // 型は → *[3]string になる
	fmt.Println(a[1])  // Banana
	fmt.Println(p6[1]) // Banana
	p6[2] = "Grape"
	fmt.Println(a[2])  // Grape
	fmt.Println(p6[2]) // Grape

	// lenやcap、スライス式の場合も配列へのポインタ型であればデリファレンスを省略できる
	p7 := &[3]int{1, 2, 3}
	fmt.Println(len(p7)) // 3
	fmt.Println(cap(p7)) // 3
	fmt.Println(p7[0:2]) // [1 2]

	// rangeを使用する場合も配列へのポインタ型のデリファレンスを省略できる
	// この様にGoは配列へのポインタ型を簡潔に操作するための仕組みが備わっている
	p8 := &[3]string{"Apple", "Banana", "Cherry"}
	for i, v := range p8 {
		fmt.Println(i, v)
	}

	// Goが合理的であれば一貫性にはあまり強くこだわらない傾向がある
	// 対象が配列型なのか、配列へのポインタ型なのか見分けがつきにくくなるデメリットがあるが合理性を優先している
	// 配列へのポインタ型はある種「特別扱い」されていると認識しておけば良い

	/*
	* 値としてのポインタ型
	 */
	// ポインタ型は「型」とメモリ上の「アドレス」を組み合わせたデータ型
	// int方が整数を持つように、ポインタ型は値として「メモリ上のアドレス」を保持する
	i4 := 5
	ip := &i4
	fmt.Printf("type=%T, address=%p\n", ip, ip) // type=*int, address=0xc0000b4050

	/*
	* 文字列型とポインタ型
	 */
	// string型はポインタ型という観点から見ると特殊
	// 文字列の部分参照s[0]をポインタ参照するとコンパイルエラーになる
	// これはGoでは1度生成された文字列に対して何らかの変更を加えることは基本的にできないように設計されている「immutable」な性質を持ち、
	// 文字列を構成するbyte型の配列に対してポインタ型を定義できると破壊的変更を許可することに他ならない為
	s2 := "ABC"
	sp := &s2
	fmt.Println(sp)    // 0xc000010270
	fmt.Println(s2[0]) // 65 （byte型）
	// fmt.Println(sp[0]) → コンパイルエラー

	/*
	* まとめ
	 */
	// Goでは「この様な場合にポインタ型を使う」といったルールさえ把握しておけば通常のプログラミングには支障はほとんどない
	// 学習を進めてポインタについてより本質的な理解が必要になった際にCとあわせて学んでみるが良いとのこと

}
