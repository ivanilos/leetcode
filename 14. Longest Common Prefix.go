func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }

    for lcp := 0; lcp < len(strs[0]); lcp++ {
        for i := 1; i < len(strs); i++ {
            if len(strs[i]) < lcp + 1 || strs[i][lcp] != strs[0][lcp] {
                return strs[i][0 : lcp]
            }
        }
    }
    return strs[0]
}
