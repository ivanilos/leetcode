func closestPrimes(left int, right int) []int {
    primes := sieve(left, right)

    result := []int{-1, -1}

    for i := 1; i < len(primes); i++ {
        if result[0] == -1 || primes[i] - primes[i - 1] < result[1] - result[0] {
            result = []int{primes[i - 1], primes[i]}
        }
    }
    return result
}

func sieve(left, right int) []int {
    result := []int{}

    notPrime := make([]bool, right + 1)
    for i := 2; i <= right; i++ {
        if !notPrime[i] {
            for j := i + i; j <= right; j += i {
                notPrime[j] = true
            }

            if i >= left {
                result = append(result, i)
            }
        }
    }

    return result
}
