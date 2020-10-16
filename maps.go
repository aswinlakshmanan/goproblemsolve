package main
import "fmt"

func main(){
	statePopulations := map[string]int{
		"california" : 257399,
		"Texas": 986246,
		"Florida" : 287634,
		"Pennsulvania" : 872934,
		"Ohio" : 826384,
	}
	for k,v := range statePopulations{
		fmt.Println(k,v)
	}
}