package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToWords(s0 string) []string {
	p1 := regexp.MustCompile(`([A-Z])`)
	s1 := p1.ReplaceAllString(s0, `_${1}`)
	words := strings.Split(s1, "_")
	if words[0] == "" {
		words = words[1:]
	}
	return words
}

func ToCamel(words []string) string {
	ws := make([]string, len(words))
	copy(ws, words)

	for i, w := range ws {
		if i == 0 {
			ws[i] = strings.ToLower(w)
		} else {
			ws[i] = strings.ToUpper(w[0:1]) + w[1:]
		}
	}
	return strings.Join(ws, "")
}

func ToSnake(words []string) string {
	ws := make([]string, len(words))
	copy(ws, words)

	for i, w := range ws {
		ws[i] = strings.ToLower(w)
	}
	return strings.Join(ws, "_")
}

func ToPascal(words []string) string {
	ws := make([]string, len(words))
	copy(ws, words)

	for i, w := range ws {
		ws[i] = strings.ToUpper(w[0:1]) + w[1:]
	}
	return strings.Join(ws, "")
}

func main() {
	var code int
	var name string

	_, _ = fmt.Scan(&code, &name)
	words := ToWords(name)

	fmt.Println(ToCamel(words))
	fmt.Println(ToSnake(words))
	fmt.Println(ToPascal(words))
}
