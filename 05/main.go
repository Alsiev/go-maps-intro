package main

import "fmt"

func main() {
	//var m map[string]int
	m := make(map[string]int)

	//В Go неинициализированная карта (map) имеет значение nil, и попытка записи в такую карту вызовет панику (runtime panic).

	m["visits"] = 1

	fmt.Println(m)

}
