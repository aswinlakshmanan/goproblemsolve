package main

import "fmt"

type author struct{
	name string
	branch string
	particles int
	salary int
}

func (a author) show(){
	fmt.Println("Author's name: ",a.name)
	fmt.Println("Branch name: ",a.branch)
	fmt.Println("Published Articles: ",a.particles)
	fmt.Println("Salary: ",a.salary)
}

func main(){
	res := author{
		name : "sona",
		branch : "cse",
		particles : 203,
		salary : 34000,
	}
	res.show()
}
