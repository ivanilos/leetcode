func kthDistinct(arr []string, k int) string {
    freq := map[string]int{}

    for _, val := range(arr) {
        freq[val]++
    }

    result := ""
    for i := 0; i < len(arr); i++ {
        if freq[arr[i]] == 1 {
            k--
            if k == 0 {
                result = arr[i]
                break
            }
        }
    }
    return result
}
