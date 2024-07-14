func countOfAtoms(formula string) string {
    tokens := Tokenize(formula)

    for _, token := range(tokens) {
        fmt.Println(token)
    }

    result := solve(tokens)

    return result
}

func Tokenize(formula string) []string {
    tokens := []string{}

    curToken := []rune{}
    for _, ch := range(formula) {
        if unicode.IsUpper(ch) {
            if len(curToken) > 0 {
                tokens = append(tokens, string(curToken))
            }
            curToken = []rune{ch}
        } else if unicode.IsLower(ch) {
            curToken = append(curToken, ch)
        } else if unicode.IsDigit(ch) {
            if len(curToken) > 0 && !unicode.IsDigit(curToken[0]) {
                tokens = append(tokens, string(curToken))
                curToken = []rune{ch}
            } else  {
                curToken = append(curToken, ch)
            }
        } else if ch == '(' || ch == ')' {
            if len(curToken) > 0 {
                tokens = append(tokens, string(curToken))
            }
            tokens = append(tokens, string(ch))
            curToken = []rune{}
        }
    }
    if len(curToken) > 0 {
        tokens = append(tokens, string(curToken))
    }
    return tokens
}

func solve(tokens []string) string {
    mapStack := []map[string]int{}
    mapStack = append(mapStack, map[string]int{})

    for idx, token := range(tokens) {
        if unicode.IsLetter(rune(token[0])) {
            qt := 1
            if idx + 1 < len(tokens) && unicode.IsDigit(rune(tokens[idx + 1][0])) {
                qt, _ = strconv.Atoi(tokens[idx + 1])
            }
            mapStack[len(mapStack) - 1][token] += qt
        }
        if token[0] == '(' {
            mapStack = append(mapStack, map[string]int{})
        } else if token[0] == ')' {
            qt := 1
            if idx + 1 < len(tokens) && unicode.IsDigit(rune(tokens[idx + 1][0])) {
                qt, _ = strconv.Atoi(tokens[idx + 1])
            }
            for key, _ := range(mapStack[len(mapStack) - 1]) {
                mapStack[len(mapStack) - 1][key] *= qt
            }

            if (len(mapStack) > 1) {
                for key, val := range(mapStack[len(mapStack) - 1]) {
                    mapStack[len(mapStack) - 2][key] += val
                }
                mapStack = mapStack[0 : len(mapStack) - 1]
            }
        }
    }

    return convertMapToString(mapStack[0])
}

func convertMapToString(mapa map[string]int) string {
    result := []string{}

    for key, val := range(mapa) {
        if val > 1 {
            result = append(result, fmt.Sprintf("%s%d", key, val))
        } else {
            result = append(result, key)
        } 
    }

    sort.Slice(result, func(a, b int) bool {
        return result[a] < result[b]
    })

    return strings.Join(result[:], "")
}
