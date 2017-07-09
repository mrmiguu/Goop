package goop

import (
	"strings"

	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

func alert(v interface{}) {
	js.Global.Call("alert", v)
}

func panicjs(v interface{}) {
	defer alert("panic: " + strings.ToLower(fmt.Sprint(v)))
	panic(v)
}

func wildcards(s string, words ...string) string {
	wildcard := "*"
	if strings.Count(s, wildcard) != len(words) {
		panicjs("wildcard count does not match words")
	}
	for _, word := range words {
		s = strings.Replace(s, wildcard, word, 1)
	}
	return s
}
