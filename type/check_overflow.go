package main

import (
	"fmt"
	"math"
)

// mathパッケージを利用してオーバーフローが発生する可能性をチェックする関数を用意
func checkOverflow(a, b uint32) bool {
	if (math.MaxUint32 - a) < b {
		// オーバーフローするのでfalse
		fmt.Println("オーバーフロー")
		return false
	}
	// チェック済みの為問題無し
	sum := a + b
	fmt.Println(sum)

	return true
}
