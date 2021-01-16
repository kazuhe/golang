package main

import (
	"fmt"
)

/*
* 型のコンストラクタ
 */

// Goにはオブジェクト指向言語の「コンストラクタ」機能はないが、慣例的に「型のコンストラクタ」というパターンを利用する
// 型のコンストラクタを表す関数は「New[型名]」のように命名するのが一般的
// また、対象の型のポインタ型を返すように定義するのが望ましい

// 型のコンストラクタをパッケージの外部に公開する場合は"NewUser"で問題ないが
// パッケージの内部でのみ利用する場合は"newUser"のようにキャメルケースがよい

// User 構造体"user"
type User struct {
	ID   int
	Name string
}

// NewUser 新しいUserを定義する型のコンストラクタ
func NewUser(id int, name string) *User {
	u := new(User)
	u.ID = id
	u.Name = name
	return u
}

func callNewUser() {
	taro := NewUser(1, "Taro")
	fmt.Println(taro) // &{1 Taro}

	taro.ID = 77
	fmt.Println(taro) // &{77 Taro}
}
