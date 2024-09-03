func getLucky(s string, k int) int {
    result := transform(s)

    for i := 1; i < k; i++ {
        result = getSum(result)
    }
    return result
}

func getSum(val int) int {
    result := 0
    for val > 0 {
        result += val % 10
        val /= 10
    }
    return result
}

func transform(s string) int {
    result := 0
    for _, c := range s {
        val := int(c - 'a') + 1
        for val > 0 {
            result += val % 10
            val /= 10
        }
    }
    return result
}
