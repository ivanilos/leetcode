func numOfSubarrays(arr []int) int {
    const MOD = int(1e9 + 7)
    result := 0

    curParity := 0
    parityFreq := []int{1, 0}

    for _, val := range arr {
        curParity = (curParity + val) % 2

        result = (result + parityFreq[1 - curParity]) % MOD
        parityFreq[curParity]++
    }

    return result
}
