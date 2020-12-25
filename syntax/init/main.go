package main

import (
	"fmt"
)

/*
* init
 */
// initはパッケージの初期化関数でmain関数に先立って実行される
// 引数もとらず、戻り値も返さない

func init() {
	fmt.Println("init()")
}

// initは同名で複数定義することができる
// 同一パッケージ内で別のファイルに初期化処理を書く際に有効
// なお、実行順はソースコードに出現した順になる
func init() {
	fmt.Println("init() 2")
}

func main() {
	fmt.Println("main()")
}
