package a2

import (
	"fmt"
	"testing"
)

func TestSearchValue(t *testing.T) {
	src := []int{4, 5, 6, 7, 0, 1, 2}
	target := 4
	index := searchValue(src, target)
	if index != 0 {
		t.Errorf("#### search case 1 failed")
		fmt.Println("eror index = %d", index)
		return
	} else {
		fmt.Println("===== search case 1 success index =%d", index)
	}

	index = searchValue(src, 1)
	if index != 5 {
		t.Errorf("#### search case 2 failed")
		fmt.Println("eror index = %d", index)
		return
	} else {
		fmt.Println("===== search case 2 success index = %d", index)
	}

	index = searchValue(src, 2)
	if index != 6 {
		t.Errorf("#### search case 3 failed")
		fmt.Println("eror index = %d", index)
		return
	} else {
		fmt.Println("===== search case 3 success index = %d", index)
	}
}

func TestBinarySearch(t *testing.T) {
	src := []int{0, 1, 3, 5, 7, 9, 18, 25}
	index := binarySearch(src, 0, len(src)-1, 3)
	if index != 2 {
		t.Errorf("######## binary search case 1 failed")
	} else {
		fmt.Println("===case 1 return %d", index)
	}

	index = binarySearch(src, 0, len(src)-1, 0)
	if index != 0 {
		t.Errorf("######## binary search case 2 failed")
	} else {
		fmt.Println("===case 2 return %d", index)
	}

	index = binarySearch(src, 0, len(src)-1, 25)
	if index != 7 {
		t.Errorf("######## binary search case 3 failed")
	} else {
		fmt.Println("===case 3 return %d", index)
	}
}
