func replaceNonCoprimes(nums []int) []int {
    result := []int{}

    for _, num := range nums {
        result = append(result, num)

        for len(result) >= 2 && !coprime(result[len(result) - 1], result[len(result) - 2]) {
            m := lcm(result[len(result) - 1], result[len(result) - 2])
            result = result[: len(result) - 2]
            result = append(result, m)
        }
    }

    return result
}

func coprime(a, b int) bool {
    return gcd(a, b) == 1
}

func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func lcm(a, b int) int {
    return (a / gcd(a, b)) * b
}
