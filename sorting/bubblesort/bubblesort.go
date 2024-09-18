package bubblesort

/**
Bubble sort
Space O(1)
Time O(N^2)
Stable sorting algorithm, meaning that elements with the same key value maintain their relative order in the sorted output.
Comparison based sorting algorythm
*/

func sort(in []int) {
	for i := 0; i < len(in); i++ {
		for j := i + 1; j < len(in); j++ {
			if in[i] > in[j] {
				in[i], in[j] = in[j], in[i]
			}
		}
	}
}
