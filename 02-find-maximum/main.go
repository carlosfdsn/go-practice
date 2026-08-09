package main

import "fmt"

func findMaximum(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}

func main() {

	nums := []int{4, 12, 7, 25, 9}

	fmt.Println(findMaximum(nums))

}
