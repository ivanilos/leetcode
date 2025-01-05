const ALPHA int = 26

func shiftingLetters(s string, shifts [][]int) string {
    todoShift := make([]int, len(s) + 1)

    for _, shift := range shifts {
        start, end, dir := getParams(shift)

        todoShift[start] = sum(todoShift[start], dir)
        todoShift[end + 1] = sum(todoShift[end + 1], -dir)
    }

    result := make([]byte, len(s))
    curShift := 0
    for i := 0 ; i < len(s); i++ {
        curShift += todoShift[i]
        curCharIdx := sum(int(s[i] - 'a'), curShift)
        result[i] = byte(curCharIdx + 'a')
    }

    return string(result)
}

func getParams(shift []int) (int, int, int) {
    start := shift[0]
    end := shift[1]
    dir := 1

    if shift[2] == 0 {
        dir = -1
    }
    return start, end, dir
}

func sum(a, b int) int {
    return (((a + b) % ALPHA) + ALPHA) % ALPHA
}
