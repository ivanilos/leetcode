type Robot struct {
    id int
    pos int
    health int
    dir rune
}

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    robots := []Robot{}

    dirs := []rune(directions)
    for i := 0; i < len(positions); i++ {
        robots = append(robots, Robot{i, positions[i], healths[i], dirs[i]})
    }

    sort.Slice(robots, func(a, b int) bool {
        return robots[a].pos < robots[b].pos
    })

    st := []Robot{}
    for i := 0; i < len(robots); i++ {
        for len(st) > 0 && robots[i].health > 0 {
            topRobot := st[len(st) - 1]
            if topRobot.dir == 'R' && robots[i].dir == 'L' {
                if topRobot.health < robots[i].health {
                    st = st[0 : len(st) - 1]
                    robots[i].health--
                } else if topRobot.health > robots[i].health {
                    st[len(st) - 1].health--
                    robots[i].health = 0
                    break
                } else {
                    st = st[0 : len(st) - 1]
                    robots[i].health = 0
                    break
                }
            } else {
                break
            }
        }
        if robots[i].health > 0 {
            st = append(st, robots[i])
        }
    }

    sort.Slice(st, func(a, b int) bool {
        return st[a].id < st[b].id
    })

    result := []int{}
    for i := 0; i < len(st); i++ {
        result = append(result, st[i].health)
    }

    return result
}
