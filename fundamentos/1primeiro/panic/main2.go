package main

import (
	"fmt"
	"log"
)

func recoveryFunction() {
	if err := recover(); err != nil {
		log.Println("panic occurred:", err)
	}
	fmt.Println("This is recovery function...")
}

func executePanic() {
	defer recoveryFunction()
	var b [3]float64 // erro proposital
	for i := 0; i < 5; i++ {
		b[i] = 1.1
	}
	//panic("This is Panic Situation")
	fmt.Println("The function executes Completely")
}

func main() {
	executePanic()
	fmt.Println("Main block is executed completely...")
}
