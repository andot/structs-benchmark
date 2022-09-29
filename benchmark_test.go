package structsbenchmark_test

import (
	"reflect"
	"testing"

	astructs "github.com/andot/structs"
	fstructs "github.com/fatih/structs"
)

func TestData(t *testing.T) {
	data := NewResponse()
	a := astructs.Map(data)
	b := fstructs.Map(data)
	if !reflect.DeepEqual(a, b) {
		t.Error("wrong")
	}
}

func BenchmarkAndotStructs(b *testing.B) {
	data := NewResponse()
	for i := 0; i < b.N; i++ {
		_ = astructs.Map(data)
	}
}

func BenchmarkFatihStructs(b *testing.B) {
	data := NewResponse()
	for i := 0; i < b.N; i++ {
		_ = fstructs.Map(data)
	}
}
