package main

import (
	"fmt"
	"dz1/internal/task_1"
	"dz1/internal/task_2"
	"dz1/internal/task_3"
)

func main(){

	// slice := []int{}
	err := task3.ScaleSlice(nil, 5)
	fmt.Println(err)

	err = task2.FindCommonWords("../internal/task_2/files/result.txt", "../internal/task_2/files/t1", "../internal/task_2/files/t2", "../internal/task_2/files/t3")
	fmt.Println(err)

	a, b, rc := task1.FilterCommonDigits(1234, 498124)
	fmt.Println(a, b, rc)

}
