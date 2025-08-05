func numOfUnplacedFruits(fruits []int, baskets []int) int {
    result := 0

    for _, qt := range fruits {
        fruitPlaced := false
        for i, cap := range baskets {
            if qt <= cap {
                fruitPlaced = true
                baskets[i] = 0
                break
            }
        }

        if !fruitPlaced {
            result++
        }
    }

    return result
}
