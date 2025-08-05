func kClosest(points [][]int, k int) [][]int {
    slices.SortFunc(points, func(p1 []int, p2 []int) int {
        d1 := p1[0] * p1[0] + p1[1] * p1[1]
        d2 := p2[0] * p2[0] + p2[1] * p2[1]

        return d1 - d2
    })

    return points[ : k]
}
