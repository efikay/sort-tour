package sort

import "github.com/charmbracelet/bubbles/list"

type ShortSortInfo struct {
	name        string
	description string
}

func (i ShortSortInfo) Title() string       { return i.name }
func (i ShortSortInfo) Description() string { return i.description }
func (i ShortSortInfo) FilterValue() string { return i.name }

func GetSortInfoListItems() []list.Item {
	return []list.Item{
		ShortSortInfo{name: "Bubble Sort", description: "Repeatedly swaps adjacent elements if they're in the wrong order"},
		ShortSortInfo{name: "Selection Sort", description: "Selects the smallest element and swaps it into place each iteration"},
		ShortSortInfo{name: "Insertion Sort", description: "Builds the sorted array one element at a time by inserting each item into its correct position"},
		ShortSortInfo{name: "Merge Sort", description: "Divides the array into halves, recursively sorts, then merges them back together"},
		ShortSortInfo{name: "Quick Sort", description: "Picks a 'pivot' element and partitions the array around it, recursively sorting sub-arrays"},
		ShortSortInfo{name: "Heap Sort", description: "Converts the array into a max-heap to repeatedly extract the largest element"},
		ShortSortInfo{name: "Counting Sort", description: "Counts occurrences of each value (works only for integer ranges)"},
		ShortSortInfo{name: "Radix Sort", description: "Sorts numbers digit-by-digit (from least to most significant)"},
		ShortSortInfo{name: "Bucket Sort", description: "Distributes elements into buckets, sorts each bucket, then concatenates"},
		ShortSortInfo{name: "Shell Sort", description: "An optimized insertion sort that compares elements spaced apart"},
		ShortSortInfo{name: "Tim Sort", description: "Hybrid of merge sort and insertion sort (used in Python and Java)"},
		ShortSortInfo{name: "Comb Sort", description: "Improves bubble sort by comparing elements spaced apart (shrinking gap)"},
		ShortSortInfo{name: "Cycle Sort", description: "Minimizes writes by cycling elements to their correct positions"},
		ShortSortInfo{name: "Cocktail Shaker Sort", description: "Bidirectional bubble sort that alternates directions"},
		ShortSortInfo{name: "Gnome Sort", description: "Similar to insertion sort but swaps backward like a garden gnome"},
		ShortSortInfo{name: "Bogo Sort", description: "Randomly shuffles the array until it's sorted (theoretically terrible)"},
		ShortSortInfo{name: "Stooge Sort", description: "Recursively sorts first 2/3, last 2/3, then first 2/3 again"},
		ShortSortInfo{name: "Pancake Sort", description: "Flips subarrays to move the largest element to its position"},
		ShortSortInfo{name: "Bitonic Sort", description: "Works best on power-of-two-sized arrays (parallelizable)"},
		ShortSortInfo{name: "Tree Sort", description: "Uses a binary search tree to sort elements"},
		ShortSortInfo{name: "Cube Sort", description: "Parallel version of bubble sort for 3D data"},
		ShortSortInfo{name: "Odd-Even Sort", description: "Compares odd/even indexed pairs (parallelizable)"},
		ShortSortInfo{name: "Strand Sort", description: "Extracts sorted 'strands' from the list and merges them"},
		ShortSortInfo{name: "Intro Sort", description: "Hybrid of quick sort, heap sort, and insertion sort"},
		ShortSortInfo{name: "Smooth Sort", description: "Adaptive variant of heap sort with near-O(n) best-case performance"},
	}
}
