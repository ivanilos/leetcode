func xorQueries(arr []int, queries [][]int) []int {
    pref := make([]int, len(arr))

    pref[0] = arr[0]
    for i := 1; i < len(arr); i++ {
        pref[i] = arr[i] ^ pref[i - 1]
    }

    result := []int{}
    for _, query := range queries {
        left := query[0]
        right := query[1]

        val := pref[right]
        if left > 0 {
            val ^= pref[left - 1]
        }
        result = append(result, val)
    }
    return result
}
