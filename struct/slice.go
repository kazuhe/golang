package main

import "fmt"

/*
* スライスと構造体
 */

// []*Pointのように複雑な宣言に対してtypeによるエイリアスを定義し、
// そのエイリアスに対してメソッドを定義することで型を扱いやすくするテクニック
// 複雑なものを複雑なままで放置せず適宜分かりやすい型の表現を模索することが重要

// Points Pointのポインタ型を格納するスライス
type Points []*Point

// ToString 受け取ったPoints型のスライスを1つも文字列として返す
func (ps Points) ToString() string {
	str := ""
	for _, p := range ps {
		if str != "" {
			str += ","
		}
		if p == nil {
			str += "<nil>"
		} else {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
		}
	}
	return str
}

func callSlice() {
	ps2 := Points{}
	ps2 = append(ps2, &Point{X: 1, Y: 2})
	ps2 = append(ps2, nil)
	ps2 = append(ps2, &Point{X: 3, Y: 4})
	fmt.Println(ps2.ToString()) // [1, 2],<nil>,[3, 4]
}
