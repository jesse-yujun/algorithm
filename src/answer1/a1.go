package a1

import (
	"errors"
)

func removeDuplicate(src []int) ([]int, error) {
	des := make([]int, 0)

	if src == nil {
		return des, errors.New("the input array is invalid")
	}

	length := len(src)
	if length == 0 {
		return des, errors.New("the input array is empty")
	}

	if length == 1 {
		des = src[0:1]
		return des, nil
	}

	tail := 0
	for i := 1; i < length; i++ {
		if src[tail] == src[i] {
			//do nothing
		} else {
			tail++
			src[tail] = src[i]
		}
	}

	des = src[0 : tail+1]

	return des, nil
}
