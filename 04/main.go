package main

import "fmt"

func main() {
	sight := map[string]string{"Урус-мартан": "Автомобильный рынок", "Закан-юрт": "Природа", "Шали": "Мечеть", "Сержень-юрт": "Река"}

	for k, v := range sight {
		fmt.Println(k, v)
	}

}

//потомучто не сортировоный тип находятся в разброс
