package main

import(
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func say(s string){
	for i:=0;i<3;i++{
		fmt.Println(s)
		time.Sleep(time.MilliSecond*100)
	}
	wg.Done()
}

func main(){
	wg.Add(1)
	go say("Hey")
	wg.add(1)
	go say("There")
	wg.Wait()
}