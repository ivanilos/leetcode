func successfulPairs(spells []int, potions []int, success int64) []int {
    result := make([]int, len(spells))

    slices.Sort(potions)

    for i, spell := range spells {
        result[i] = search(spell, potions, success)
    }

    return result
}

func search(spell int, potions []int, success int64) int {
    low := 0
    high := len(potions) - 1
    result := 0

    for low <= high {
        mid := (low + high) / 2

        if  int64(spell) * int64(potions[mid]) >= success {
            result = len(potions) - mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return result
}
