func decodeString(s string) string {
    result := _decodeString(s)

    return string(result)
}

func _decodeString(s string) []byte {
    if len(s) == 0 {
        return []byte{}
    }

    times := 0

    bracketDepth := 0
    startInner := 0
    endInner := 0
    curStr := []byte{}
    for i := 0; i < len(s); i++ {
        if bracketDepth == 0 && !unicode.IsDigit(rune(s[i])) && s[i] != '[' {
            for i < len(s) && !unicode.IsDigit(rune(s[i])) {
                curStr = append(curStr, s[i])
                i++
            }
            result := append(curStr, _decodeString(s[i : len(s)])...)
            return result
        }
        if bracketDepth == 0 && unicode.IsDigit(rune(s[i])) {
            times = 10 * times + int(s[i] - '0')
        } else if s[i] == '[' {
            if bracketDepth == 0 {
                startInner = i + 1
            }
            bracketDepth++
        } else if s[i] == ']' {
            bracketDepth--
            if bracketDepth == 0 {
                endInner = i
                break
            }
        }
    }

    result := append(multiply(times, _decodeString(s[startInner : endInner])), _decodeString(s[endInner + 1 : len(s)])...)
    return result
}

func multiply(times int, str []byte) []byte {
    result := []byte{}

    for i := 0; i < times; i++ {
        result = append(result, str...)
    }

    return result
}
