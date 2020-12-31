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
}
