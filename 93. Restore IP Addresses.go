const IP_PARTS = 4

const MIN_ADDRESS_VAL = 0
const MAX_ADDRESS_VAL = 255

func restoreIpAddresses(s string) []string {
    validIPs := restoreIpAddressesParts(s, IP_PARTS)

    result := make([]string, len(validIPs))
    for i := 0; i < len(validIPs); i++ {
        result[i] = strings.Join(validIPs[i], ".")
    }

    return result
}

func restoreIpAddressesParts(s string, partsLeft int) [][]string {
    if partsLeft == 0 {
        result := [][]string{}
        if len(s) == 0 {
            result = append(result, []string{})
        }
        return result
    }

    result := [][]string{}
    for i := len(s) - 1; i >= 0; i-- {
        if isValidSubAddress(s[i:]) {
            partialAddresses := restoreIpAddressesParts(s[: i], partsLeft - 1)

            for _, partialAddress := range partialAddresses {
                address := append(partialAddress, s[i:])
                result = append(result, address)
            }
        }
    }

    return result
}

func isValidSubAddress(s string) bool {
    if len(s) > 1 && s[0] == '0' {
        return false
    }

    num, _ := strconv.Atoi(s)
    if num < MIN_ADDRESS_VAL ||  MAX_ADDRESS_VAL < num {
        return false
    }

    return true
}
