func countSeniors(details []string) int {
    result := 0
    for _, detail := range(details) {
        ageStart := 10 + 1
        age, _ := strconv.Atoi(detail[ageStart : ageStart + 2])
        if age > 60 {
            result++
        }
    }
    return result
}
