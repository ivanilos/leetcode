func countCollisions(directions string) int {
    st := []byte{}

    result := 0
    for _, direction := range directions {
        if direction == 'R' {
            st = append(st, 'R')
        } else if direction == 'S' {
            for len(st) > 0 && st[len(st) - 1] == 'R' {
                result++
                st = st[: len(st) - 1]
            }
            st = append(st, 'S')
        } else {
            if len(st) > 0 && st[len(st) - 1] == 'R' {
                result += 2
                st = st[: len(st) - 1]

                for len(st) > 0 && st[len(st) - 1] == 'R' {
                    result++
                    st = st[: len(st) - 1]
                }
                st = append(st, 'S')
            } else if len(st) > 0 && st[len(st) - 1] == 'S' {
                result++
            }
        }
    }

    return result
}
