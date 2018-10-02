package main

import (
	"fmt"

	"github.com/hodgesds/offheap"
)

func main() {
	var l *offheap.Large
	b := make([]*offheap.Small, 10000000, 10000000)
	for {
		l = offheap.NewLarge()
		for z := 0; z < 10000000; z++ {
			b[z] = offheap.NewSmall()
		}
	}
	fmt.Printf("%+T\n", l)
}
