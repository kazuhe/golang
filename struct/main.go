package main

import (
	"fmt"
)

func main() {
	/*
	* 構造体
	 */
	// 「構造体（struct）」とは複数の任意の型の値を1つにまとめたもの
	// Goでは複雑なデータ構造を1つの「型」として取り扱うことができる構造体の定義は重要な位置を占める

	/*
	* type
	 */
	// typeは既に定義されている型をもとに新しい型を定義するための機能
	// 基本型のint型にMyIntのようなエイリアスを定義する
	type MyInt int
	var n1 MyInt = 5
	n2 := MyInt(7)
	fmt.Println(n1) // 5
	fmt.Println(n2) // 7

	// メリットはmap[string][2]float64のような複雑な型にエイリアスを定義し見通しがよくなる
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)
	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)
	fmt.Println(pair) // [1 2]
	fmt.Println(strs) // [Apple Banana Cherry]
	fmt.Println(amap) // map[Tokyo:[35.689488 139.691706]]
	fmt.Println(ich)  // 0xc00008c060

	// 関数型にエイリアスを定義することもできる
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	fmt.Println(n) // 30

	/*
	* エイリアス型の互換性
	 */
	// ↓のT0とT1はいずれもint型へのエイリアスでこれらはint型への互換性が成立する
	type T0 int
	type T1 int

	t0 := T0(5)
	i0 := int(t0)
	fmt.Printf("t0=%T, i0=%T\n", t0, i0) // t0=main.T0, i0=int

	t1 := T1(5)
	i1 := int(t1)
	fmt.Printf("t1=%T, i1=%T\n", t1, i1) // t1=main.T1, i1=int

	// T0とT1の間に互換性は成立しない
	// t0 = t1 → コンパイルエラー

	/*
	* 構造体の定義
	 */
	// 構造体を使用するには一般的にはtypeと組み合わせて新しい型を定義する
	// ※structで定義された構造体にtypeを使って新しい型名を与えるという順序
	// structの中には任意の型を持つ「フィールド」を任意の数だけ並べることができる
	type Point struct {
		X, Y int
	}

	// 構造体は値型の一種で、構造体の変数定義時にメモリ領域が確保されフィールドは初期値をとる
	// 構造体のフィールドの値を参照するには「構造体.フィールド名」のように指定する
	var pt Point
	fmt.Println(pt.X, pt.Y) // 0 0

	// フィールドに値を代入することもできる
	pt.X = 10
	pt.Y = 8
	fmt.Println(pt.X, pt.Y) // 10 8

	/*
	* 複合リテラル
	 */
	// 構造体に各フィールドの初期値を指定しつつ生成する為の「複合リテラル（Composite literals）」が用意されている
	// この時構造体のフィールドが定義された順序に対応しなければならない
	pt2 := Point{1, 2}
	fmt.Println(pt2.X, pt2.Y) // 1 2

	// フィールドを明示的に指定して値を定義することもできる
	pt3 := Point{X: 5, Y: 7}
	fmt.Println(pt3.X, pt3.Y) // 5 7
	// この場合は順序を気にする必要がなく'Point{Y: 7, X: 5}'のように順序を入れ替えても問題ない

	// また、明示的に指定する複合リテラルを使うことでフィールドの一部のみを初期化することもできる
	// 指定されなかった方は構造体の定義時の初期値のままである
	// 可読性と柔軟性を考慮して基本的には明示的に指定する方法で記述するべき
	pt4 := Point{Y: 33}
	fmt.Println(pt4.X, pt4.Y) // 0 33

	/*
	 * 構造体を含む構造体
	 */
	// ＊フィールド名のある埋め込み構造体
	// 先頭が英大文字の英数字によるフィールド名がGoの慣例
	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed Feed
	}
	// ↑の Feed Feed の定義を
	// type Animal struct {
	// 	Name string
	// 	Feed
	// }
	// ↑のように省略することができる
	// この場合、「フィールド名が一意に定まる」の条件にクリアしたフィールドは「a.Amount」のように中間フィールド名を省略することができる

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	// a.Feed.Name のように階層的にたどってアクセスすることができる
	fmt.Println(a.Name)        // Monkey
	fmt.Println(a.Feed.Name)   // Banana
	fmt.Println(a.Feed.Amount) // 10

	a.Feed.Amount = 15
	fmt.Println(a.Feed.Amount) // 15

	// フィールド名を省略した埋め込み構造体は異なる構造体型に共通の性質をもたせることができる
	type Base struct {
		ID    int
		Owner string
	}

	type A struct {
		Base // 共通フィールド
		Name string
		Area string
	}

	type B struct {
		Base   // 共通フィールド
		Title  string
		Bodies []string
	}

	sa := A{
		Base: Base{17, "taro"},
		Name: "Taro",
		Area: "Tokyo",
	}

	sb := B{
		Base:   Base{81, "Hanako"},
		Title:  "no title",
		Bodies: []string{"A", "B"},
	}

	fmt.Println(sa)    // {{17 taro} Taro Tokyo}
	fmt.Println(sa.ID) // 17

	fmt.Println(sb)    // {{81 Hanako} no title [A B]}
	fmt.Println(sb.ID) // 81

	// ＊暗黙的にフィールドの注意点
	// ↓ ポインタ型の修飾子やパッケージ名のプリフィックス部分は無視され純粋な型名の部分が暗黙的なフィールド名として利用される

	// struct {
	// 	T1    // フィールド名は「T1」
	// 	*T2   // フィールド名は「T2」
	// 	P.T3  // フィールド名は「T3」
	// 	*P.T4 // フィールド名は「T4」
	// }

	// また、構造体のフィールドに構造体自身の型を埋め込むような再起的な定義はコンパイルエラーとなる

	/*
	 * 無名の構造体型
	 */
	// typedで定義していないstructそのものを型として扱うこともできる
	// しかし、煩雑な定義をあえてする必要性はほぼなく原則typeと組み合わせてエイリアスを定義するべき
	noName := struct{ X, Y int }{X: 1, Y: 2}
	noNameStruct(noName) // {1 2}

	/*
	 * 構造体とポインタ
	 */
	// 構造体は値型なので関数の引数として渡した場合はコピーが生成され元の構造体に影響はない
	// 以下の例でp1の値に何も影響がないのがわかる
	p1 := DoubleInt{X: 5, Y: 10}
	swap(p1)
	fmt.Println(p1.X, p1.Y) // 5 10

	// 構造体は主にポインタ型を経由して使用する
	// 構造体型のポインタを生成してポインタを受け取れる関数へ渡す
	// p2の値が変更されていることがわかる
	p2 := &DoubleInt{X: 5, Y: 10}
	pointerSwap(p2)
	fmt.Println(p2.X, p2.Y) // 10 5

	/*
	* new
	 */
	// 指定した型のポインタ型を生成する
	// &演算子を使った複合リテラルによる構造体型のポインタ生成（&DoubleInt{X: 5, Y: 10}）とnewは動作上ほとんど違いがない
	// 状況に応じて使い分ける必要がある
	type Person struct {
		ID   int
		Name string
		Area string
	}
	p3 := new(Person)

	fmt.Printf("%v\n", p3.ID)   // 0
	fmt.Printf("%v\n", p3.Name) // ""
	fmt.Printf("%v\n", p3.Area) // ""

	/*
	* メソッド
	 */
	// オブジェクト指向言語によくあるメソッドとは異なる
	// 任意の型に特化した関数を定義する為の仕組み

	// メソッドの定義
	rendarCall()
	distanceCall()
	callPlus()
	callJoin()

	/*
	* 型のコンストラクタ
	 */
	callNewUser()

}
