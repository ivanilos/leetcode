func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    result := []string{}

    if sign(numerator) * sign(denominator) < 0 {
        result = append(result, "-")
    }

    dividend := int64(abs(numerator))
    divisor := int64(abs(denominator))

    integerPart := dividend / divisor
    strIntegerPart := strconv.FormatInt(integerPart, 10)
    result = append(result, strIntegerPart)

    remainder := dividend % divisor
    if remainder == 0 {
        return strings.Join(result, "")
    }

    result = append(result, ".")

    used := map[int64]int{}
    for remainder != 0 {
        if _, ok := used[remainder]; ok {
            result = append(result[: used[remainder]], "(", strings.Join(result[used[remainder]:], ""))
            result = append(result, ")")
            break
        }
        used[remainder] = len(result)
        remainder *= 10

        strVal := strconv.FormatInt(remainder / divisor, 10)
        result = append(result, strVal)
        remainder %= divisor
    }

    return strings.Join(result, "")
}

func sign(x int) int {
    if x < 0 {
        return -1
    }
    return 1
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
