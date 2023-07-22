package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 20
	return &p
}

func main() {
	fmt.Println(person{"Blob", 20})

	p1 := &person{name: "Ann", age: 10}
	fmt.Println(p1)

	p2 := p1
	p2.age = 22
	fmt.Println("p2 ", p2)
	fmt.Println("p1 ", p1)

	fmt.Println(newPerson("Joe"))

	dog := struct {
		name string
	}{
		name: "wangcai",
	}

	fmt.Println(dog)

}
