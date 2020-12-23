package main

import (
	"fmt"
)

// recoverによってpanicから復帰させるテスト関数
// ある種の例外処理が実現できているが、panicを起こす可能性があることを
// 前提にデザインされる様なことは絶対に避けるべき
func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
		}
	}()
	panic(src)
}

func main() {
	/*
	* panic
	 */
	// panicを実行すると即座に「ランタイムパニック」が発生し実行中の関数が中断される
	// これ以上処理を継続しようがない状態、すなわち回復不能な事態を表すので一般的なエラーとは区別する
	// また、中断時までに登録されたdeferは全て実行される

	// panic("runtime panic!!!") →以降の処理は中断される
	// fmt.Println("Hello")

	/*
	* recover
	 */
	// panicによって中断されたプログラムの中断を回復する
	// recoverはdeferと組み合わせて使うのが原則
	// panicは中断後にdeferに登録された式の評価に移行する為deferの中でしか動作しない
	// recoverはinterface{}型の値を戻しnilでなければpanicが実行されたと判断する
	testRecover(128)
	testRecover("foo")
	testRecover([...]int{1, 2, 3})
}
