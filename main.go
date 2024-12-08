package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index, err := BinarySearch(nums, 11)
	if err != nil {
		panic(err)
	}
	fmt.Println("The number was found at index:", index)

}
