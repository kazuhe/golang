package foo

/*
* 識別子の命名規則による可視範囲
 */
// パッケージに定義された定数・関数・変数などが他のパッケージから参照可能であるかは
// 「識別子の1文字目が大文字」であるかどうかで決定される

// 定数
const (
	MAX           = 1 // 公開される定数
	internalConst = 1 // パッケージ内のみで参照できる定数
)

// パッケージ変数
var (
	m = 256 // パッケージ内のみで参照できる変数
	N = 512 // 公開される変数
)

// DoSomething は公開される関数
func DoSomething() string {
	return "foo"
}

// doSomething はパッケージ内でのみ参照できる関数
func doSomething() string {
	return "foo"
}
