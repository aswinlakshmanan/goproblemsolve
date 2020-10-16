package main

import ( "fmt"\
"math/rand"
"time"
)


func caluculatevalue(c chan int){
 
	v := rand.Intn(20)
	fmt.Println("Random Values are : ",v)
	time.Sleep(1000 * time.Millisecond)
	c <- v



}

func main(){
	fmt.Println("Welcome to go routine channels tutorials")

	valuechannel := make(chan int)
	defer close(valuechannel)

	go calculatevalue(valuechannel)

	go calculatevalue(valuechannel)

	
	values := <-valuechannel

	fmt.Println(values)



	

}