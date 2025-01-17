func doesValidArrayExist(derived []int) bool {
    curXor := 0
    for _, val := range derived {
        curXor ^= val
    }

    return curXor == 0
}
