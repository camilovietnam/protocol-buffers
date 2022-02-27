package main

import (
	"fmt"
	abpb "src/gen/gen/addressbookpb"
)

func main() {
	ab := &abpb.AddressBook{}

	person1 := newPerson(1, "Dylan", "dylan.pastor@gmail.com")
	person2 := newPerson(2, "Hector", "lavoe.hector@latino.com")

	// Test the write code
	writePerson(ab, person1)
	writePerson(ab, person2)
	printBook(ab)

	// Test the read code
	foundPerson := readPerson(ab, 1)
	fmt.Println("Search person by ID 1:")
	printPerson(foundPerson)
}

func newPerson(id int32, name, email string) *abpb.Person {
	return &abpb.Person{
		Name:  name,
		Id:    id,
		Email: email,
		Phones: []*abpb.PhoneNumber{
			{Number: "972-234567", Type: abpb.PhoneType_HOME},
			{Number: "633-112233", Type: abpb.PhoneType_MOBILE},
		},
	}
}

func writePerson(ab *abpb.AddressBook, person *abpb.Person) {
	ab.People = append(ab.People, person)
}

func readPerson(ab *abpb.AddressBook, id int32) *abpb.Person {
	for _, p := range ab.People {
		if p.GetId() == id {
			return p
		}
	}

	return nil
}

func printBook(ab *abpb.AddressBook) {
	for _, p := range ab.People {
		printPerson(p)
	}
}

func printPerson(p *abpb.Person) {
	url := fmt.Sprintf("[%d] %s. Contact: %s", p.GetId(), p.GetName(), p.GetEmail())
	fmt.Println(url)
}
