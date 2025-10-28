package main

import "fmt"

func main() {

	weekend := map[string]int{"Mon": 2, "Tue": 3, "Wen": 1, "Thu": 1, "Fri": 3, "Saturday": 3}

	fmt.Println(weekend["Tue"])
}
