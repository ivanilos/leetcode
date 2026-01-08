func maxDotProduct(nums1 []int, nums2 []int) int {
    n := len(nums1)
    m := len(nums2)

    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
    }

    dp[0][0] = max(0, nums1[0] * nums2[0])
    for i := 1; i < n; i++ {
        dp[i][0] = max(dp[i - 1][0], nums1[i] * nums2[0])
    }
    for j := 1; j < m; j++ {
        dp[0][j] = max(dp[0][j - 1], nums1[0] * nums2[j])
    }

    for i := 1; i < n; i++ {
        for j := 1; j < m; j++ {
            dp[i][j] = max(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1], nums1[i] * nums2[j] + dp[i - 1][j - 1])
        }
    }

    result := dp[0][0]
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Print(dp[i][j], " ")
            result = max(result, dp[i][j])
        }
        fmt.Println()
    }

    // hack - 0 is the value of the empty subsequence
    // if it is the max, a subsequence of size 1 has the second largest result
    if result == 0 {
        result = nums1[0] * nums2[0]
        for i := 0; i < n; i++ {
            for j := 0; j < m; j++ {
                result = max(result, nums1[i] * nums2[j])
            }
        }
    }

    return result
}
