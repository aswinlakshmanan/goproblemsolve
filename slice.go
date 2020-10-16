package main

import "fmt"

func main(){

	var arr = [3]string{"geeks","for","geeks"}

	arr1 := arr[0:]
	fmt.Println(arr1)

	slice := []int{1,2,3,4,5}
	fmt.Println(slice[:])
	fmt.Println(slice[0:])
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

}

