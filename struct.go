package main

import "fmt"

func main(){

	type address struct {
		name string
		street string
		city string
		pincode int
	}

	a1 := address{"Aswin Lakshmanan","Indira Gandhi Street","Erode",638011}

	fmt.Println(a1)
}