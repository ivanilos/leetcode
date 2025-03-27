func minimumIndex(nums []int) int {
    dominant, count := getDominant(nums)

    curQt := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == dominant {
            curQt++
        }

        if curQt * 2 > i + 1 && (count - curQt) * 2 > len(nums) - 1 - i {
            return i
        }
    }

    return -1
}

func getDominant(nums []int) (int, int) {
    candidate := nums[0]
    qt := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] == candidate {
            qt++
        } else if qt == 0 {
            candidate = nums[i]
            qt = 1
        } else {
            qt--
        }
    }

    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == candidate {
            count++
        }
    }

    return candidate, count
}
