package main

import "fmt"

/*
* インターフェースのメリット
 */
// インターフェースの最もポピュラーな使用方法は「異なる型に共通の性質を付与する」使い方

// Stringify 文字列化できることを示すインターフェース
type Stringify interface {
	ToString() string
}

// Person 人を表す構造体
type Person struct {
	Name string
	Age  int
}

// ToString Stringifyインターフェースのシグネチャ通りにstringを返すメソッドを定義
func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

// Car 車を表す構造体
type Car struct {
	Number string
	Model  string
}

// ToString Stringifyインターフェースのシグネチャ通りにstringを返すメソッドを定義
func (c *Car) ToString() string {
	return fmt.Sprintf("[%s] %s", c.Number, c.Model)
}

// また、インターフェースを活用して以下の様な汎用性の高い関数やメソッドを定義することができる

// Println Stringifyインターフェースさえ実装していればどの様な型を使っても呼び出すことができる
func Println(s Stringify) {
	fmt.Println(s.ToString())
}

// callStringify file内で分かりやすく呼び出す為の仮関数
func callStringify() {

	// 異なる型を共通するインターフェース型にまとめることができる
	vs := []Stringify{
		&Person{Name: "Taro", Age: 20},
		&Car{Number: "XXX-0123", Model: "PX512"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}
	// 出力結果
	// Taro(20)
	// [XXX-0123] PX512

	Println(&Person{Name: "Hanako", Age: 30})         // Hanako(30)
	Println(&Car{Number: "XYZ-9999", Model: "RT-12"}) // [XYZ-9999] RT-12
}
