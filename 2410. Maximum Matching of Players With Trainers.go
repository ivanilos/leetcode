func matchPlayersAndTrainers(players []int, trainers []int) int {
    slices.Sort(players)
    slices.Sort(trainers)

    result := 0
    nextPlayer := 0

    for _, trainer := range trainers {
        if nextPlayer < len(players) && players[nextPlayer] <= trainer {
            result++
            nextPlayer++
        }
    }

    return result
}
