func largestGoodInteger(num string) string {
    result := ""

    for i := 2; i < len(num); i++ {
        if num[i] == num[i - 1] && num[i - 1] == num[i - 2] {
            candidate := num[i - 2 : i + 1]
            if result == "" || candidate > result {
                result = candidate
            }
        }
    }

    return result
}
