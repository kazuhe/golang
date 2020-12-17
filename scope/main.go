package main

import (
	"fmt"
	m "math" // mathパッケージにパッケージ名「m」を指定

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

	/*
	 * パッケージ名の省略
	 */
	// パッケージ名「m」を指定してimport時したmathパッケージを使う
	// この時に本来の「math」というパッケージ名は使用不可
	fmt.Println(m.Pi)
}
