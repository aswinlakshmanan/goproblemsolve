package main

import(
	"fmt"
	"strings"
	"sort"
)

func main(){

	arr := []int{32,654,12,1,34,56,8,1,9,23,5}
	fmt.Println(arr)
	sort.Ints(arr)
	fmt.Println("Sorted array is: ", arr)

	fmt.Println("Index value of aswin in aswinlakshmanan is : ", strings.Index("aswinlakshmanan","w"))

}