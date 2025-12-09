const MOD = int(1e9 + 7)

func specialTriplets(nums []int) int {
    prefFreq := map[int]int{}
    sufFreq := map[int]int{}

    for _, num := range nums {
        sufFreq[num]++
    }

    result := 0
    for _, num := range nums {
        sufFreq[num]--

        result = (result + prefFreq[num * 2] * sufFreq[num * 2]) % MOD

        prefFreq[num]++
    }

    return result
}
