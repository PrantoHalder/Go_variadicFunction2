package main

import "fmt"

func variadicFunction3 () {
  
  num := [] int {1,2,3,4,5,6,7,8,9,10}
  name , res := add(num...)
  
  fmt.Println("the addition of the values ",name,res)
}

func add (num ...int) (string ,int) {
  res := 0
  for _,value := range num{
	res += value 
  }
  return "pranto" , res
}