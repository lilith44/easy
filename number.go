package easy

type number interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64
}

// Max figures out the maximum value of the given nums.
func Max[n number](nums ...n) (result n) {
	if len(nums) == 0 {
		return
	}

	result = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > result {
			result = nums[i]
		}
	}
	return
}

// Min figures out the minimum value of the given nums.
func Min[n number](nums ...n) (result n) {
	if len(nums) == 0 {
		return
	}

	result = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < result {
			result = nums[i]
		}
	}
	return
}

// Sum gets the sum of the given nums.
func Sum[n number](nums ...n) (result n) {
	if len(nums) == 0 {
		return
	}

	for i := range nums {
		result += nums[i]
	}
	return
}
