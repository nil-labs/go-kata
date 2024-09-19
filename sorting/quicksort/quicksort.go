// Quicksort is a divide-and-conquer algorithm. It works by selecting a 'pivot' element from the array
// and partitioning the other elements into two sub-arrays, according to whether they are less than or
// greater than the pivot. For this reason, it is sometimes called partition-exchange sort.
// The sub-arrays are then sorted recursively. This can be done in-place, requiring small additional
// mounts of memory to perform the sorting.
// Quicksort is a comparison sort, meaning that it can sort items of any type for which a "less-than"
// relation (formally, a total order) is defined. Efficient implementations of Quicksort are not a
// stable sort, meaning that the relative order of equal sort items is not preserved.

// Mathematical analysis of quicksort shows that, on average, the algorithm takes O(n log n) comparisons
// to sort n items. In the worst case, it makes O(n2) comparisons, though this behavior is rare.
package quicksort

func sort(arr []int, low, high int) {

	if low < high {
		pivot := partition(arr, low, high)
		sort(arr, low, pivot-1)
		sort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < len(arr); j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
