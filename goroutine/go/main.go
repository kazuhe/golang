package main

import (
	"fmt"
	"runtime"
)

func sub() {
	for i := 0; i < 100; i++ {
		fmt.Println("sub loop")
	}
}

func printRuntime() {
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())             // ランタイムが動作するCPUの使用できるコア数
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine()) // ランタイム上で動作しているゴルーチンの数
	fmt.Printf("Version: %s\n", runtime.Version())
}

func main() {
	/*
	* go
	 */
	// go文は並行処理を司る特別な機能
	// go文によってsub関数を呼び出し並行してループ処理がされる
	// スレッドよりも小さい処理単位である「goroutine（ゴルーチン）」が並行して動作する
	// go文はこのゴルーチンを新たに生成して新しい処理の流れををランタイムに追加する機能
	go sub()
	for i := 0; i < 100; i++ {
		fmt.Println("main loop")
	}

	printRuntime()
}
