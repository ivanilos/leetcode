const INF = int(1e9)

func minimumCost(nums []int) int {
    mini := []int{INF, INF}

    for i := 1; i < len(nums); i++ {
        num := nums[i]
        idxToReplace := len(mini)

        for i := 0; i < len(mini); i++ {
            if num < mini[i] {
                idxToReplace = i
                break
            }
        }

        for i := len(mini) - 1; i > idxToReplace; i-- {
            mini[i] = mini[i - 1]
        }

        if idxToReplace < len(mini) {
            mini[idxToReplace] = num
        }

        fmt.Println(mini)
    }

    return nums[0] + mini[0] + mini[1]
}
