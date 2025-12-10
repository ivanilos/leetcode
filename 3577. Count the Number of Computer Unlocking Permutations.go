const MOD = int64(1e9 + 7)

func countPermutations(complexity []int) int {
    mini := complexity[0]

    for i := 1; i < len(complexity); i++  {
        if complexity[i] <= mini {
            return 0
        }
    }

    return factorial(int64(len(complexity)) - 1)
}

func factorial(n int64) int {
    result := int64(1)

    for i := int64(1); i <= n; i++ {
        result = (result * i) % MOD
    }

    return int(result)
}
