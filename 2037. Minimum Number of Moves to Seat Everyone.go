func minMovesToSeat(seats []int, students []int) int {
    slices.Sort(seats)
    slices.Sort(students)

    result := 0
    for i := 0; i < len(seats); i++ {
        result += abs(seats[i] - students[i])
    }
    return result
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
