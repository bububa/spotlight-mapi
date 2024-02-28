package util

import (
	"bytes"
	"net/url"
	"strings"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func NewBufferPool() *bytes.Buffer {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func ReleaseBufferPool(buf *bytes.Buffer) {
	bufferPool.Put(buf)
}

var stringsBuilder = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func NewStringsBuilder() *strings.Builder {
	b := stringsBuilder.Get().(*strings.Builder)
	b.Reset()
	return b
}

func ReleaseStringsBuilder(b *strings.Builder) {
	stringsBuilder.Put(b)
}

var urlValuesPool = sync.Pool{
	New: func() any {
		return make(url.Values)
	},
}

func NewUrlValues() url.Values {
	vals := urlValuesPool.Get().(url.Values)
	for k := range vals {
		vals.Del(k)
	}
	return vals
}

func ReleaseUrlValues(vals url.Values) {
	urlValuesPool.Put(vals)
}
