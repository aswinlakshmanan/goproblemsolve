package main

import "fmt"

type employee struct {
	firstname,lastname string
	age, salary int

}

func main(){
	emp8 := &employee{"sam","anderson",5,6000}
	fmt.Println("First name: ",(*emp8).firstname)
	fmt.Println("Age: ",(*emp8).age)

}
