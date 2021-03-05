package main

import (
	"fmt"
)

func main() {
	x := 12

	if x >= 5 && x < 10 {
		fmt.Println("5以上10未満")
	} else {
		fmt.Println("それ以外")
	}
}
