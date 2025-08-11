const MOD = int(1e9 + 7)

func productQueries(n int, queries [][]int) []int {
    powers := getPowers(n)

    result := make([]int, len(queries))
    for i, query := range queries {
        result [i] = 1

        for j := query[0]; j <= query[1]; j++ {
            result[i] = (result[i] * powers[j]) % MOD
        }
    }

    return result
}

func getPowers(n int) []int {
    result := []int{}

    curPower := 1
    for n > 0 {
        if (n & 1) == 1 {
            result = append(result, curPower)
        }
        curPower *= 2
        n /= 2
    }

    return result
}
