func candy(ratings []int) int {
    candies := make([]int, len(ratings))

    pos := make([]int, len(ratings))
    for i := 0; i < len(ratings); i++ {
        pos[i] = i
        candies[i] = 1
    }

    sort.Slice(pos, func(a, b int) bool {
        return ratings[pos[a]] < ratings[pos[b]]
    })

    for _, idx := range pos {
        neededVal := 1
        if idx - 1 >= 0 && ratings[idx - 1] < ratings[idx] {
            neededVal = max(neededVal, candies[idx - 1] + 1)
        }
        if idx + 1 < len(ratings) && ratings[idx + 1] < ratings[idx] {
            neededVal = max(neededVal, candies[idx + 1] + 1)
        }
        candies[idx] = neededVal
    }

    result := sum(candies)

    return result
}

func sum (v []int) int {
    result := 0
    for i := 0; i < len(v); i++ {
        result += v[i]
    }
    return result
}
