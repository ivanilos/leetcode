func numTeams(rating []int) int {
    result := 0
    for mid := 0; mid < len(rating); mid++ {
        bigLeft, smallLeft := countLeft(rating, mid)
        bigRight, smallRight := countRight(rating, mid)

        result += bigLeft * smallRight
        result += bigRight * smallLeft
    }
    return result
}

func countLeft(v []int, idx int) (int, int) {
    bigger := 0
    smaller := 0

    for i := idx - 1; i >= 0; i-- {
        if v[i] < v[idx] {
            smaller++
        } else if v[i] > v[idx] {
            bigger++
        }
    }
    return bigger, smaller
}

func countRight(v []int, idx int) (int, int) {
    bigger := 0
    smaller := 0

    for i := idx + 1; i < len(v); i++ {
        if v[i] < v[idx] {
            smaller++
        } else if v[i] > v[idx] {
            bigger++
        }
    }
    return bigger, smaller
}
