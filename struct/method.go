package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

// メソッド定義では関数とは異なりfuncとメソッド名の間に「レシーバー（Receiver）」の型とその変数名が必要になる
// 下記の例では「(p *Point)」がそれに該当する

// Rendar メソッドの説明用関数
func (p *Point) Rendar() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

func rendarCall() {
	// 型に定義されたメソッドは「p4.Rendar()」のような形式で呼び出すことができる
	p4 := &Point{X: 5, Y: 12}
	p4.Rendar() // <5, 12>
}

// メソッドには関数と同様に任意の引数と戻り値を定義できる
// レシーバーの定義が必要なことをのぞけば通常の関数と異なる点はない

// Distance 2点間の距離を求めるメソッド
func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func distanceCall() {
	p := &Point{X: 0, Y: 0}
	// Point型に特化したメソッド「Distance」を呼び出す
	distance := p.Distance(&Point{X: 1, Y: 1})
	fmt.Println(distance) // 1.4142135623730951
}

// 同一パッケージ内に同名の関数を複数定義することはできないが、メソッドはレシーバーの型さえ異なっていれば複数定義できる
