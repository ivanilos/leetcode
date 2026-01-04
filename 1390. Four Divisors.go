const MAX_NUM = int(1e5)

func sumFourDivisors(nums []int) int {
    divisors, divisorsSum := getDivisorsCount(MAX_NUM)

    result := 0
    for _, num := range nums {
        if divisors[num] == 4 {
            result += divisorsSum[num]
        }
    }

    return result
}

func getDivisorsCount(maxNum int) ([]int, []int) {
    divisorsCount := make([]int, maxNum + 1)
    divisorsSum := make([]int, maxNum + 1)

    for i := 1; i <= maxNum; i++ {
        for j := i; j <= maxNum; j += i {
            divisorsCount[j]++
            divisorsSum[j] += i
        }
    }

    return divisorsCount, divisorsSum
}
