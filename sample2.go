package main

import "fmt"

func area(length, width int)int{
	ar := length * width
	return ar
}

func  main(){

	fmt.Printf("area of the rectangle is: %d ", area(2,3))
}