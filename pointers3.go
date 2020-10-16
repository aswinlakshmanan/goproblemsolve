 package main

 import "fmt"

 func main(){

	arr := [3]int{1,2,3}
	a := &arr[0]
	b := &arr[1]
	c := &arr[2]

	fmt.Println("%v,%p,%p",a,b,c)
 }