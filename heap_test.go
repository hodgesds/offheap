package offheap

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_NewSmall(t *testing.T) {
	s := NewSmall()
	s.s = int8(2)

	if int(reflect.TypeOf(*s).Size()) <= 0 {
		t.Fail("Failed to allocate Small")
	}
}

func Test_NewLarge(t *testing.T) {
	l := NewLarge()
	if int(reflect.TypeOf(*l).Size()) <= 0 {
		t.Fail("Failed to allocate Large")
	}
}

func Benchmark_Small(b *testing.B) {
	var ss *Small
	for i := 0; i < b.N; i++ {
		ss = NewSmall()
		for z := 0; z < 1000; z++ {
			ss = NewSmall()
		}
	}
}

func Benchmark_Large(b *testing.B) {
	var ll *Large
	for i := 0; i < b.N; i++ {
		ll = NewLarge()
		lg := NewLarge()
		ll = lg
		a := NewLarge()
		b := NewLarge()
		c := NewLarge()
		d := NewLarge()
		e := NewLarge()
	}
}
