# golang
Go言語勉強用のリポジトリ

## Goが目指すコト
Goは「Cの高速性とスクリプト言語の手軽さ」を両立させるコトを目指して設計されている。  
その為、所々に奇妙な挙動（セミコロン省略によるコンパイルエラー等）が見受けられるが、「コードの表現力を上げる」ことよりも「コンパイル速度を速くする」ことを優先した方が最終的には開発者の利益になると判断したと言える。

## Goの原則
- 1つのディレクトリには1つのパッケージのみ
- ディレクトリ名とパッケージ名は同一にしておく

## 基本コマンド
```bash
# Goのバージョン情報を確認
$ go version

# Goのビルドシステムに関係する環境変数の内容を確認
$ go env

# Goのソースコードを推奨される形式へ自動的に整形
# -n: 実行されるコマンドの表示（ファイルは書き換えない）
# -x: 実行されるコマンドの表示
$ go fmt [-n] [-x] [packages]

# Goのパッケージのドキュメントを参照
$ go dog [-u] [-c] [package.]symbol[.method]
$ go doc math/rand # (例)math/randパッケージを指定
$ go doc time.Time.Unix # (例)timeパッケージのTime型のUnixメソッドを指定

# ビルドプロセスを隠蔽しつつ直接プログラムを実行
$ go run hello.go

# カレントディレクトリ内の全てのGoファイルをビルド
$ go build

# -oオプションを使用して出力するファイル名を指定して実行ファイル形式にコンパイル
$ go build -o hello hello.go

# -vオプションを使用して個々のテスト結果も表示させる
$ go test -v ./zoo/animals
```

## 基本の言語仕様

### 識別子の命名規則
変数名や定数名・関数名は全てGoでは「識別子」になる。  
これらの識別子に利用できる文字種は半角英数などの範囲に限定されず「Unicodeにおいて『文字』または『数字』と定義されたもの」に「_(アンダースコア)」を加えたものが許容される。

### セミコロン
Goでは文の終端に「;（セミコロン）」が必要だが、コンパイラが自動で挿入してくれるので明示的に使用することほぼない。  
しかし、視認性を考慮して以下の様に整形していると最後の要素（"Suzuki"）が文末と判定される。その為に不要なセミコロンが挿入されbuidエラーになってしまう。この様な場合は最後の要素にも「,（カンマ）」を置くことで解決できる。
```go
a := [3]string{
  "Yamada",
  "Sato",
  "Suzuki" // 文末と判定される
}
```

## テスト
Goには標準でパッケージのテスト機能が組み込まれている。  
`***_test.go`のファイルはテストの為の特別なファイルとして扱われる。

## 定義済み識別子
| 識別子の種類 | 識別子 |
----|---- 
| 型 | bool, byte, complex64, complex128, error, float32, float64, int, int8, int16, int32, int64, rune, string, uint, uint8, uint16, uint32, uint64, uintptr |
| 定数 | true, false, iota |
|ゼロ値 | nil |
|関数 |	append, cap, close, complex, copy, delete, imag, len, make, new, panic, print, println, real, recover |


## Goが認識するファイル
| 拡張子 | 内容 |
----|---- 
| .go | Goのソースファイル |
| .c / .h | Cのソースファイル |
| .cc / .cpp / .cxx / .hh / .hpp / .hxx | C++のソースファイル |
| .m | Objective-Cのソースファイル |
| .s / .S | アセンブラソースファイル |
| .swig / .swigcxx | SWIG定義ファイル |
| .syso | システムオブジェクトファイル |

## 参考書籍
[スターティングGo言語](https://www.shoeisha.co.jp/book/detail/9784798142418)