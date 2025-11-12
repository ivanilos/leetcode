const INF = int(1e9)

func minOperations(nums []int) int {
    onesQt := 0

    for _, num := range nums {
        if num == 1 {
            onesQt++
        }
    }

    if onesQt > 0 {
        return len(nums) - onesQt
    }

    minLenToUnitGCD := INF
    for start := 0; start < len(nums); start++ {
        curGCD := 0
        for i := start; i < len(nums); i++ {
            curGCD = gcd(curGCD, nums[i])

            if curGCD == 1 {
                minLenToUnitGCD = min(minLenToUnitGCD, i - start + 1)
                break
            }
        }
    }

    if minLenToUnitGCD == INF {
        return -1
    } else {
        return minLenToUnitGCD - 1 + len(nums) - 1
    }
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}
