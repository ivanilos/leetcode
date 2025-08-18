const TARGET = 24.0
const EPS = 1e-6

func judgePoint24(cards []int) bool {
    results := getValues(cards, (1 << len(cards) - 1))

    fmt.Println(results)
    for _, val := range results {
        if abs(val - 24) < EPS {
            return true
        }
    }

    return false
}

func getValues(cards []int, mask int) []float64 {
    if isSingleValue(mask) {
        return []float64{float64(getValue(cards, mask))}
    }

    results := []float64{}
    for leftMask := (mask - 1) & mask; leftMask > 0; leftMask = (leftMask - 1) & mask {
        leftOperand := getValues(cards, leftMask)
        rightOperand := getValues(cards, mask ^ leftMask)

        for _, a := range leftOperand {
            for _, b := range rightOperand {
                results = append(results, a + b)
                results = append(results, a - b)
                results = append(results, a * b)
                if b != 0 {
                    results = append(results, a / b)
                } 
            }
        }
    }
    return results
}

func abs(x float64) float64 {
    if x < 0 {
        return -x
    }
    return x
}

func isSingleValue(mask int) bool {
    return mask == 1 || mask == 2 || mask == 4 || mask == 8
}

func getValue(cards []int, mask int) int {
    if mask == 1 {
        return cards[0]
    } else if mask == 2 {
        return cards[1]
    } else if mask == 4 {
        return cards[2]
    } else {
        return cards[3]
    }
}
