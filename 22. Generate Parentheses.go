func generateParenthesis(n int) []string {
    result := []string{}

    cur := []byte{}
    DFS(n, n, 0, &cur, &result)

    return result
}

func DFS(open, close, balance int, cur *[]byte, result *[]string) {
    if open == 0 && close == 0 {
        if balance == 0 {
            *result = append(*result, string(*cur))
        }
        return
    }

    if open > 0 {
        *cur = append(*cur, '(')
        DFS(open - 1, close, balance + 1, cur, result)
        *cur = (*cur)[: len(*cur) - 1]
    }

    if close > 0 && balance > 0 {
        *cur = append(*cur, ')')
        DFS(open, close - 1, balance - 1, cur, result)
        *cur = (*cur)[: len(*cur) - 1]
    }
}
