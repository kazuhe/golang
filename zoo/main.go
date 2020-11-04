package main

import (
	"fmt"

	// 通常は環境変数GOPATHで指定されたディレクトリ内のパッケージから探索されるが
	// 相対パスで記述することで相対位置に置かれたディレクトリを指定することも可能
	// ※この時にディレクトリ名とパッケージ名を揃える
	"./animals"
)

func main() {
	// 同パッケージであれば、別ファイルに分割していてもパッケージ名の指定は必要ない
	fmt.Println(AppName())

	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())
}
