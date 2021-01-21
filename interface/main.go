package main

/*
* インターフェース
 */
// インターフェースは型の一種で任意の型が「どのようなメソッドを実装するべきか」を規程するための枠組み
// 「interface {メソッドのシグネチャの列挙}」という形式で、型が実装するべきメソッドのシグネチャを任意の数だけ列挙できる

func main() {
	/*
	* 代表的なインターフェース error
	 */
	callMyError()

	/*
	* インターフェースのメリット
	 */
	callStringify()

	/*
	* インターフェースを含むインターフェース
	 */
	// 「インターフェースI1」は「インターフェースI0」を含んでいる為、I1が要求するメソッドは「Method1とMethod2」の2つになる
	type I0 interface {
		Method1() int
	}

	type I1 interface {
		I0 // インターフェースI0を含む
		Method2() int
	}
}
