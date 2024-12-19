func maxChunksToSorted(arr []int) int {
    result := 0
    curChunkStart := 0

    curChunk := 0
    for i := 0; i < len(arr); i++ {
        curChunk |= 1 << arr[i]

        needForNewChunk := (1 << (i - curChunkStart + 1)) - 1
        if ((curChunk  >> curChunkStart) & needForNewChunk) == needForNewChunk {
            result++
            curChunkStart = i + 1
        }
    }

    return result
}
