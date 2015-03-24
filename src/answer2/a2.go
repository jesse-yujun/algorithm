package a2

import (
//"errors"
)

/*
Suppose a sorted array is rotated at some pivot unknown to you beforehand.
(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
You are given a target value to search. If found in the array return its index, otherwise return -1.
You may assume no duplicate exists in the array.
*/

func searchValue(src []int, target int) int {
	headIndex := 0
	tailIndex := len(src) - 1
	middleIndex := (tailIndex-headIndex)/2 + headIndex

	for middleIndex > headIndex && middleIndex < tailIndex {
		if src[middleIndex] == target {
			return middleIndex
		}

		if (src[headIndex] < src[middleIndex]) && (src[middleIndex] > target) && (src[headIndex] <= target) { //左侧是有序数组且包含target
			return binarySearch(src, headIndex, middleIndex, target)
		} else {
			headIndex = middleIndex + 1
			middleIndex = (tailIndex-headIndex)/2 + headIndex
		}
	}

	if src[headIndex] == target {
		return headIndex
	}

	if src[tailIndex] == target {
		return tailIndex
	}
	//err := errors.New("the input array is invalid")
	return -1
}

func binarySearch(src []int, headIndex int, tailIndex int, target int) int {
	middleIndex := (tailIndex-headIndex)/2 + headIndex

	for middleIndex > headIndex && middleIndex < tailIndex {
		if src[middleIndex] == target {
			return middleIndex
		} else {
			if src[middleIndex] > target {
				tailIndex = middleIndex - 1
			} else {
				headIndex = middleIndex + 1
			}
			middleIndex = (tailIndex-headIndex)/2 + headIndex
		}
	}

	if src[headIndex] == target {
		return headIndex
	}

	if src[tailIndex] == target {
		return tailIndex
	}

	return -1
}
