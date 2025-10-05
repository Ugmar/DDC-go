package main

import (
	"fmt"
	// "dz1/task_1"
	"dz1/task_2"
)

func main(){
	rc := task2.FindCommonWords("result", "../task_2/files/t1", "../task_2/files/t2", "../task_2/files/t3")
	fmt.Println(rc)
}
