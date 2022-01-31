package main

type person struct {
	name, surname string
}

func (p *person) String() string {
	return p.surname + ", " + p.name
}

var people = []*person{
	{"Hans", "Emil"},
	{"Max", "Mustermann"},
	{"Roma", "Tisch"},
}
