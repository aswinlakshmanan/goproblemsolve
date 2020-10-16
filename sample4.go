package main

import "fmt"

func  main(){

	array := [4]int{1,2,3,4}
	
	array1 := array

	fmt.Println("array1 = %T", array1)

	for i:=0;i<=3;i++{
		fmt.Println(array1[i])
	}
}