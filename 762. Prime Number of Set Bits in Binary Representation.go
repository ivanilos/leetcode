import "math/bits"

func countPrimeSetBits(left int, right int) int {
    result := 0
    for i := left; i <= right; i++ {
        onBits := bits.OnesCount(uint(i))

        if isPrime(onBits) {
            result++
        }
    }

    return result
}

func isPrime(a int) bool {
    if a == 1 {
        return false
    }
    for i := 2; i * i <= a; i++ {
        if a % i == 0 {
            return false
        }
    }
    return true
}
