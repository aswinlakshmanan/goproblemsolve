package main

import "fmt"

func area(length, width int)int{
	Ar := length * width
	return Ar

}


func main(){
	fmt.Printf("Area of rectangle: %d", area(12,10))
}