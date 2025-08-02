func minCost(basket1 []int, basket2 []int) int64 {
    freq1 := getFreq(basket1)
    freq2 := getFreq(basket2)

    freqSum, ok := getFreqSum(freq1, freq2)
    if !ok {
        return int64(-1)
    }

    out1, out2, elementsToChange := getQuantityToChangeAndElements(freq1, freq2, freqSum)

    if out1 != out2 {
        return int64(-1)
    }

    minElement := min(slices.Min(basket1), slices.Min(basket2))
    return getMinCost(minElement, elementsToChange)
}

func getFreq(basket []int) map[int]int {
    result := map[int]int{}

    for _, val := range basket {
        result[val]++
    }
    return result
}

func getFreqSum(freq1, freq2 map[int]int) (map[int]int, bool) {
    freqSum := map[int]int{}

    for key, val := range freq1 {
        freqSum[key] += val
    }

    for key, val := range freq2 {
        freqSum[key] += val
    }

    for _, val := range freqSum {
        if val % 2 != 0 {
            return freqSum, false
        }
    }

    return freqSum, true
}

func getQuantityToChangeAndElements(freq1, freq2, freqSum map[int]int) (out1, out2 int, elementsToChange []int) {
    for key, val := range freqSum {
        // maybe encapsulate in a function
        if freq1[key] > val / 2 {
            changeQuantity := freq1[key] - val / 2
            out1 += changeQuantity
            for i := 0; i < changeQuantity; i++ {
                elementsToChange = append(elementsToChange, key)
            }
        }
        if freq2[key] > val / 2 {
            changeQuantity := freq2[key] - val / 2
            out2 += changeQuantity
            for i := 0; i < changeQuantity; i++ {
                elementsToChange = append(elementsToChange, key)
            }
        }
    }

    return out1, out2, elementsToChange
}

func getMinCost(minElement int, elementsToChange []int) int64 {
    result := int64(0)

    slices.Sort(elementsToChange)

    for i := 0; i < len(elementsToChange) / 2; i++ {
        val := elementsToChange[i]
        if val > 2 * minElement {
            result += int64(2 * minElement)
        } else {
            result += int64(val)
        }
    }

    return result
}
