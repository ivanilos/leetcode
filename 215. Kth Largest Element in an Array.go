func findKthLargest(nums []int, k int) int {
    return findKth(nums, len(nums) - k)
}

func findKth(nums []int, k int) int {
    pivotPos := rand.IntN(len(nums))
    pivot := nums[pivotPos]

    left := 0
    mid := 0
    right := len(nums) - 1

    for mid <= right {
        if nums[mid] < pivot {
            nums[left], nums[mid] = nums[mid], nums[left]
            left++
            mid++
        } else if nums[mid] == pivot {
            mid++
        } else {
            nums[mid], nums[right] = nums[right], nums[mid]
            right--
        }
    }


    if left - 1 >= k {
        return findKth(nums[: left], k)
    } else if k > right {
        return findKth(nums[right + 1 : len(nums)], k - (right + 1))
    } else {
        return pivot
    }
}
