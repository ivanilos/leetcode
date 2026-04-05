func judgeCircle(moves string) bool {
    movesMap := map[rune][]int{
        'D': []int{1, 0},
        'U': []int{-1, 0},
        'L': []int{0, -1},
        'R': []int{0, 1},
    }

    x, y := 0, 0
    for _, move := range moves {
        x, y = x + movesMap[move][0], y + movesMap[move][1]
    }

    return x == 0 && y == 0
}
