func threeConsecutiveOdds(arr []int) bool {
    for i := 0; i < len(arr) - 2; i++ {
        if AllOdd(arr[i], arr[i + 1], arr[i + 2]) {
            return true
        }
    }
    return false
}

func AllOdd(vals ...int) bool {
    for i := 0; i < len(vals); i++ {
        if vals[i] % 2 == 0 {
            return false
        }
    }
    return true
}
