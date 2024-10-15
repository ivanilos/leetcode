func search(nums []int, target int) int {
    return solve(0, len(nums) - 1, target, nums)
}

func solve(left int, right int, target int, nums []int) int {
    if left > right {
        return -1
    }
    mid := (left + right) / 2

    if nums[mid] == target {
        return mid
    }

    if nums[left] > nums[right] {
        if nums[mid] >= nums[left] {
            return solveCrescentWithMidBigger(left, right, target, nums)
        } else {
            return solveCrescentWithMidSmaller(left, right, target, nums)
        }
    } else {
        return solveCrescent(left, right, target, nums)
    }
}

func solveCrescent(left, right, target int, nums []int) int {
    mid := (left + right) / 2

    if nums[mid] < target {
        return solve(mid + 1, right, target, nums)
    } else {
        return solve(left, mid - 1, target, nums)
    }
}

func solveCrescentWithMidBigger(left, right, target int, nums []int) int {
    mid := (left + right) / 2

    if nums[left] <= target && nums[mid] > target {
        return solve(left, mid - 1, target, nums)
    } else {
        return solve(mid + 1, right, target, nums)
    }
}

func solveCrescentWithMidSmaller(left, right, target int, nums []int) int {
    mid := (left + right) / 2

    if nums[mid] < target && target <= nums[right] {
        return solve(mid + 1, right, target, nums)
    } else {
        return solve(left, mid - 1, target, nums)
    }
}
