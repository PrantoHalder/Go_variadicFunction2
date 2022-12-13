package main

import "fmt"

func variadicFunction4 () {
	
	names := [] string {"pranto","halder","shovon"}
    repetation(names...)
	
}
func repetation( names ...string) {
        
		 for _,value := range names {
			fmt.Println(value)
		 }
		
}