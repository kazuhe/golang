package main

import (
	"fmt"

	"./foo"
)

func main() {

	// パッケージ名と識別子を「.(ドット)」で繋いで参照

	// fooパッケージの定数MAXを参照
	fmt.Printf("%v\n", foo.MAX)
	// ↓ internalConstはエクスポートされていないのでコンパイルエラー
	// fmt.Printf("%v\n", foo.internalConst)

	// fooパッケージの定数MAXを参照
	fmt.Printf("%v\n", foo.DoSomething())
	// fmt.Printf("%v\n", foo.doSomething()) → コンパイルエラー
}
