package main

import "fmt"

func main(){

	var x int = 0

	Label1: for x < 8 {
		if x == 5 {
			x = x + 1;
			goto Label1
		}
		fmt.Printf("value is: %d\n", x);
		x++;
	}
}