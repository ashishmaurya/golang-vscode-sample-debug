package main

import (
	"fmt"
)

func main() {
	fmt.Println("Namste!")
	for i := 0; i < 10; i++ {
		//You can add debug and step through each iteration
		result := fmt.Sprintf("task %d", i)
		fmt.Println(result)
	}
}
