package main

import (
	"html/template"
	"os"
)

type User struct {
	Name   string
	Age    int
	Height float32
	Address
	Exam []Exam
}

type Address struct {
	City string
}

type Exam struct {
	Subject string
	Score   float32
}

func main() {

	t, err := template.ParseFiles("exercises.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:   "GrayJunZi",
		Age:    25,
		Height: 169.9,
		Address: Address{
			City: "Shanghai",
		},
		Exam: []Exam{
			{
				Subject: "计算机",
				Score:   90.8,
			},
			{
				Subject: "物理",
				Score:   55.5,
			},
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
