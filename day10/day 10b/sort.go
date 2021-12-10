package main

func mergeSort(nums []int, start int, end int) []int {
	if start == end {
		return []int{nums[start]}
	}

	var mid int = ((end - start) / 2) + start

	//Assign values back into Left and right
	left := mergeSort(nums, start, mid)
	right := mergeSort(nums, mid+1, end)

	var combined []int

	//Pointers for new array
	leftPointer, rightPointer := 0, 0

	for leftPointer <= len(left)-1 || rightPointer <= len(right)-1 {

		if leftPointer == len(left) {
			addValue(&combined, right[rightPointer], &rightPointer)
		} else if rightPointer == len(right) {
			addValue(&combined, left[leftPointer], &leftPointer)
		} else {
			if left[leftPointer] <= right[rightPointer] {
				addValue(&combined, left[leftPointer], &leftPointer)
			} else {
				addValue(&combined, right[rightPointer], &rightPointer)
			}
		}
	}
	return combined
}

func addValue(nums *[]int, value int, pointer *int) {
	*nums = append(*nums, value)
	*pointer++
}
