package main

import (
	"fmt"
	"dz1/internal/task_1"
	"dz1/internal/task_2"
	"dz1/internal/task_3"
)

func main(){

	slice := []int{1, 2, 3}
	err := task3.ScaleSlice(&slice, 3)
	fmt.Println(slice, err)

	err = task2.FindCommonWords("../task_2/files/result.txt", "../task_2/files/t1", "../task_2/files/t2", "../task_2/files/t3")
	fmt.Println(err)

	a, b, rc := task1.FilterCommonDigits(1234, 498124)
	fmt.Println(a, b, rc)

}
