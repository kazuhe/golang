package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", plus(1, 2))
	hello()
	q, r := div(19, 7)
	fmt.Printf("q=%v r=%v\n", q, r)
}
