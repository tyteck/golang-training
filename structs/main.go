package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	postalAddress
}

type postalAddress struct {
	streetname string
	postalCode string
	city       string
}

func main() {
	johnAddress := postalAddress{"451 chemin de la prairie", "06790", "Aspremont"}

	johnConnor := person{
		firstname:     "john",
		lastname:      "connor",
		postalAddress: johnAddress,
	}
	johnConnor.print()

	var alex person

	alex.firstname = "Alex"
	alex.lastname = "Anderson"
	alex.updateLastname("Andersen")

	alex.print()
}

func (p *person) updateLastname(lastname string) {
	p.lastname = lastname
}

func (p person) print() {
	fmt.Println(p)
}
