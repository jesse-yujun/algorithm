package a1

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicate(t *testing.T) {
	src := []int{1, 2, 2, 2, 3, 5, 5, 7}
	des, err := removeDuplicate(src)
	if len(des) != 5 {
		t.Errorf("there is an error")
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Print("case 1 pass")
		for _, v := range des {
			fmt.Println(v)
		}
	}
}

func TestRemoveDuplicate2(t *testing.T) {
	src := []int{}
	des, err := removeDuplicate(src)
	if len(des) != 0 {
		t.Errorf("there is an error")
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Print("case 2 pass")
	}
}
