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

func (p person) String() string {
	return fmt.Sprintf("%v(%v )",p.name,p.city)
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

	var vasya = person{
		name:    "Vasya",
		year:    1991,
		city:    "shchuchynshchyna",
		drivlic: "abcde",
		contacts: contacts{
			mail:     "vas@tut.by",
			phone:    "+375294444444",
			telegram: "tg_uuuuu444",
		},
	}

	var olga = person{
		name:    "Olga",
		year:    1999,
		city:    "Moskow",
		drivlic: "b",
		contacts: contacts{
			mail:     "olya@tut.by",
			phone:    "+375292929292",
			telegram: "tg_olya29",
		},
	}

	
	var users []person =[]person{kolya,vasya,olga}
	for index, value := range users{
    fmt.Println(index," | ",value)
}

}


