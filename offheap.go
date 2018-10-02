package offheap

import (
	"reflect"
	"syscall"
	"unsafe"
)

var b byte

// Mmap is a helper method for allocating slices via mmap. It returns a uinptr
// which then can be reflected to a slice.
func Mmap(count, size int) (uintptr, int, error) {
	var pages, fd int

	pageSize := syscall.Getpagesize()
	pages = (size * count) / pageSize
	if pages%pageSize != 0 {
		pages = pageSize*pages + pageSize
	}

	fd = -1

	ptr, uFd, errno := syscall.Syscall6(
		syscall.SYS_MMAP,
		uintptr(0),
		uintptr(pages),
		syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC,
		syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE|syscall.MAP_POPULATE,
		uintptr(fd),
		uintptr(0),
	)

	var err error
	if errno != 0 {
		err = errno
		return ptr, -1, err
	}

	return ptr, int(uFd), err
}

// Bytes returns a slice of bytes that are allocated off heap using mmap as
// well as an associated file descriptor.
func Bytes(count int) ([]byte, int, error) {
	var bs []byte
	byteSize := int(reflect.TypeOf(b).Size())
	ptr, fd, err := Mmap(count, byteSize)
	if err != nil {
		return bs, fd, err
	}

	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}

// IntSlice returns a slice of ints that are allocated off heap using mmap as
// well as an associated file descriptor.
func IntSlice(count int) ([]int, int, error) {
	var i []int
	byteSize := int(reflect.TypeOf(count).Size())
	ptr, fd, err := Mmap(count, byteSize)
	if err != nil {
		return i, fd, err
	}

	return *(*[]int)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}

// IntPSlice returns a slice of ints that are allocated off heap using mmap as
// well as an associated file descriptor.
func IntPSlice(count int) ([]*int, int, error) {
	var i []*int
	ptr, fd, err := Mmap(count, 8)
	if err != nil {
		return i, fd, err
	}

	return *(*[]*int)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}

// Int32Slice returns a slice of int32s that are allocated off heap using mmap as
// well as an associated file descriptor.
func Int32Slice(count int) ([]int32, int, error) {
	var i []int32
	ptr, fd, err := Mmap(count, 32)
	if err != nil {
		return i, fd, err
	}

	return *(*[]int32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}

// Uint32Slice returns a slice of uint32s that are allocated off heap using
// mmap as well as an associated file descriptor.
func Uint32Slice(count int) ([]uint32, int, error) {
	var i []uint32
	ptr, fd, err := Mmap(count, 32)
	if err != nil {
		return i, fd, err
	}

	return *(*[]uint32)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}

// Int64Slice returns a slice of int64s that are allocated off heap using mmap as
// well as an associated file descriptor.
func Int64Slice(count int) ([]int64, int, error) {
	var i []int64
	ptr, fd, err := Mmap(count, 64)
	if err != nil {
		return i, fd, err
	}

	return *(*[]int64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}

// Uint64Slice returns a slice of uint64s that are allocated off heap using
// mmap as well as an associated file descriptor.
func Uint64Slice(count int) ([]uint64, int, error) {
	var i []uint64
	ptr, fd, err := Mmap(count, 64)
	if err != nil {
		return i, fd, err
	}

	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{
		Data: ptr,
		Len:  count,
		Cap:  count,
	})), fd, nil
}
