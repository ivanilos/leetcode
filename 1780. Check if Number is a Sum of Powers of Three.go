func checkPowersOfThree(n int) bool {
    if n == 1 {
        return true
    }
    if n % 3 == 0 {
        return checkPowersOfThree(n / 3)
    } else if (n - 1) % 3 == 0{
        return checkPowersOfThree((n - 1) / 3)
    } else {
        return false
    }
}
