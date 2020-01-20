package main

import "fmt"

func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
	return arr
}

func main() {
	arr := []int{2, 3, 8, 9, 12, 0}
	fmt.Println(bubbleSort(arr))
}
