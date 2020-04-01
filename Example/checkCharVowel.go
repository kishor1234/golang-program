package main
import (
	"fmt"
	"strings"
)

//Package strings implements simple functions to manipulate UTF-8 encoded strings.

//For information about UTF-8 strings in Go, see https://blog.golang.org/strings.

func main(){
	c :="kishor"
	
	if(strings.ToUpper(c)==c){//Should print true beacuse our input is in Upper Case
		fmt.Println("%d is Uppercase ",c)
	}else{
		fmt.Println("%d is LowerCase ",c)
	}

}