package main

import "fmt"

// BubbleSort
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Exchange elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var test func()
	test = func() {
		fmt.Println("test func")
		defer fmt.Println("defer func")
	}
	fmt.Println("main func")
	test()
}
