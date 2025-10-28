package main

import "fmt"

func main() {

	weekend := map[string]int{"Mon": 2, "Tue": 3, "Wen": 1, "Thu": 1, "Fri": 3, "Saturday": 3}

	fmt.Println("\n", "До удаления", weekend, len(weekend), "\n")
	delete(weekend, "Tue")
	fmt.Println("После удаления", weekend, len(weekend), "\n")
	delete(weekend, "Holiday")
	fmt.Println("После удаления не существуещего слова", weekend, len(weekend), "\n")
}
