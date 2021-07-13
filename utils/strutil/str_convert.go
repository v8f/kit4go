package strutil

import (
	"reflect"
	"unsafe"
)

// the Data field is not sufficient to guarantee the data it references will not be garbage collected,
// so programs must keep a separate,
// correctly typed pointer to the underlying data.
func String2BytesUnRecommend(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sb := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&sb))
}

func String2Bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var b []byte
	pBytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pBytes.Data = stringHeader.Data
	pBytes.Len = stringHeader.Len
	pBytes.Cap = stringHeader.Len
	return b
}
