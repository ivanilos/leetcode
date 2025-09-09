const MOD = int(1e9 + 7)

func peopleAwareOfSecret(n int, delay int, forget int) int {
    result := 1

    forgetQuantity := map[int]int{}
    forgetQuantity[forget - 1] = 1
    for i := 2; i <= n; i++ {
        result = sub(result, forgetQuantity[0])
        addedPeople := 0

        for forgetIdx := 1; forgetIdx < forget; forgetIdx++ {
            curDelay := forgetIdx - (forget - delay)
            if curDelay <= 0 {
                addedPeople = add(addedPeople, forgetQuantity[forgetIdx])

                result = add(result, forgetQuantity[forgetIdx])
            }

            forgetQuantity[forgetIdx - 1] = forgetQuantity[forgetIdx]
        }

        forgetQuantity[forget - 1] = addedPeople
    }

    return result
}

func add(a, b int) int {
    return (a + b) % MOD
}

func sub(a, b int) int {
    return (a - b + MOD) % MOD
}
