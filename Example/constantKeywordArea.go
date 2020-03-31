package main
import "fmt"
func main(){
	const LENGHT=10 //Note that it is a good programming practice to define constants in CAPITALS.
	const HEIGHT=5
	var area int

	area=LENGHT*HEIGHT

	fmt.Printf("Area = %d\n",area)
}