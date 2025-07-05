func findLucky(arr []int) int {
    freq := map[int]int{}

    for _, val := range arr {
        freq[val]++
    }

    result := -1
    for key, val := range freq {
        if key == val && key > result {
            result = key
        }
    }

    return result
}
