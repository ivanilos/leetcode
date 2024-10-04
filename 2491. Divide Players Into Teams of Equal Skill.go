func dividePlayers(skill []int) int64 {
    sort.Ints(skill)

    result := int64(0)
    need := skill[0] + skill[len(skill) - 1]
    for i, j := 0, len(skill) - 1; i < j; i, j = i + 1, j - 1 {
        if skill[i] + skill[j] != need {
            return -1
        }
        result += int64(skill[i] * skill[j])
    }
    return result
}
