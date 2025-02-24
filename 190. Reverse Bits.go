func reverseBits(num uint32) uint32 {
    result := uint32(0)
    for i := 0; i < 32; i++ {
        result = result | (((num >> i) & 1) << (31 - i))
    }
    return result
}
