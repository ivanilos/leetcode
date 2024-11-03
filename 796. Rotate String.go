func rotateString(s string, goal string) bool {
    if len(s) != len(goal) {
        return false
    }

    for startIdx := 0; startIdx < len(s); startIdx++ {
        if isRotation(s, goal, startIdx) {
            return true
        }
    }
    return false
}

func isRotation(s, goal string, startIdx int) bool {
    for i := 0; i < len(s); i++ {
        if s[(i + startIdx) % len(s)] != goal[i] {
            return false
        }
    }
    return true
}
