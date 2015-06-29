// Benchmark and sample data lifted from go source,
// We'll say it is all copyright 2011 The Go Authors.
// We'll claim all rights are reserved by them and and
// use of this source code is governed by their BSD-style
// license that can be found in the LICENSE file at
// https://github.com/golang/go

package test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/lloyd/goj"
)

var codeJSON []byte

func codeInit() {
	f, err := os.Open("./code.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	codeJSON = data
}

func BenchmarkGojScanning(b *testing.B) {
	if codeJSON == nil {
		b.StopTimer()
		codeInit()
		b.StartTimer()
	}
	parser := goj.NewParser()
	for i := 0; i < b.N; i++ {
		err := parser.Parse([]byte(codeJSON), func(t goj.Type, k []byte, v []byte) bool {
			return true
		})
		if err != nil {
			b.Fatal("Scanning:", err)
		}
	}
	b.SetBytes(int64(len(codeJSON)))
}
