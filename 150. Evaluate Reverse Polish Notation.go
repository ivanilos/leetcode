func evalRPN(tokens []string) int {
    ops := map[string]bool{
        "+": true,
        "-": true,
        "*": true,
        "/": true,
    }

    st := []int{}
    for _, token := range tokens {
        if  _, ok := ops[token]; ok {
            val := opResult(st[len(st) - 2], st[len(st) - 1], token)
            st[len(st) - 2] = val
            st = st[0 : len(st) - 1]
        } else {
            val, _ := strconv.Atoi(token)
            st = append(st, val)
        }
    }
    return st[0]
}

func opResult(a, b int, op string) int {
    if op == "+" {
        return a + b
    } else if op == "-" {
        return a - b
    } else if op == "*" {
        return a * b
    } else if op == "/" {
        return a / b
    } else {
        panic("Not implemented op")
    }
}
