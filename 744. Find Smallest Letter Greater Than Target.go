func nextGreatestLetter(letters []byte, target byte) byte {
    result := target

    for i := 0; i < len(letters); i++ {
        if letters[i] > target && (letters[i] < result || result == target) {
            result = letters[i]
        }
    }

    if result == target {
        return letters[0]
    }

    return result
}
