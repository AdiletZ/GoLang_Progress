package main

import "fmt"

// part-1
type Person struct {
	name        string
	phoneNumber int
	location    string
}

//func (p Person) get_per() {
//	fmt.Println("Name: ", p.name)
//	fmt.Println("Phone number: ", p.phoneNumber)
//	fmt.Println("Location: ", p.location)
//}
//
//func (p *Person) set_per() {
//	fmt.Println("Name: ", &p.name)
//	fmt.Println("Phone number: ", &p.phoneNumber)
//	fmt.Println("Location: ", &p.location)
//}

// part-2
type Person2 []string

func (s *Person2) IsEmpty() bool {
	return len(*s) == 0
}
func (s *Person2) Push(str string) {
	*s = append(*s, str)
}
func (s *Person2) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {

	// part-1
	//per := Person{"Adilet", 87473433170, "Kaskelen"}
	//per.get_per()
	//
	//per2 := Person{"Da3y", 87770541656, "Almaty"}
	//per2.set_per()
	// part-2
	var stack Person2

	stack.Push("Data Structures")
	stack.Push("And")
	stack.Push("Algorithms")

	for !stack.IsEmpty() {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
