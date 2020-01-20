package main

import "fmt"

func partition(arr []int, l int, r int) int {
	x := arr[l]
	for l < r {
		for arr[r] >= x && l < r {
			r--
		}
		arr[l] = arr[r]
		for arr[l] <= x && l < r {
			l++
		}
		arr[r] = arr[l]

		//if l<r {
		//	tmp := arr[l]
		//	arr[l] = arr[r]
		//	arr[r] = tmp
		//}
	}
	arr[r] = x
	return r
}

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	i := partition2(arr, l, r)
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)

}

func partition2(arr []int, l int, r int) int {
	x := arr[l]
	start := l
	for l < r {

		for arr[r] >= x && l < r {
			r--
		}
		for arr[l] <= x && l < r {
			l++
		}

		if l < r {
			tmp := arr[r]
			arr[r] = arr[l]
			arr[l] = tmp
		}
	}
	arr[start] = arr[l]
	arr[l] = x

	return l

}

func main() {
	arr := []int{2, 3, 8, 9, 12, 0}
	quickSort(arr, 0, 5)
	fmt.Println(arr)
}
