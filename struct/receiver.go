package main

import (
	"fmt"
)

/*
*  レシーバーとポインタ型
 */

// 基本的には「構造体に定義するメソッドのレシーバーはポインタ型」にすべき
// ここではレシーバーを値型にした場合との差異を確認する

// PairInt intのペア
type PairInt struct{ X, Y int }

// Set 値型のPairInt型のレシーバー
func (p PairInt) Set(x, y int) {
	p.X = x
	p.Y = y
}

// PointerSet ポインタ型のPairInt型のレシーバー
func (p *PairInt) PointerSet(x, y int) {
	p.X = x
	p.Y = y
}

// メソッドに対してレシーバーが値渡しされるか参照渡しされるかの違いは、レシーバーの型が値型かポインタ型かによってのみ決定される
// 呼び出し側はポインタ型だろうが値型でも関係ない
func callReceiver() {
	// PairInt型
	p1 := PairInt{}
	p1.Set(1, 2)
	fmt.Println(p1.X, p1.Y) // 0 0
	p1.PointerSet(1, 2)
	fmt.Println(p1.X, p1.Y) // 1 2

	// *PairInt型
	p2 := &PairInt{}
	p2.Set(1, 2)
	fmt.Println(p2.X, p2.Y) // 0 0
	p2.PointerSet(1, 2)
	fmt.Println(p2.X, p2.Y) // 1 2
}
