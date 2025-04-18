func countAndSay(n int) string {
    result := make([][]byte, n + 1)

    result[1] = []byte{'1'}
    for i := 2; i <= n; i++ {
        last := byte('0')
        qt := 0
        for j := 0; j < len(result[i - 1]); j++ {
            if result[i - 1][j] == last {
                qt++
            } else {
                if qt > 0 {
                    result[i] = append(result[i], getQtAsBytes(qt)...)
                    result[i] = append(result[i], last)
                }
                last = result[i - 1][j]
                qt = 1
            }
        }
        result[i] = append(result[i], getQtAsBytes(qt)...)
        result[i] = append(result[i], last)
    }

    return string(result[n])
}

func getQtAsBytes(qt int) []byte {
    result := []byte{}
    for qt > 0 {
        d := qt % 10
        qt /= 10

        dByte := byte(d + '0')
        result = append(result, dByte)
    }

    for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
        result[i], result[j] = result[j], result[i]
    }

    return result
}
