package main

import "fmt"

func main(){

	s := []int{1,2,3}
	k,v := range s{
		fmt.Println(k,v)
	}
}