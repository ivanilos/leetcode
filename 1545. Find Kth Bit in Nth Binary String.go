func findKthBit(n int, k int) byte {
    sz := make([]int, n + 1)

    sz[1] = 1
    for i := 2; i <= n; i++ {
        sz[i] = 2 * sz[i - 1] + 1
    }

    return solve(n, k, sz, false)
}

func solve(n, k int, sz []int, isRev bool) byte {
    if n == 1 && isRev {
        return '1'
    }

    if n == 1 && !isRev {
        return '0'
    }

    leftSz := sz[n - 1]

    if leftSz >= k {
        return solve(n - 1, k, sz, isRev)
    } else if leftSz + 1 == k {
        if isRev {
            return '0'
        } else {
            return '1'
        }
    } else {
        kWithoutReverse := k - (sz[n - 1] + 1)
        kWithReverse := sz[n - 1] - kWithoutReverse + 1

        if isRev {
            return solve(n - 1, kWithReverse, sz, false)
        } else {
            return solve(n - 1, kWithReverse, sz, true)
        }
    }
}
