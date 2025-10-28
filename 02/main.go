package main

import "fmt"

func getVisits(m map[string]int, day string) (int, bool) {

	d, ok := m[day]

	return d, ok
}

func main() {
	weekend := map[string]int{"Mon": 2, "Tue": 3, "Wen": 1, "Thu": 1, "Fri": 3, "Saturday": 3}

	a, b := getVisits(weekend, "Mon1")

	if b == true {

		fmt.Println("Значение найдено", a, b)
	} else {
		fmt.Println("Значение не найдено")
	}

}
