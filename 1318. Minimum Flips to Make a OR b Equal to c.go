func minFlips(a int, b int, c int) int {
    result := 0
    for i := 0; i < 32; i++ {
        bitA := (a >> i) & 1
        bitB := (b >> i) & 1
        bitC := (c >> i) & 1

        result += flipsToMakeEqual(bitA, bitB, bitC)
    }
    return result
}

func flipsToMakeEqual(bitA, bitB, bitC int) int {
    if (bitC == 0) {
        return bitA + bitB
    } else {
        if (bitA == 0 && bitB == 0) {
            return 1
        }
    }
    return 0
}
