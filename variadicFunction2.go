package main 

import "fmt"

func Variadic() {
    num := [] int {1,2,3,4,5,6,7,8,9,10}
 fmt.Println("print the sum of numbers ",sum(num...))
}

func sum (nums ...int) (int){
	res := 0
	for _,value := range nums{
		res = res + value
	}
	return  res
}