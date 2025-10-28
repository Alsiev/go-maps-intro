package main

import "fmt"

func clone(src map[string]int) map[string]int {
	a := make(map[string]int)
	for i, v := range src {
		a[i] = v
	}
	return a
}
func main() {

	a := map[string]int{"Mon": 10}

	b := a

	fmt.Println("функция", clone(b))
	b["Mon"] = 97 //копируете указатель на ту же самую область памяти. Поэтому любые изменения, сделанные через b, будут видны через a, и наоборот.

	fmt.Println("оригинал:", a, "Клон", b)

}
