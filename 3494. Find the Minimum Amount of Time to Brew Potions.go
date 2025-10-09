func minTime(skill []int, mana []int) int64 {
    skill = append(skill, 0)

    times := make([]int64, len(skill))
    times[len(skill) - 1] = 0
    for _, potion := range mana {
        for i := 1; i < len(skill); i++ {
            lastSkillSpend := int64(skill[i - 1]) * int64(potion)
            times[i] = max(times[i], times[i - 1] + lastSkillSpend)
        }

        for i := len(skill) - 2; i >= 0; i-- {
            nextSkillSpend := int64(skill[i + 1]) * int64(potion)
            times[i] = times[i + 1] - nextSkillSpend 
        }
    }

    return times[len(skill) - 1]
}
