type elem struct {
    val int
    idx int
}

func lexicographicallySmallestArray(nums []int, limit int) []int {
    v := make([]elem, len(nums))
    for idx, num := range nums {
        v[idx] = elem{num, idx}
    }

    sort.Slice(v, func(a, b int) bool {
        return v[a].val < v[b].val
    })

    result := make([]int, len(nums))

    last := v[0].val
    firstIdx := 0
    for i := 0; i < len(v); i++ {
        if last != -1 && v[i].val - last > limit {
            updateResult(firstIdx, i - 1, result, v)

            firstIdx = i
        }
        last = v[i].val
    }
    updateResult(firstIdx, len(nums) - 1, result, v)    

    return result
}

func updateResult(firstIdx, lastIdx int, result []int, v []elem) {
    aux := make([]elem, lastIdx -firstIdx + 1)

    for i := 0; i <= lastIdx - firstIdx; i++ {
        aux[i] = elem{v[i + firstIdx].idx, v[i + firstIdx].val}
    }

    sort.Slice(aux, func(a, b int) bool {
        return aux[a].val < aux[b].val
    })


    for i := 0; i <= lastIdx - firstIdx; i++ {
        result[aux[i].val] = v[i + firstIdx].val
    }
}
