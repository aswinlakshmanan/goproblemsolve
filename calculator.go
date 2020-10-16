package main

import "fmt"

func add(n1,n2 int)int{
	var ans int
	ans=n1 + n2
	return ans
}
func sub(n1,n2 int)int{
	var answer int
	answer=n1 - n2
	return answer
}
func mul(n1,n2 int)int{
	var an int
	an=n1 * n2
	return an
}

func  main(){
	
	var num int
	var a,b int
	fmt.Scanln(&num)

switch num{
case 1:
	fmt.Println("First Number: ")
	fmt.Scan(&a)
	fmt.Println("Second Number: ")
	fmt.Scan(&b)

	fmt.Println("addition: %d", add(a,b))
case 2:
	fmt.Println("First Number: ")
	fmt.Scan(&a)
	fmt.Println("Second Number: ")
	fmt.Scan(&b)
	fmt.Println("Subtraction: %d", sub(a,b))
case 3:
	fmt.Println("First Number: ")
	fmt.Scan(&a)
	fmt.Println("Second Number: ")
	fmt.Scan(&b)
	fmt.Println("multiplication: %d", mul(a,b))
default:
	fmt.Println("INVALID")
}


}