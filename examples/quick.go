package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"

	"github.com/hodgesds/offheap"
)

func main() {
	i, _, err := offheap.IntSlice(10000000)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	for {
		for z := 0; z < 10000000; z++ {
			v := rand.Intn(100)
			i[z] = v
		}
		runtime.GC()
	}
}
