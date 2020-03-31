package main
import "fmt"

type car struct{
	car_name string
	car_speed float64
	car_no int32
}

func main(){
	a_car := car{"Java",20.0,12}
	fmt.Printf(a_car.car_name)
}

