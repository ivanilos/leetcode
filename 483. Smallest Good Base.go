const INF = int64(1e18)

func smallestGoodBase(n string) string {
    N, _ := strconv.ParseInt(n, 10, 64)

    result := N - 1

    for b := int64(2); b * (1 + b + b * b) <= N - 1; b++ {
        candidate := int64(0)
        for len := 0;; len++ {
            if candidate > INF / b {
                break
            }

            candidate = candidate * b + 1

            if candidate >= N {
                break
            }
        }

        if candidate == N {
            result = b
        }
    }

    if result == N - 1 {
        result = getClosestBase(N)
    }

    return strconv.FormatInt(result, 10)
}

func getClosestBase(N int64) int64 {
    root := int64(math.Sqrt(float64(N)))

    for root * (root + 1) < N - 1 {
        root++
    }
    for root * (root + 1) > N - 1 {
        root--
    }
    
    if root > 1 && root * (root + 1) == N - 1 {
        return root
    } else {
        return N - 1
    }
}
