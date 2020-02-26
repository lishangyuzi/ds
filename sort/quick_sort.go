package sort

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
	}
	arr[r] = x
	return r
}

func QuickSort(arr []int, l int, r int) []int {
	if l >= r {
		return nil
	}
	i := partition(arr, l, r)
	QuickSort(arr, l, i-1)
	QuickSort(arr, i+1, r)
	return arr
}
