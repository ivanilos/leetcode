func distance(nums []int) []int64 {
    valToIdx := map[int][]int{}

    for i := 0; i < len(nums); i++ {
        valToIdx[nums[i]] = append(valToIdx[nums[i]], i)
    }

    result := make([]int64, len(nums))
    for _, list := range valToIdx {

        sum := int64(0)
        for i := 0; i < len(list); i++ {
            idx := list[i]

            if i > 0 {
                sum += int64(i) * int64(list[i] - list[i - 1])
                result[idx] += sum
            }
        }

        sum = 0
        for i := len(list) - 1; i >= 0; i-- {
            idx := list[i]
            prev := int64(len(list) - 1 - i)

            if i < len(list) - 1 {
                sum += (prev) * int64(list[i + 1] - list[i])
                result[idx] += sum
            }
        }
    }

    return result
}
