func maxSumDivThree(nums []int) int {
    sum := 0

    modNum := make([][]int, 3)
    for i := 0; i < 3; i++ {
        modNum[i] = []int{}
    }

    for _, num := range nums {
        modNum[num % 3] = append(modNum[num % 3], num)

        sum += num
    }

    for i := 0; i < 3; i++ {
        slices.Sort(modNum[i])
    }

    if sum % 3 == 0 {
        return sum
    } else if sum % 3 == 1 {
        a := 0
        b := 0

        if len(modNum[1]) >= 1 {
            a = sum - modNum[1][0]
        }
        if len(modNum[2]) >= 2 {
            b = sum - modNum[2][0] - modNum[2][1]
        }

        return max(a, b)
    } else {
        a := 0
        b := 0

        if len(modNum[1]) >= 2 {
            a = sum - modNum[1][0] - modNum[1][1]
        }
        if len(modNum[2]) >= 1 {
            b = sum - modNum[2][0]
        }

        return max(a, b)
    }
}
