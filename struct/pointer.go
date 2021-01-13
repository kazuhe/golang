package main

// DoubleInt 構造体とポインタを使う例の型
type DoubleInt struct {
	X, Y int
}

func swap(p DoubleInt) {
	// フィールドX, Yの値を入れ替える
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

// ↑の関数をポインタを受け取れるように変更
func pointerSwap(p *DoubleInt) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}
