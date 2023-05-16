package search_sort

// BubbleSort sorts the given slice in increasing order using the bubble sort algorithm.
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Iterate over the unsorted part of the slice (last i elements are already in place)
		for j := 0; j < n-i-1; j++ {
			// Traverse the array from 0 to n-i-1
			// Swap adjacent elements if they are in the wrong order
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// SelectionSort sorts the given slice in increasing order using the selection sort algorithm.
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// Find the index of the minimum element in the unsorted part of the slice
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum element with the first element in the unsorted part of the slice
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort sorts the given slice in increasing order using the insertion sort algorithm.
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		// Store the current element and its index
		current := arr[i]
		j := i - 1
		// Move the larger elements to the right to make space for the current element
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		// Insert the current element into its correct position
		arr[j+1] = current
	}
}

// MergeSort sorts the given slice in increasing order using the merge sort algorithm.
func MergeSort(arr []int) {
	n := len(arr)
	if n > 1 {
		// Divide the slice into two halves
		mid := n / 2
		left := make([]int, mid)
		right := make([]int, n-mid)
		copy(left, arr[:mid])
		copy(right, arr[mid:])
		// Recursively sort the two halves
		MergeSort(left)
		MergeSort(right)
		// Merge the sorted halves
		i, j, k := 0, 0, 0
		for i < len(left) && j < len(right) {
			if left[i] < right[j] {
				arr[k] = left[i]
				i++
			} else {
				arr[k] = right[j]
				j++
			}
			k++
		}
		for i < len(left) {
			arr[k] = left[i]
			i++
			k++
		}
		for j < len(right) {
			arr[k] = right[j]
			j++
			k++
		}
	}
}
