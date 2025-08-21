func numSubmat(mat [][]int) int {
    pref := make([][]int, len(mat))
    for i := 0; i < len(mat); i++ {
        pref[i] = make([]int, len(mat[i]))
    }

    for i := 0; i < len(mat); i++ {
        pref[i][0] = mat[i][0]

        for j := 1; j < len(mat[i]); j++ {
            if mat[i][j] == 0 {
                pref[i][j] = 0
            } else {
                pref[i][j] = 1 + pref[i][j - 1]
            }
        }
    }

    result := 0
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[i]); j++ {
            curMin := pref[i][j]
            for k := i; k >= 0 && curMin > 0; k-- {
                curMin = min(curMin, pref[k][j])

                result += curMin
            }
        }
    }

    return result
}
