func twoSum(numbers []int, target int) []int {
    leftIdx := 0
    rightIdx := len(numbers) - 1

    for leftIdx < rightIdx {
        if numbers[leftIdx] + numbers[rightIdx] < target {
            leftIdx++
        } else if numbers[leftIdx] + numbers[rightIdx] > target {
            rightIdx--
        } else {
            return []int{leftIdx + 1, rightIdx + 1}
        }
    }
    panic("No solution! Should not reach here")
}
