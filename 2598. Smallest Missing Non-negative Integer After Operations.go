func findSmallestInteger(nums []int, value int) int {
    hasResidue := make([]int, value + 1)

    for _, num := range nums {
        hasResidue[((num % value) + value) % value]++
    }

    result := 0
    for hasResidue[result % value] > 0 {
        hasResidue[result % value]--
        result++
    }

    return result
}
