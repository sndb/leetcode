package main

import (
	"fmt"
	"strings"
)

type tree struct {
	child [26]*tree
	end   bool
}

func (t *tree) add(word string) {
	for _, w := range word {
		c := w - 'a'
		if t.child[c] == nil {
			t.child[c] = &tree{}
		}
		t = t.child[c]
	}
	t.end = true
}

func (t *tree) root(word string) int {
	n := 0
	for _, w := range word {
		if t == nil {
			return len(word)
		}
		if t.end {
			break
		}
		t = t.child[w-'a']
		n++
	}
	return n
}

func replaceWords(dictionary []string, sentence string) string {
	t := &tree{}
	for _, w := range dictionary {
		t.add(w)
	}

	a := strings.Fields(sentence)
	for i, w := range a {
		a[i] = w[:t.root(w)]
	}
	return strings.Join(a, " ")
}

func main() {
	dictionary := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dictionary, sentence))
}
