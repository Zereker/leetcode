package main

import (
	"crypto/md5"
	"encoding/base64"
	"reflect"
	"unsafe"
)

type BloomFilter struct {
	size  int
	k     int
	slice []byte
}

func New(size, k int) *BloomFilter {
	return &BloomFilter{
		size:  size,
		k:     k,
		slice: make([]byte, 0, size),
	}
}

// TODO hash 算法有问题
func hash(i interface{}) string {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		if !v.CanAddr() {
			return ""
		}

		v = v.Addr()
	}

	size := unsafe.Sizeof(v.Interface())
	b := (*[1 << 10]uint8)(unsafe.Pointer(v.Pointer()))[:size:size]

	h := md5.New()
	return base64.StdEncoding.EncodeToString(h.Sum(b))
}

func (b *BloomFilter) hash(x interface{}) int {
	s := hash(x)

	var result int32
	for _, v := range s {
		result += v
	}

	return int(result)
}

func (b *BloomFilter) Add(x interface{}) {
	for i := 0; i < b.k; i++ {
		v := b.hash(x) % b.size
		b.slice[v] = 1
	}
}

func (b *BloomFilter) Have(x interface{}) bool {
	for i := 0; i < b.k; i++ {
		v := b.hash(x) % b.size
		if data := b.slice[v]; data == 0 {
			return false
		}
	}

	return true
}

func main() {

}
