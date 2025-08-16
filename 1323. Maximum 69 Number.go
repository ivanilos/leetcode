func maximum69Number (num int) int {
    str := []byte(strconv.Itoa(num))

    for i := 0; i < len(str); i++ {
        if str[i] == '6' {
            str[i] = '9'
            break
        }
    }

    result, _ := strconv.Atoi(string(str))
    return result
}
