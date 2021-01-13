package main

import (
	"fmt"
)

func noNameStruct(s struct{ X, Y int }) {
	fmt.Println(s)
}
