func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandiesWithAKid := slices.Max(candies)

    result := make([]bool, len(candies))
    for i := 0; i < len(candies); i++ {
        if candies[i] + extraCandies >= maxCandiesWithAKid {
            result[i] = true
        }
    }

    return result
}
