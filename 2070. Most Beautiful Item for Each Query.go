func maximumBeauty(items [][]int, queries []int) []int {
    sort.Slice(items, func(a, b int) bool {
        return items[a][0] < items[b][0] || items[a][0] == items[b][0] && items[a][1] < items[b][1]
    })

    maxBeautyAtPrice := getMaxBeautyAtPrice(items)

    result := make([]int, len(queries))
    for i, q := range queries {
        ans := bsearch(maxBeautyAtPrice, q)

        result[i] = ans
    }
    return result
}

func bsearch(v [][]int, val int) int {
    low := 0
    high := len(v) - 1
    best := 0

    for low <= high {
        mid := (low + high) / 2

        if v[mid][0] <= val {
            best = v[mid][1]
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return best
}

func getMaxBeautyAtPrice(items [][]int) [][]int {
    maxBeautyAtPrice := make([][]int, len(items))

    for i := 0; i < len(items); i++ {
        maxBeautyAtPrice[i] = make([]int, 2)
    }

    maxBeautyAtPrice[0][0] = items[0][0]
    maxBeautyAtPrice[0][1] = items[0][1]

    for i := 1; i < len(items); i++ {
        maxBeautyAtPrice[i][0] = items[i][0]
        maxBeautyAtPrice[i][1] = max(items[i][1], maxBeautyAtPrice[i - 1][1])
    }

    return maxBeautyAtPrice
}
