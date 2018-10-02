package main

import (
	"math/rand"
)

func main() {
	i := make([]*int, 10000000, 10000000)
	for {
		for z := 0; z < 10000000; z++ {
			v := rand.Intn(100)
			i[z] = &v
		}
	}
}
