func numTilePossibilities(tiles string) int {
    freq := map[rune]int{}

    for _, c := range tiles {
        freq[c]++
    }

    v := slices.Collect(maps.Values(freq))

    used := make([]int, len(v))
    result := solve(v, used, -1)

    return result - 1
}

func solve(v, used []int, lastIdx int) int {
    result := 1
    for i := 0; i < len(v); i++ {
        if i == lastIdx {
            continue
        }
        for j := 1; j <= v[i] - used[i]; j++ {
            used[i] += j
            result += solve(v, used, i)
            used[i] -= j
        }
    }

    return result
}
