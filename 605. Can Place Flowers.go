func canPlaceFlowers(flowerbed []int, n int) bool {
    total := 0
    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 && notOccupied(flowerbed, i - 1, i + 1) {
            flowerbed[i] = 1
            total++
        }
    }

    return total >= n
}

func notOccupied(flowerbed []int, positions ...int) bool {
    for _, pos := range positions {
        if pos < 0 || pos >= len(flowerbed) {
            continue
        }
        if flowerbed[pos] == 1 {
            return false
        }
    }
    return true
}
