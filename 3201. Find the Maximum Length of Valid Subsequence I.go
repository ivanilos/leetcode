func maximumLength(nums []int) int {
    allOdd := 0
    allEven := 0
    alternatingWithLastOdd := 0
    alternatingWithlastEven := 0

    for _, num := range nums {
        if num % 2 == 0 {
            allEven++
            alternatingWithlastEven = max(alternatingWithlastEven, alternatingWithLastOdd + 1)
        } else {
            allOdd++
            alternatingWithLastOdd = max(alternatingWithLastOdd, alternatingWithlastEven + 1)
        }
    }

    return max(allOdd, allEven, alternatingWithLastOdd, alternatingWithlastEven)
}
