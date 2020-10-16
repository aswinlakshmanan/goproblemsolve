package main

import "fmt"

type author struct {
	name string
	branch string
	year int
}
type hr struct {
	details author
}

func main(){

	result := hr{
		details : author{"aswin","cse",2013},
	}

	fmt.Println("\n details of the author: ")
	fmt.Println(result)

}