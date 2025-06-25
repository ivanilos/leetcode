func kthSmallestProduct(nums1 []int, nums2 []int, k int64) int64 {
    nums2Pos, nums2Neg  := getSortedValuesBySign(nums2)

    low := int64(-1e10)
    high := int64(1e10)
    result := low

    for low <= high {
        mid := (low + high) / 2

        if hasEnoughValues(nums1, nums2Pos, nums2Neg, mid, k) {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}

func getSortedValuesBySign(nums []int) ([]int, []int) {
    pos := []int{}
    neg := []int{}

    for _, num := range nums {
        if num >= 0 {
            pos = append(pos, num)
        } else {
            neg = append(neg, num)
        }
    }

    slices.Sort(pos)
    slices.Sort(neg)

    return pos, neg
}

func hasEnoughValues(nums1, pos, neg []int, maxVal int64, k int64) bool {
    total := int64(0)

    for _, num := range nums1 {
        curNum := int64(num)
        
        if curNum == 0 {
            if maxVal >= 0 {
                total += int64(len(pos) + len(neg))
            }
        } else if curNum < 0 {
            total += searchLeftIncreases(neg, curNum, maxVal) + searchLeftIncreases(pos, curNum, maxVal)
        } else {
            total += searchRightIncreases(neg, curNum, maxVal) + searchRightIncreases(pos, curNum, maxVal)
        }
    }

    return total >= k
}

func searchLeftIncreases(v []int, cur int64, maxVal int64) int64 {
    low := 0
    high := len(v) - 1
    result := 0

    for low <= high {
        mid := (low + high) / 2

        val := cur * int64(v[mid])

        if val > maxVal {
            low = mid + 1
        } else {
            result = len(v) - mid
            high = mid - 1
        }
    }

    return int64(result)
}

func searchRightIncreases(v []int, cur int64, maxVal int64) int64 {
    low := 0
    high := len(v) - 1
    result := 0

    for low <= high {
        mid := (low + high) / 2

        val := cur * int64(v[mid])

        if val > maxVal {
            high = mid - 1
        } else {
            result = mid + 1
            low = mid + 1
        }
    }

    return int64(result)
}
