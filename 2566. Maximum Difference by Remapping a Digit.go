func minMaxDifference(num int) int {
    mini, maxi := num, num
    for startDigit := 0; startDigit <= 9; startDigit++ {
        for endDigit := 0; endDigit <= 9; endDigit++ {
            if startDigit != endDigit {
                nextNum := remap(num, startDigit, endDigit)

                mini = min(mini, nextNum)
                maxi = max(maxi, nextNum)
            }
        }
    }

    return maxi - mini
}

func remap(num, startDigit, endDigit int) int {
    str := []byte(fmt.Sprintf("%d", num))

    for i := 0; i < len(str); i++ {
        if int(str[i] - '0') == startDigit {
            str[i] = byte(endDigit + '0')
        }
    }

    nextNum, _ := strconv.Atoi(string(str))

    return nextNum
}
