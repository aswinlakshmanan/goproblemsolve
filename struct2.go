package main

import "fmt"

type student struct{
	id int
	name string
	branch string

}

func main(){
	emp := &student{206,"aswin","cse"}
	fmt.Println("id is: " ,(*emp).id)
	fmt.Println("name is: ",(*emp).name)
	fmt.Println("brach is: ",(*emp).branch)

}