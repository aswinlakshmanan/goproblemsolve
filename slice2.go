package main
import (
	"fmt"
	"sort"
)

func main(){

	slc1 := []string{"python","java","c","go","ruby"}
	slc2 := []int{45,54,65,1,20}

	fmt.Println(slc1)
	fmt.Println(slc2)

	sort.Strings(slc1)
	sort.Ints(slc2)

	fmt.Println("Sorted slice is: ",slc1)
	fmt.Println("Sorted slice is: ",slc2 )
}