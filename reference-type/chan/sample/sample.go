package main

import (
	"fmt"
	"time"
)

// ゴルーチンを3つ起動して非同期に実行されるゴルーチンとチャネルによるデータ共有が機能していることを検証

func receiver(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			// 受信できなくなったら終了
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is done.")
}

func main() {
	ch := make(chan int, 20)

	go receiver("1st goroutine", ch)
	go receiver("2st goroutine", ch)
	go receiver("3st goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	// ゴルーチンの完了を3秒待つ
	time.Sleep(3 * time.Second)
}
