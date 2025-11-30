package demoday2

import "fmt"

func SliceDemo() {
	var slice []int = make([]int, 10)
	for i := 0; i < len(slice); i++ {
		slice[i] = i
	}
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i])
	}
	for _, value := range slice {
		fmt.Print(value)

	}
	slice = append(slice, 1, 2, 3)

}
