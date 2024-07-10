func minOperations(logs []string) int {
    depth := 0

    for _, cmd := range(logs) {
        if cmd == "../" {
            depth = max(depth - 1, 0)
        } else if cmd != "./" {
            depth++
        }
    }
    return depth
}
