package main

import "fmt"

type student struct{
	name string
	id int
	branch string
}

type teacher struct {
	name string
	id int
	department string
	details student
}

func main(){
	result := teacher{
		name : "Aswin",
		id : 207,
		branch : "cse",
		details: student{"lakshmanan",206,"cse"},

	}
	fmt.Printn("details of the teacher")
	fmt.Println("Teacher's name: ",result.name)
	fmt.Println("teacher's id: ",result.id)
	fmt.Println("Teacher's branch ",result.branch)
	
	fmt.Println("\nDetals of Student")
	fmt.Println("Student's name: ",result.details.name)
	fmt.Println("Student's id: ",result.details.id)
	fmt.Println("Student's department: ",result.details.department)
	
}