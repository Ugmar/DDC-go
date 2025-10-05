package main

import (
	"fmt"
	// "dz1/task_1"
	// "dz1/task_2"
	"dz1/task_3"
)

func main(){
	slice := []int{1, 2, 3}

		// When
	err := task3.ScaleSlice(&slice, 3)
	fmt.Println(slice, err)
}
