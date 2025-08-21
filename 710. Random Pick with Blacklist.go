type Solution struct {
    n int
    orderedBlacklist []int
}

const INF = int(1e9 + 1)

func Constructor(n int, blacklist []int) Solution {
    slices.Sort(blacklist)
    blacklist = append(blacklist, INF)
    return Solution{
        n,
        blacklist,
    }
}

func (this *Solution) Pick() int {
    chosenIdx := rand.IntN(this.n - len(this.orderedBlacklist) + 1) // account for INF guard

    return getValidIdx(chosenIdx, this.orderedBlacklist)
}

func getValidIdx(chosenIdx int, blacklist []int) int {
    low := 0
    high := len(blacklist) - 1
    result := low

    for low <= high {
        mid := (low + high) / 2

        if blacklist[mid] - mid > chosenIdx {
            result = mid
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    return chosenIdx + result
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
