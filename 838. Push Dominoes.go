const INF = int(1e9)

func pushDominoes(dominoes string) string {
    n := len(dominoes)

    timer_left := make([]int, n)
    timer_right := make([]int, n)
    for i := 0; i < n; i++ {
        timer_left[i] = INF
        timer_right[i] = INF
    }

    for i := 0; i < n; i++ {
        if dominoes[i] == 'L' {
            timer_left[i] = 0
            for j := i - 1; j >= 0; j-- {
                if dominoes[j] == 'L' || dominoes[j] == 'R' {
                    break
                }
                timer_left[j] = i - j
            }
        }
    }

    for i := n - 1; i >= 0; i-- {
        if dominoes[i] == 'R' {
            timer_right[i] = 0
            for j := i + 1; j < n; j++ {
                if dominoes[j] == 'L' || dominoes[j] == 'R' {
                    break
                }
                timer_right[j] = j - i
            }
        }
    }

    result := make([]rune, n)
    for i := 0; i < n; i++ {
        if timer_left[i] < timer_right[i] {
            result[i] = 'L'
        } else if timer_left[i] > timer_right[i] {
            result[i] = 'R'
        } else {
            result[i] = '.'
        }
    }

    return string(result)
}
