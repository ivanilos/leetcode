func subarrayBitwiseORs(arr []int) int {
    nextOn := make([]int, 32)
    for i := 0; i < 32; i++ {
        nextOn[i] = len(arr)
    }

    orSet := map[int]bool{}

    for i := len(arr) - 1; i >= 0; i-- {

        nextPosWithOnBit := make([]int, 32)
        for j := 0; j < 32; j++ {
            if ((arr[i] >> j) & 1) == 1 {
                nextOn[j] = i
            }
            nextPosWithOnBit[j] = nextOn[j]
        }

        slices.Sort(nextPosWithOnBit)

        curOr := arr[i]
        orSet[curOr] = true
        for _, pos := range nextPosWithOnBit {
            if pos < len(arr) {
                curOr |= arr[pos]
            } else {
                break
            }
            orSet[curOr] = true
        }
    }

    return len(orSet)
}
