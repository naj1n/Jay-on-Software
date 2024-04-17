package main

import (
	"fmt"
	"math/rand"
	"time"
)

// SelectionSort is a simple sorting algorithm.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[k] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
}

// BubbleSort1 is a simple sorting algorithm.
func BubbleSort1(arr []int) {
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				// Exchange elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// BubbleSort2 is a simple sorting algorithm.
func BubbleSort2(arr []int) {
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

// BubbleSort3 is a simple sorting algorithm.
func BubbleSort3(arr []int) {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if arr[i] > arr[i+1] {
				// Exchange elements
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
}

// InsertionSort is a simple sorting algorithm.
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// partition is the important part of QuickSort.
func partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]

	return i
}

// QuickSort is a good sorting algorithm.
func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition(arr, left, right)
	QuickSort(arr, left, pivot-1)
	QuickSort(arr, pivot+1, right)
}

// merge is the important part of MergeSort.
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// MergeSort is a simple algorithm.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := generateRandomSlice(10)
	fmt.Println("Original slice:", slice)
	start := time.Now()
	res := MergeSort(slice)
	duration := time.Since(start)
	fmt.Println("Sorted slice:", res)
	fmt.Println("Time required for bubble sort: ", duration)
}

// generateRandomSlice generates a slice of size n filled with random numbers
func generateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Intn(100) // generate integers up to 100
	}
	return slice
}
