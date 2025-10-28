package main

import "fmt"

type Student struct {
	Name   string
	Active bool
}

func main() {

	people := map[string]Student{"15": Student{Name: "Yusup", Active: true}, "29": Student{Name: "Yunus", Active: false}}

	user := people["15"]

	fmt.Println("До изминения ", user, "\n")

	user.Active = false
	people["15"] = user

	fmt.Println("После изминеня", user, "\n")

	fmt.Println("Всякарта после изминения", people)

}
