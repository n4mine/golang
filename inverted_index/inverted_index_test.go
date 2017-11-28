package main

import (
	"strings"
	"testing"

	"github.com/cespare/xxhash"
)

func Test_InvertedIndex(t *testing.T) {
	iIndex := NewInvertedIndex()
	fIndex := NewForwardIndex()
	if iIndex.Size() != 0 {
		t.Error("want 0")
	}

	if fIndex.Size() != 0 {
		t.Error("want 0")
	}

	tDatas := []struct {
		s string
	}{
		{"__metric__=cpu.idle" + SEP + "__ns__=ns1" + SEP + "host=host1"},
		{"__metric__=cpu.idle.core" + SEP + "__ns__=ns1" + SEP + "host=host1" + SEP + "core=1"},
		{"__metric__=cpu.idle.core" + SEP + "__ns__=ns1" + SEP + "host=host1" + SEP + "core=2"},
		{"__metric__=cpu.idle.core" + SEP + "__ns__=ns1" + SEP + "host=host1" + SEP + "core=3"},
		{"__metric__=cpu.idle.core" + SEP + "__ns__=ns1" + SEP + "host=host1" + SEP + "core=10"},
	}

	for _, data := range tDatas {
		iIndex.AddDoc(xxhash.Sum64String(data.s), data.s)
		fIndex.AddDoc(xxhash.Sum64String(data.s), data.s)
	}

	for docID := range (Intersect(iIndex.Search("__metric__=cpu.idle.core"), (iIndex.Search("host=host1")))).Iter() {
		if !strings.Contains(fIndex.ToCounterString(docID.(uint64)), "__metric__=cpu.idle.core") {
			t.Errorf("want contain __metric__=cpu.idle.core, got %v ", fIndex.ToCounterString(docID.(uint64)))
		}
	}

	for docID := range iIndex.Search("core=10").Iter() {
		if fIndex.ToCounterString(docID.(uint64)) != "__metric__=cpu.idle.core,__ns__=ns1,core=10,host=host1" {
			t.Error("want equal")
		}
	}
}
