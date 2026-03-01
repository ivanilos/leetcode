func minPartitions(n string) int {
    digits := strings.Split(n, "")

    result := 0
    for i := 0; i < len(digits); i++ {
        d, _ := strconv.Atoi(digits[i])
        result = max(result, d)
    }

    return result
}
