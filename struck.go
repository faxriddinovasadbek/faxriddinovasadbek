package main

import "fmt"

type Person struct{
	FirstName string
	LastName string
	Age int
	IsOnline bool
	Salary float32
	PhoneNumber []string
}

func (a *Person) FullName() string {
	return a.FirstName + " " + a.LastName
}

func (a *Person) SetSalary(sum float32) {
	a.Salary = sum
}


func main(){

	a := Person{
		FirstName: "Tom",
		LastName: "Jerry",
		Salary: 1000,
		Age: 19,
		IsOnline: false,
		PhoneNumber: []string{"99 478-60-66", "99 432-43-43"},
	}

	fmt.Println(a.Salary)
	a.SetSalary(700)
	fmt.Println(a.Salary)

}
