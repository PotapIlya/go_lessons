package main

import "fmt"

type Person struct {
	name string
	age  int
}

type User struct {
	Person
	city     string
	Children []Person
}

func (u *User) setName(name string) {
	u.name = name
}

func (u *User) setCity(city string) {
	u.city = city
}

func (u *User) addChildren(user Person) {
	u.Children = append(u.Children, user)
}

func (u *User) countChildren() int {
	return len(u.Children)
}

func main() {

	user := User{
		Person: Person{"Bob", 21},
		city:   "San Jose",
		Children: []Person{
			Person{"Bob1", 21},
		},
	}

	user.setName("Potap")
	user.setCity("London")

	user.addChildren(Person{
		name: "Bob2",
		age:  22,
	})

	fmt.Println(
		user,
	)
	//fmt.Println(user)

}
