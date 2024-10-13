package sliding_window

// Find the length of the longest subarray with the same value in each position

func longestSubarray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	left := 0

	maxLen := 0

	for right := range arr {
		if arr[right] != arr[left] {
			left = right
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

//find the minimum length subarray where the sum is greater than or equal to the target.
//assume all values are positive

func minSubArrayLen(target int, nums []int) int {
	l, total := 0, 0

	// SETTING THE MAXIMUM LENGTH OF THE POSSIBLE SUBARRAY
	length := 4000

	for r := range nums {
		total += nums[r]

		for total >= target {
			length = min(length, r-l+1)
			total -= nums[l]
			l++
		}
	}
	if length == 4000 {
		return 0
	}
	return length

}
