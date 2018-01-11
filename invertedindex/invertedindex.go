package invertedindex

import (
	"sort"
	"strings"

	"github.com/deckarep/golang-set"
)

const SEP = "\u2318"

// k-v to docID
type InvertedIndex map[string]mapset.Set

func NewInvertedIndex() *InvertedIndex {
	i := make(InvertedIndex)
	return &i
}

func (x *InvertedIndex) Size() int {
	return len(*x)
}

func (x *InvertedIndex) AddDoc(docID uint64, doc string) {
	for _, word := range tokenizer(doc) {
		_, exist := (*x)[word]
		if !exist {
			(*x)[word] = mapset.NewSet()
		}
		(*x)[word].Add(docID)
	}
}

func (x *InvertedIndex) Search(query string) mapset.Set {
	if ref, exist := (*x)[query]; exist {
		return ref
	}

	return nil
}

func Intersect(x, y mapset.Set) mapset.Set {
	if x == nil || y == nil {
		return mapset.NewSet()
	}
	return x.Intersect(y)
}

// docID to counter string
type ForwardIndex map[uint64]string

func NewForwardIndex() *ForwardIndex {
	f := make(ForwardIndex)
	return &f
}

func (x *ForwardIndex) Size() int {
	return len(*x)
}

func (x *ForwardIndex) AddDoc(docID uint64, doc string) {
	temp := tokenizer(doc)
	sort.Strings(temp)
	(*x)[docID] = strings.Join(temp, ",")
}

func (x *ForwardIndex) ToCounterString(docID uint64) string {
	if ref, exist := (*x)[docID]; exist {
		return ref
	}
	return ""
}

func tokenizer(doc string) []string {
	return strings.Split(doc, SEP)
}
