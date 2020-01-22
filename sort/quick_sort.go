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

		//if l<r {
		//	tmp := arr[l]
		//	arr[l] = arr[r]
		//	arr[r] = tmp
		//}
	}
	arr[r] = x
	return r
}

func QuickSort(arr []int, l int, r int) []int {
	if l >= r {
		return nil
	}
	i := partition2(arr, l, r)
	QuickSort(arr, l, i-1)
	QuickSort(arr, i+1, r)
	return arr
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
