package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {

	c := strings.Fields(s)

	a := make(map[string]int)

	for i, v := range c {
		a[v] = i
	}

	return a

}

func main() {
	fmt.Println(WordCount("go go gopher"))
	fmt.Println(WordCount("интукод школа программирования интокод"))
}
