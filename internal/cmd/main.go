package main

import (
	"fmt"
	"dz1/task_1"
)

func main(){
	a, b, err := task1.FilterCommomDigits(1235, 55702)
	fmt.Printf("FilterCommomDigits(1235, 55702) = (%d, %d, %v)\n", a, b, err)
}
