package offheap

import "testing"

func Test_Bytes(t *testing.T) {
	s, _, err := Bytes(10000000)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(s); i++ {
		s[i] = byte(1)
	}
}

func Test_IntSlice(t *testing.T) {
	is, _, err := IntSlice(10000000)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(is); i++ {
		is[i] = i
	}
}

func Test_Int32Slice(t *testing.T) {
	is, _, err := Int32Slice(10000000)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(is); i++ {
		is[i] = int32(i)
	}
}

func Test_Uint32Slice(t *testing.T) {
	is, _, err := Uint32Slice(10000000)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(is); i++ {
		is[i] = uint32(i)
	}
}

func Test_Int64Slice(t *testing.T) {
	is, _, err := Int64Slice(10000000)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(is); i++ {
		is[i] = int64(i)
	}
}

func Test_Uint64Slice(t *testing.T) {
	is, _, err := Uint64Slice(1000000)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(is); i++ {
		is[i] = uint64(i)
	}
}

func Benchmark_Int64Slice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int64Slice(1)
	}
}
