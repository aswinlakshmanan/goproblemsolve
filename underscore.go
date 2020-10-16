package main

import"fmt"

func call(n1,n2 int)(int,int){
	return n1 * n2, n1/n2
}

func main(){

	mul, _ := call(105,7)
	fmt.Println("Multiplication is: ",mul)
	

}