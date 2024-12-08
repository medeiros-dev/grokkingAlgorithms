package main

import "fmt"

func BinarySearch(nums []int, item int) (int, error) {

	down := 0
	up := len(nums) - 1

	for down <= up {
		try := (down + up) / 2

		if nums[try] == item {
			return try, nil
		}
		if nums[try] > item {
			up = try - 1
		} else {
			down = try + 1
		}
	}

	return -1, fmt.Errorf("This is not a valid number")
}
