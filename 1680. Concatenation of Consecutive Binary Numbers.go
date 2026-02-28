const MOD int = int(1e9 + 7)

func concatenatedBinary(n int) int {
    if n == 0 {
        return 0
    }

    size := getSize(n)

    return (((concatenatedBinary(n - 1) * getPot2(size)) % MOD) + n) % MOD
}

func getPot2(size int) int {
    result := 1

    for size > 0 {
        result = (result * 2) % MOD
        size--
    }

    return result
}

func getSize(n int) int {
    size := 0

    for n > 0 {
        size++
        n /= 2
    }

    return size
}
