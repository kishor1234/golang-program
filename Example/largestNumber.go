package main

import "fmt"

func main(){
	n1 := 100
	n2 := 50
	n3 := 30
	if(n1>=n2){
		if(n1>=n3){
			fmt.Println("%d is Largest Number n1",n1)
		}else{
			fmt.Println("%d is Largest Number n2",n3)
		}
	}else{
		if(n2>=n3){
			fmt.Println("%d is Largest Number n2",n2)
		}else{
			fmt.Println("%d is Largest Number n3",n3)
		}
	}
}