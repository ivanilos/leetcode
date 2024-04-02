func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    mapa := map[byte]byte{}
    rev := map[byte]byte{}
    for i := 0; i < len(s); i++ {
        if val, ok := mapa[s[i]]; ok && val != t[i] {
            return false
        } else {
            mapa[s[i]] = t[i]
        }
        if val, ok := rev[t[i]]; ok && val != s[i] {
            return false
        } else {
            rev[t[i]] = s[i]
        }
    }
    return true
}
