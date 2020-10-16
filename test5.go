package main

import "fmt"

func main(){
	var value int = 2

	switch{
	case value == 1:
		fmt.Println("Hello")
	case value == 2:
		fmt.Println("Bonjour")
	case value == 3:
		fmt.Println("Namastey")
	case value == 4:
		fmt.Println("Invalid")
	}


}