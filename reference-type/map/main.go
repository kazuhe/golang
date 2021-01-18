package main

import (
	"fmt"
)

func main() {
	/*
	* map
	 */
	// mapはいわゆる連想配列に類するデータ構造「map[キーの型]要素の型」
	// ↓はint型のキーとstring型の値を保持するマップ
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	m[86] = "China"
	fmt.Println(m) // map[1:US 81:Japan 86:China]

	// マップは一意に定まるキーと任意の要素の組み合わせで構成される
	// よってキーが重複する代入を行うと要素は上書きされる
	// キーを浮動小数点数の型なし定数の場合は自動的にまるめられるのでキーの重複に注意
	m[81] = "追加Japan"
	fmt.Println(m) // map[1:US 81:追加Japan 86:China]

	/*
	* マップのリテラル
	 */
	// 配列やスライス同様にマップもキーと要素のペアとまとめて生成する為のリテラルが用意されている
	m2 := map[int]string{1: "Taro", 2: "Hanako", 3: "Jiro"}
	fmt.Println(m2) // map[1:Taro 2:Hanako 3:Jiro]

	// マップを生成するリテラルの内部にさらにスライスを生成するリテラルなどを含める複雑な定義も可能
	// キー: int 値: []int(スライス)
	m3 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	fmt.Println(m3) // map[1:[1] 2:[1 2] 3:[1 2 3]]

	// ↑のマップは要素である[]int型のリテラルを省略可能
	m4 := map[int][]int{
		1: {1},
		2: {1, 2},
		3: {1, 2, 3},
	}
	fmt.Println(m4) // map[1:[1] 2:[1 2] 3:[1 2 3]]

	// これはマップの要素型がマップの場合も同じで省略記法を使う事で各要素の見通しがよくなる
	m5 := map[int]map[float64]string{
		1: map[float64]string{3.14: "円周率"}, // 省略しない場合
		2: {1.6: "省略記法"},                   // 省略する場合
	}
	fmt.Println(m5) // map[1:map[3.14:円周率] 2:map[1.6:省略記法]]

	// 構造体やスライスも同様に型名の省略が可能

	/*
	* 要素の参照
	 */
	// マップの要素を参照するには[]にキーの値を指定する
	// 登録されていないキーを参照した場合はそれぞれの型の初期値（stringの場合は""）が与えられる
	m6 := map[int]string{1: "A", 2: "B", 3: "C"}
	fmt.Println(m6[1]) // A
	fmt.Println(m6[9]) // ""

	// Goの基本型はnilのような特別な状態を持たないので意図した代入が実行されずミスに繋がりやすい
	// この様な問題を防ぐ為にマップの要素参照するには他の↓の様な方法がある
	s, ok := m[9]
	fmt.Println(s, ok) // "" false
	// 2つの変数に代入して第2引数に、存在すればtrue 存在しなければfalse が代入される
	// これで明確に値が存在するか確認することができる
	// また、この第2引数をokとすることはGoにおける一種のイディオムとされている

	// ifの簡易式とマップの要素参照を組み合わせた例
	if _, ok := m6[1]; ok {
		fmt.Printf("[%v]値は存在する\n", ok)
	}

	// Goのマップは要素の値にnilを使用できる
	m7 := map[int][]int{
		1: nil,
		2: {1, 2},
	}
	fmt.Println(m7) // map[1:[] 2:[1 2]]
	// 参照型の初期値がnilであることを利用して↓の様にnilを使って要素チェックするのはプログラムとして成り立たない
	s2 := m7[1]
	if s2 != nil {
		fmt.Printf("[%v]は初期値ではない\n", s2)
	}
	// 先述の変数okを使用して明確な要素チェックを行うべき

	/*
	* マップとfor
	 */
	// スライスで有用な「範囲節によるfor」はマップでも同じ様に活躍する
	m8 := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Cherry",
	}
	// 第1引数: キー 第2引数: バリュー
	for k, v := range m8 {
		fmt.Printf("%d => %s\n", k, v)
	}
	// 出力結果
	// 2 => Banana
	// 3 => Cherry
	// 1 => Apple

	// 「範囲節によるfor」を使用する場合は「キーの順序は保証されない」事に注意
	// マップは要素の追加や削除といった副作用から仕様の上でも「不定」である

	/*
	* len
	 */
	// capはマップには使用不可
	fmt.Println(len(m8)) // 3
	m8[4] = "Grape"
	fmt.Println(len(m8)) // 4

	/*
	* delete
	 */
	// マップから任意の要素を取り除く
	// キーが存在すれば除去され、存在しなければ何もされない
	m9 := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}
	// delete(マップ, キー)
	delete(m9, 2)
	fmt.Println(m9) // map[1:A 3:C]

	/*
	* 要素数に最適化したmake
	 */
	// makeを使ってマップを生成する場合に2番目の引数に「要素数に対応した初期スペース」を整数で指定できる
	// スライスにおける「容量（capacity）」とは意味合いが異なる
	// Goのランタイムが最適なメモリ領域を確保するための一種の「ヒント」として機能する
	// 要素数の巨大なマップを構築する場合であればパフォーマンスの向上を期待できるかもしれないが、
	// 要素数の少ないマップに細かく指定するのは煩雑なだけで意味をなさない
	m10 := make(map[int]string, 100) // 初期スペース100
	fmt.Println(m10)

}
