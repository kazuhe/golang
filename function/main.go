package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", plus(1, 2))
	hello()
	q, r := div(19, 7)
	fmt.Printf("q=%v r=%v\n", q, r)

	/*
	* 戻り値の破棄
	 */
	//「_」を使って戻り値の一部を破棄することが可能
	q2, _ := div(19, 7)
	fmt.Printf("q2=%v\n", q2)

	/*
	* 関数とエラー処理
	 */
	// Goには例外機構がない為複数の戻り値を使用してエラーを検知する
	result, err := doError()
	if err != nil {
		// エラー処理
		fmt.Printf("result=%v err=%v\n", result, err)
	}
	// 第二引数のerrでエラーの有無を表す
	// これはGoの一種のイディオムで頻出する表現
	// 変数名がerrなのも慣例となっているのでこの形式に従う
}
