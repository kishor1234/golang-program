package main

import "fmt"

func main(){
	num := 25
	if(num%2 == 0)	{
		fmt.Println("%d Number is Even",num)
	}else{
		fmt.Println("%d Number is Odd",num)
	}
}