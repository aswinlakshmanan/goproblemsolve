package main

import "fmt"

func main(){


	slice := [6]int{55,66,77,88,99,22}
	slc := slice[0:4]

	fmt.Println("Original array: ",slice)
	fmt.Println("Original slice: ", slc)

	slc[0] = 100
	slc[1] = 1000
	slc[2] = 1000

	fmt.Println("\nnew array: ", slice)
	fmt.Println("New slice is: ",slc)
}