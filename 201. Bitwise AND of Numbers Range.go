func rangeBitwiseAnd(left int, right int) int {
    biggerThanLeft := right - left

    base := 2
    result := 0
    for i := 0;; i++{
        mod := left % base

        fmt.Println(i, (left >> i) & 1)
        if mod + biggerThanLeft < base && ((left >> i) & 1) == 1 {
            result = result + (1 << i)
        }
        if base > right {
            break
        }
        base *= 2
    }

    return result
}
