func findDifference(nums1 []int, nums2 []int) [][]int {
    result := [][]int{[]int{}, []int{}}

    slices.Sort(nums1)
    slices.Sort(nums2)

    p1, p2 := 0, 0
    for p1 < len(nums1) && p2 < len(nums2) {
        if nums1[p1] == nums2[p2] {
            curNum := nums1[p1]

            skip(curNum, nums1, &p1)
            skip(curNum, nums2, &p2)
        } else if nums1[p1] < nums2[p2] {
            curNum := nums1[p1]
            result[0] = append(result[0], nums1[p1])

            skip(curNum, nums1, &p1)
        } else {
            curNum := nums2[p2]
            result[1] = append(result[1], nums2[p2])

            skip(curNum, nums2, &p2)
        }
    }

    result[0] = append(result[0], getUnique(nums1, p1)...)
    result[1] = append(result[1], getUnique(nums2, p2)...)

    return result
}

func skip(curNum int, nums []int, p *int) {
    for *p < len(nums) && nums[*p] == curNum {
        (*p)++
    }
}

func getUnique(nums []int, p int) []int {
    result := []int{}
    for p < len(nums) {
        curNum := nums[p]
        result = append(result, nums[p])

        skip(curNum, nums, &p)
    }

    return result
}
