package main

import "fmt"

func receiver(ch <-chan int) {
	// main関数内のゴルーチンから"受信した整数を出力し続ける"
	for {
		i := <-ch
		fmt.Printf("receiver関数による出力=%d\n", i)
	}
}

func main() {
	/*
	* chan
	 */
	// チャネルはゴルーチンとゴルーチンの間でデータの受け渡しを司るためにデザインされたGo特有のデータ構造

	/*
	* チャネルの型
	 */
	// チャネルの型名はchanと型定義の間にスペースが含まれる
	var ch chan int        // int型のチャネルを表す
	fmt.Printf("%T\n", ch) // chan int

	// チャネルにはGoの他のデータ型には存在しない特殊なサブタイプを指定できる
	// 「<-chan」は受信専用チャネルで「chan<-」は送信専用チャネル
	// オプションを指定しない場合は受診も送信も可能な双方向のチャネルとして機能する
	var ch1 <-chan int      // 受信専用
	var ch2 chan<- int      // 送信専用
	fmt.Printf("%T\n", ch1) // <-chan int
	fmt.Printf("%T\n", ch2) // chan<- int

	// Goのデータ型は厳密で異なる型の変数同士の代入は原則的にコンパイルエラーとなるがチャネルのサブタイプについては少し異なる
	// サブタイプの指定がないchanはサブタイプ付きのchanに代入可能だが、それ以外のサブタイプを越えて代入はできない
	// 基本はサブタイプ指定がないchanで局面に応じて受信専用or送信専用に切り替えることを意図した仕組みである
	ch1 = ch
	ch2 = ch
	// ch = ch1 ←NG
	// ch = ch2 ←NG
	// ch1 = ch2 ←NG
	// ch2 = ch1 ←NG

	/*
	* チャネルの生成と送受信
	 */
	// チャネルはmake関数によって生成でき、第2引数にバッファサイズを指定できる(何も指定しなければ0となる)
	// バッファサイズは「キュー」を格納する領域で先入先出の性質を持つ
	ch3 := make(chan int, 8) // バッファサイズ8のチャネル
	fmt.Printf("%T\n", ch3)  // chan int

	// チャネルが保持するデータに対する操作は「送信」か「受信」のみで送受信ともに <- を使用する
	ch3 <- 5       // ch3チャネルに整数5を送信
	i := <-ch3     // ch3チャネルから整数値を受信
	fmt.Println(i) // 5

	/*
	* チャネルとゴルーチン
	 */
	// チャネルはあくまでもゴルーチン間でデータを共有する為の仕組みである
	// その為↓の様に1つだけ存在するゴルーチンがチャネルからの受信の為に
	// 眠ってしまって「デッドロック」であると検知しランタイムエラーが発生する
	ch4 := make(chan int)
	// fmt.Println(<-ch4) // エラー「all goroutines are asleep - deadlock!」
	// つまり、複数のゴルーチン間でデータを共有する為に複数のゴルーチンが存在する必要がある

	// goを使ってゴルーチンを生成し、関数receiverにチャネルを共有
	go receiver(ch4)
	// 1ずつ増分する整数をチャネルに"送信し続ける"
	for i := 0; i < 10; i++ {
		ch4 <- i
	}

	// 出力結果
	// receiver関数による出力=0
	// receiver関数による出力=1
	// receiver関数による出力=2
	// receiver関数による出力=3
	// receiver関数による出力=4
	// receiver関数による出力=5
	// receiver関数による出力=6
	// receiver関数による出力=7
	// receiver関数による出力=8
	// receiver関数による出力=9

	// チャネル処理を行っているゴルーチンが停止するかどうかはバッファサイズにもよる
	// ↓の例では4つめのデータ送信処理でランタイムパニックが発生する
	ch5 := make(chan rune, 3)
	ch5 <- 'A'
	ch5 <- 'B'
	ch5 <- 'C'
	// ch5 <- 'D' エラー「デッドロック」

	// ゴルーチンがチャネル操作によって停止する条件は以下2パターン
	// 「バッファサイズが0またはバッファ内が空のチャネルからの受信」
	// 「バッファサイズが0またはバッファ内に空きがないチャネルへの送信」

	/*
	* len
	 */
	// lenによって取得できるのはチャネルのバッファ内に留められているデータの個数
	ch6 := make(chan string, 3)
	ch6 <- "Apple"
	fmt.Println(len(ch6)) // 1
	ch6 <- "Banana"
	fmt.Println(len(ch6)) // 2

	// プログラムの実行時と次の瞬間以降で他のゴルーチンによってチャネルの状態が変化する可能性があるので↓の様なコードは避けるべき
	// 【 NG例 】
	if len(ch6) > 0 {
		i := <-ch6
		fmt.Println(i)
	}

	/*
	* cap
	 */
	// capによって取得できるのはチャネルのバッファサイズ
	// 1度生成したチャネルのバッファサイズが変動することはないので有効に利用できる局面は少ない？
	fmt.Println(cap(ch6)) // 3

	/*
	* close
	 */
	// チャネルは「オープン」状態から始まるが、送信したチャネルを明示的に「クローズ」させることができる
	// クローズされたチャネルに対して送信を行うとランタイムパニックが発生する
	ch7 := make(chan int, 3)
	ch7 <- 7
	close(ch7)
	// ch7 <- 1 エラー「panic: send on closed channel」

	// しかし、クローズされたチャネルから受信をおこなうことは可能
	fmt.Println(<-ch7) // 7
	fmt.Println(<-ch7) // 0(バッファ内が空になっても初期値がを受信し続ける)

	// チャネルがクローズされているか確認するには2つの変数を割り当てる
	// 第1引数: チャネル内の値 第2引数: チャネルの「バッファ内が空でかつクローズされた状態」であればfalse
	i, ok := <-ch7
	fmt.Println(i, ok) // 0 false

	/*
	* チャネルとfor(range)
	 */
	// チャネルからひたすらrangeによって受信し続けることもできるが、クローズを検出するタイミングが得られずエラーになる
	ch8 := make(chan int, 3)
	ch8 <- 111
	ch8 <- 222
	ch8 <- 333
	// for i := range ch8 {
	// 	fmt.Println(i)
	// }
	// ↑ fatal error: all goroutines are asleep - deadlock!

	/*
	* select
	 */
	// 1つのチャネルが受信待ちの状態で停止してしまうとその後に記述されたチャネルはいつまでも受信できない
	// そんな問題を解決する為に複数のチャネルをコントロールするselect構文が用意されている
	// select分のcase節の式はすべてチャネルへの処理を伴っている必要がある

	// チャネルへの処理とは
	// <-ch による受信処理
	// ch <- e による送信処理
	// ch1 <- ch2 の様なチャネル間を直接繋ぐ処理

	ch9 := make(chan int, 1)
	ch10 := make(chan int, 1)
	ch11 := make(chan int, 1)
	ch9 <- 9
	ch10 <- 10

	select {
	case e9 := <-ch9:
		// ch9から受信が成功した場合に処理される
		fmt.Println(e9)
		fmt.Println("ch9から受信")
	case e10 := <-ch10:
		// ch10から受信が成功した場合に処理される
		fmt.Println(e10)
		fmt.Println("ch10から受信")
	case ch11 <- 11:
		// ch11へ送信が成功した場合に処理される
		fmt.Println(<-ch11)
		fmt.Println("ch11へ送信")
	default:
		// case節の条件が成立しなかった場合に処理される
		fmt.Println("ここへは到達しない")
	}
	// switchと違い、selectは複数のcase節の処理が可能な場合にはランダムに選択して処理する

	/*
	* 複数のチャネルと複数のゴルーチンにselect文を組み合わせた処理
	 */
	// 各々非同期に処理されるチャネルのデータを適切に処理できていることが確認できる
	// ch12 から ch13 から ch14 へ非同期で値を渡して出力しているだけ

	ch12 := make(chan int)
	ch13 := make(chan int)
	ch14 := make(chan int)

	// ch12から受信した整数を2倍してch13へ送信
	go func() {
		for {
			i := <-ch12
			ch13 <- (i * 2)
		}
	}()

	// ch13から受信した整数を1減産してch14へ送信
	go func() {
		for {
			i := <-ch13
			ch14 <- (i - 1)
		}
	}()

	n := 1
LOOP:
	for {
		select {
		// 整数を増分させつつch12へ送信
		case ch12 <- n:
			n++
		// ch14から受信した整数を出力
		case i := <-ch14:
			fmt.Println("received!! ->", i)
		default:
			if n > 30 {
				break LOOP
			}
		}
	}
}
