package twosum

import "fmt"

func twoSum(nums []int, target int) []int {
	var indices []int

	// prevIdx := 0

	if len(nums) == 2 {
		indices = append(indices, 0)
		indices = append(indices, 1)
		return indices
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums)-1; j++ {
			if (nums[i] + nums[j]) == target {
				indices = append(indices, i)
				indices = append(indices, j)
				// indices[0] = prevIdx
				// indices[1] = i
				fmt.Println("idx1 = ", i)
				fmt.Println("idx2 = ", j)
				return indices
			}
		}

		// prevIdx = i
	}

	return indices
}
