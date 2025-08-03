func maxTotalFruits(fruits [][]int, startPos int, k int) int {
    totalFruits := 0
    for i := 0; i < len(fruits); i++ {
        totalFruits += fruits[i][1]
    }

    leftThenRight := getMaxTotalFruits(fruits, startPos, k, totalFruits)
    
    slices.Reverse(fruits)

    for i := 0; i < len(fruits); i++ {
        fruits[i][0] = -fruits[i][0]
    }

    rightThenLeft := getMaxTotalFruits(fruits, -startPos, k, totalFruits)

    return max(leftThenRight, rightThenLeft)
}

func getMaxTotalFruits(fruits [][]int, startPos, k, totalFruits int) int {
    right := len(fruits) - 1

    result := 0
    for left := 0; left < len(fruits) && fruits[left][0] <= startPos; left++ {
        if fruits[left][0] > startPos {
            break
        }

        for right > left && dist(startPos, left, right, fruits) > k {
            totalFruits -= fruits[right][1]
            right--
        }

        for right + 1 < len(fruits) && dist(startPos, left, right + 1, fruits) <= k {
            right++
            totalFruits += fruits[right][1]
        }

        if dist(startPos, left, right, fruits) <= k {
            result = max(result, totalFruits)
        }

        totalFruits -= fruits[left][1]
    }

    return result
}

func dist(startPos, left, right int, fruits [][]int) int {
    if right >= 0 && fruits[right][0] > startPos {
        return startPos - fruits[left][0] + (fruits[right][0] - fruits[left][0])
    } else {
        return startPos - fruits[left][0]
    }
}
