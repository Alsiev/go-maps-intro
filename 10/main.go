package main

import "fmt"

func Add(phones map[string]string, name, number string) {
	phones[name] = name
	phones[number] = number

}

func Get(phones map[string]string, name string) (string, bool) {
	a, ok := phones[name]

	if ok {
		return a, ok
	}

	return "Абонент не найден", ok

}

func Remove(phones map[string]string, name string) {

	delete(phones, "Yusup")

	phones[name] = name

}

func main() {
	//Опиши карту: phones := map[string]string{} где ключ — имя, значение — номер.

	phones := map[string]string{}

	Add(phones, "Yusup", "89380160290")

	fmt.Println(Get(phones, "Yusup"))

	Remove(phones, "Yunus")

	fmt.Println(phones)

}
