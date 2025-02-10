package main

import "fmt"

func main(){
	
	for i := 0; i < 100; i++ {
		if(i % 3 == 0){
			fmt.Printf("Number: %d - Pin!\n", i)
		}else if(i % 5 == 0){
			fmt.Printf("Number: %d - Pan!\n", i)
		}
	}
}