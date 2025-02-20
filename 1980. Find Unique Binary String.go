func findDifferentBinaryString(nums []string) string {
    generated := []string{}

    curString := []rune{}
    DFS(0, len(nums), curString, &generated)

    for _, str := range generated {
        if !isIn(str, nums) {
            return str
        }
    }
    return ""
}

func isIn(str string, nums []string) bool {
    for _, a := range nums {
        if str == a {
            return true
        }
    }
    return false
}

func DFS(pos, sz int, curString []rune, generated *[]string) {
    if pos == sz {
        *generated = append(*generated, string(curString))
        return
    }

    if len(*generated) >= sz + 1 {
        return
    }

    for i := 0; i <= 1; i++ {
        curString = append(curString, rune(i + int('0')))
        DFS(pos + 1, sz, curString, generated)
        curString = curString[: len(curString) - 1]
    }
}
