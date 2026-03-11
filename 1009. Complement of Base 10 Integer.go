func bitwiseComplement(n int) int {
    if n == 0 {
        return 1
    }
    
    allOnes := getAllOnes(n)

    return allOnes ^ n
}

func getAllOnes(n int) int {
    for i := 0; i < 8; i++ {
        n = n | (n >> i)
    }
    return n
}
