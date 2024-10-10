func closeStrings(word1 string, word2 string) bool {
    freqMap1 := buildFreqMap(word1)
    freqMap2 := buildFreqMap(word2)

    if differentKeys(freqMap1, freqMap2) {
        return false
    }

    freqArray1 := buildFreqArray(freqMap1)
    freqArray2 := buildFreqArray(freqMap2)

    sort.Ints(freqArray1)
    sort.Ints(freqArray2)

    return equal(freqArray1, freqArray2)

}

func buildFreqMap(word string) map[rune]int {
    result := map[rune]int{}
    for _, c := range word {
        result[c]++
    }
    return result
}

func buildFreqArray(freqMap map[rune]int) []int {
    result := []int{}

    for _, val := range freqMap {
        result = append(result, val)
    }
    return result
}

func equal(arr1 []int, arr2 []int) bool {
    if len(arr1) != len(arr2) {
        return false
    }

    for i := 0; i < len(arr1); i++ {
        if arr1[i] != arr2[i] {
            return false
        }
    }
    return true
}

func differentKeys(map1 map[rune]int, map2 map[rune]int) bool {
    if len(map1) != len(map2) {
        return true
    }

    for key, _ := range map1 {
        if _, ok := map2[key]; !ok {
            return true
        }
    }
    return false
}
