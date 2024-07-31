func minHeightShelves(books [][]int, shelfWidth int) int {
    result := solve(books, shelfWidth)

    return result
}

func solve(books [][]int, shelfWidth int) int {
    dp := make([]int, len(books) + 1)

    dp[0] = 0

    for curBook := 1; curBook <= len(books); curBook++ {
        bookIdx := curBook - 1
        dp[curBook] = books[bookIdx][1] + dp[curBook - 1]

        curWidth := books[bookIdx][0]
        maxHeight := books[bookIdx][1]
        for i := bookIdx - 1; i >= 0; i-- {
            curWidth += books[i][0]
            maxHeight = max(maxHeight, books[i][1])

            if curWidth <= shelfWidth {
                dp[curBook] = min(dp[curBook], maxHeight + dp[i])
            } else {
                break
            }
        }
    }
    return dp[len(books)]
}
