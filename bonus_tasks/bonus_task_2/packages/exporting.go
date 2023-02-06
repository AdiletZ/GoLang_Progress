package packages

import "fmt"

type Footballers struct {
	name    string
	surname string
	age     int
	from    string
}

type Clubs struct {
	name    string
	created int
	country string
}

func (f Footballers) getFootballerInfo() (string, string, int, string) {
	return f.name, f.surname, f.age, f.from
}

func (c Clubs) getClubInfo() (string, int, string) {
	return c.name, c.created, c.country
}

func Funct() {
	f1 := Footballers{"Ronaldo", "Cristiano", 37, "Portugal"}
	c1 := Clubs{"Al-Nasr", 1972, "Saudi Arabia"}
	fmt.Printf("%s %s is %d years old. He is from %s.\nHe plays in %s football club which was created in %d in %s.", f1.surname, f1.name, f1.age, f1.from, c1.name, c1.created, c1.country)
}
