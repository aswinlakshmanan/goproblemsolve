package main

import "fmt"

func main(){

	nums := []int{1,2,3}
	sum := 0
	for _,num := range nums{
		sum += num
	}
	fmt.Println("sum:",sum)

	for i,num := range nums{
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	
}