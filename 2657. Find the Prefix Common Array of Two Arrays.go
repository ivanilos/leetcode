func findThePrefixCommonArray(A []int, B []int) []int {
    result := make([]int, len(A))
    gotA := make([]bool, len(A) + 1)
    gotB := make([]bool, len(B) + 1)

    counter := 0
    for i := 0; i < len(A); i++ {
        gotA[A[i]] = true

        if gotB[A[i]] {
            counter++
        }

        gotB[B[i]] = true

        
        if gotA[B[i]] {
            counter++
        }

        result[i] = counter
    }

    return result
}
