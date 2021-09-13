package main

import "fmt"

type person struct {
	name    string
	year    int
	city    string
	drivlic string
	contacts
}
type contacts struct {
	mail     string
	phone    string
	telegram string
}

func century(y int) {
	x := (y + 99) / 100
	fmt.Println("century of birth ", x)
}

func (p person) String() string {
	return fmt.Sprintf("%v(%v )",p.name,p.mail)
	
}

func main() {

	var kolya = person{
		name:    "Kolya",
		year:    1997,
		city:    "Minsk",
		drivlic: "abc",
		contacts: contacts{
			mail:     "kol@tut.by",
			phone:    "+3752977777777",
			telegram: "tg_kolyan777",
		},
	}

fmt.Println(kolya)
	century(kolya.year)
	
}

