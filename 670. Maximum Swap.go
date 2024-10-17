func maximumSwap(num int) int {
    val := strconv.Itoa(num)
    valSlice := []byte(val)

    for i := 0; i < len(valSlice); i++ {
        curMaxi := valSlice[i]
        curIdx := i
        for j := i + 1; j < len(valSlice); j++ {
            if valSlice[j] > curMaxi || valSlice[j] == curMaxi && curIdx != i {
                curMaxi = valSlice[j]
                curIdx = j
            }
        }

        if curIdx != i {
            valSlice[i], valSlice[curIdx] = valSlice[curIdx], valSlice[i]
            break
        }
    }

    result, _ := strconv.Atoi(string(valSlice))

    return result
}
