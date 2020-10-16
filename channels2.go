package main

import "fmt"

func myfun(c chan string){

	for v:=0;v<4;v++{
		c <- "Geeksforgeeks"
	}
	close(c)
}


func main(){
	c := make(chan string)

	go myfun(c)

	for{
		res,ok := <-c
		if ok == false {
			fmt.Println("channel close ", ok)
			break
		}
		fmt.Println("channel open",res,ok)
	}
}