package main

import (
	"fmt"
	"reflect"
)

/*
* タグ
 */
// タグ（Tag）は構造体のフィールドにメタ情報を付与する機能
// タグには文字列リテラルかRAW文字列リテラルを使用できる
// タグ内のダブルクォーテーションが意味を持つ場合もあるためRAW文字列リテラルが好まれるケースが多い
// あくまでも単なる文字列なのでコンパイラーは教えてくれない

// TagUser 説明
type TagUser struct {
	ID   int    "ユーザー名"
	Name string "名前"
	Age  uint   "年齢"
}

// タグはあくまでも構造体のフィールドに付与するメタ情報で、文字列リテラルに問題がない限りはプログラムに影響は及ぼさない
// タグの利用方法はライブラリやプログラムによって異なるが、例としては構造体をJSONに変換する時のキー名として利用することがある

// callTag 構造体に付与されたタグを参照するプログラム例
func callTag() {
	u := TagUser{ID: 1, Name: "Taro", Age: 32}

	// 変数tはreflect.TypeOf(u)型
	t := reflect.TypeOf(u)

	// 構造体の全フィールドを処理するループ
	for i := 0; i < t.NumField(); i++ {
		// i番目のフィールドを取得
		f := t.Field(i)
		fmt.Println(f.Name, f.Tag)
	}

	// 出力結果
	// ID ユーザー名
	// Name 名前
	// Age 年齢
}
