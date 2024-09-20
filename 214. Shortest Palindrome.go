const MOD int = int(1e9) + 7
const BASE int = 31

func shortestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }
    prefixHash := getPrefixHash(s)
    suffixHash := getSuffixHash(s)
    inverses := getInverses(len(s))

    for i := len(s) - 1; i >= 0; i-- {
        if isPalindrome(prefixHash, suffixHash, inverses, 0, i) {
            return solve(s, i)
        }
    }
    panic("should not reach here")
}

func getPrefixHash(s string) []int {
    hash := make([]int, len(s))

    hash[0] = int(s[0] - 'a' + 1)
    curBase := 1
    for i := 1; i < len(s); i++ {
        curBase = (curBase * BASE) % MOD
        hash[i] = (int(s[i] - 'a' + 1) * curBase + hash[i - 1]) % MOD
    }
    return hash
}

func getSuffixHash(s string) []int {
    hash := make([]int, len(s))

    hash[len(s) - 1] = int(s[len(s) - 1] - 'a' + 1)
    curBase := 1
    for i := len(s) - 2; i >= 0; i-- {
        curBase = (curBase * BASE) % MOD
        hash[i] = (int(s[i] - 'a' + 1) * curBase + hash[i + 1]) % MOD
    }
    return hash
}

func getInverses(sz int) []int {
    inverses := make([]int, sz)

    inverses[0] = 1

    curBase := 1
    for i := 1; i < sz; i++ {
        curBase = (curBase * BASE) % MOD
        inverses[i] = getInverse(curBase, MOD)
    }

    return inverses
}

func getInverse(val, mod int) int {
    var x, y int

    gcd(val, mod, &x, &y)
    return (x + mod) % mod
}

func gcd(a, b int, x, y *int) {
    if b == 0 {
        *x = 1
        *y = 0
        return 
    }

    gcd(b, a % b, y, x)
    *y -= (a / b) * (*x)
}

func isPalindrome(prefixHash, suffixHash, inverses []int, start, end int) bool {
    pref := prefixHash[end]
    suf := suffixHash[0]

    if end < len(suffixHash) - 1 {
        suf =  (suf - suffixHash[end + 1] + MOD) % MOD
    }

    suf = (suf * inverses[len(suffixHash) - (end + 1)]) % MOD

    return pref == suf 
}

func solve(s string, end int) string {
    result := []byte{}

    for i := len(s) - 1; i > end; i-- {
        result = append(result, s[i])
    }
    return string(result[:]) + s
}
