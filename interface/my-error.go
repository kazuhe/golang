package main

import (
	"fmt"
)

/*
* 代表的なインターフェース error
 */

// 組み込み型error ↓ はインターフェースで、文字列を返すメソッドのErrorのみが定義されている
/*
type error interface {
	Error() string
}
*/

// errorインターフェースを実装した型を定義してみる↓流れ

// 構造体MyError型を定義
// errorインターフェースが要求する'Error() string'をメソッドのシグネチャ通りに定義する
// これらの定義によって関数'RaiseError'が返すerror型の値としてMyError型を使うことが可能になった

// MyError 独自定義のエラーを表す型
type MyError struct {
	Message string
	ErrCode int
}

// Error errorインターフェースのメソッドを実装
func (e *MyError) Error() string {
	return e.Message
}

// RaiseError MyError型のポインタを返す
func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

// 変数errの実際の型はMyErrorだが、プログラム上ではあくまでもerror型の変数であることに注意
// error型を経由してMyError型のフィールドにアクセスすることはできない
func callMyError() {
	err := RaiseError()
	fmt.Println(err.Error()) // エラーが発生しました
}
