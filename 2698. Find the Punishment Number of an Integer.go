func punishmentNumber(n int) int {
    result := 0

    for i := 1; i <= n; i++ {
        if shouldAdd(i) {
            result += i * i
        }
    }

    return result
}

func shouldAdd(num int) bool {
    str := strconv.Itoa(num * num)

    return addUp(num, str)
}

func addUp(num int, str string) bool {
    if len(str) == 0 {
        return num == 0
    }
    if num < 0 {
        return false
    }

    for i := 0; i < len(str); i++ {
        partial, _ := strconv.Atoi(str[0 : i + 1])

        if addUp(num - partial, str[i + 1 : len(str)]) {
            return true
        }
    }
    return false
}
