func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
    seto := make([]int, 26)
    for i := 0; i < 26; i++ {
        seto[i] = i
    }

    for i := 0; i < len(s1); i++ {
        id1 := int(s1[i] - 'a')
        id2 := int(s2[i] - 'a')

        unite(id1, id2, seto)
    }

    result := []rune{}
    for _, c := range baseStr {
        id := int(c - 'a')

        newC := rune('a' + find(id, seto))
        result = append(result, newC)
    }

    return string(result)
}

func find(id int, seto []int) int {
    if seto[id] == id {
        return id
    }

    seto[id] = find(seto[id], seto)
    return seto[id]
}

func unite(x, y int, seto []int) {
    x = find(x, seto)
    y = find(y, seto)

    if x > y {
        x , y = y, x
    }

    seto[y] = seto[x]
}
