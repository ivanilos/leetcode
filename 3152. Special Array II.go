func isArraySpecial(nums []int, queries [][]int) []bool {
    specialToLeft, specialToRight := getSpecialToSidePrefixSum(nums)

    fmt.Println(specialToLeft, specialToRight)

    result := make([]bool, len(queries))
    for i := 0; i < len(queries); i++ {
        leftIdx := queries[i][0]
        rightIdx := queries[i][1]

        if isLeftSpecial(leftIdx, rightIdx, specialToLeft) && isRightSpecial(leftIdx, rightIdx, specialToRight) {
            result[i] = true
        } else {
            result[i] = false
        }
    }

    return result
}

func getSpecialToSidePrefixSum(nums []int) ([]int, []int) {
    specialToLeft := make([]int, len(nums))
    specialToLeft[0] = 1

    for i := 1; i < len(nums); i++ {
        toAdd := 0
        if nums[i] % 2 != nums[i - 1] % 2 {
            toAdd = 1
        }
        specialToLeft[i] = specialToLeft[i - 1] + toAdd
    }

    specialToRight := make([]int, len(nums))
    specialToRight[len(nums) - 1] = 1

    for i := len(nums) - 2; i >= 0; i-- {
        toAdd := 0
        if nums[i] % 2 != nums[i + 1] % 2 {
            toAdd = 1
        }
        specialToRight[i] = specialToRight[i + 1] + toAdd
    }

    return specialToLeft, specialToRight
}

func isLeftSpecial(l, r int, special []int) bool {
    total := special[r] - special[l]
    return total == r - l
}

func isRightSpecial(l, r int, special []int) bool {
    total := special[l] - special[r]
    return total == r - l
}
