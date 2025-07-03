func kthCharacter(k int) byte {
    sz := 1

    for sz < k {
        sz *= 2
    }

    changes := 0
    for k > 1 {
        if k > sz / 2 {
            changes++
            k -= sz / 2
        }
        sz /= 2
    }

    return byte((changes % 26) + 'a')
}
