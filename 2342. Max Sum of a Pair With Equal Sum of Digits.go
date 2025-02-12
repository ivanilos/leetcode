func maximumSum(nums []int) int {
    digitSumMax := map[int]int{}

    result := -1
    for _, num := range nums {
        digitSum := getDigitSum(num)

        if _, ok := digitSumMax[digitSum]; ok {
            result = max(result, num + digitSumMax[digitSum])
            digitSumMax[digitSum] = max(digitSumMax[digitSum], num)
        } else {
            digitSumMax[digitSum] = num
        }
    }

    return result
}

func getDigitSum(num int) int {
    result := 0
    for num > 0 {
        result += num % 10
        num /= 10
    }

    return result;
}
