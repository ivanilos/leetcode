var digitToLetter = map[byte]string {
    '2': "abc",
    '3': "def",
    '4': "ghi",
    '5': "jkl",
    '6': "mno",
    '7': "pqrs",
    '8': "tuv",
    '9': "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }
    
    result := []string{}

    curCombination := []byte{}
    generateCombinations(0, digits, &curCombination, &result)
    return result
}

func generateCombinations(pos int, digits string, curCombination *[]byte, result *[]string) {
    if pos >= len(digits) {
        *result = append(*result, string(*curCombination))
        return
    }

    for _, val := range digitToLetter[digits[pos]] {
        c := byte(val)
        *curCombination = append(*curCombination, c)
        generateCombinations(pos + 1, digits, curCombination, result)
        *curCombination = (*curCombination)[: len(*curCombination) - 1]
    }
}
