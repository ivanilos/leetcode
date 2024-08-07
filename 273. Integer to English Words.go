func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }

    groups := splitGroups(num)
    return solve(groups)
}

func solve(groups [][]string) string {
    result := []string{}

    tripleNames := []string{"", "Thousand", "Million", "Billion"}

    for i := len(groups) - 1; i >= 0; i-- {
        curName := getName(groups[i])
        if curName != "" {
            result = append(result, curName)

            if tripleNames[i] != "" {
                result = append(result, tripleNames[i])
            }
        }
    }

    return strings.Join(result, " ")
}

func getName(group []string) string {
    result := []string{}

    twoDigitsNames := map[string]string{
        "Zero": "Ten", 
        "One": "Eleven", 
        "Two": "Twelve", 
        "Three": "Thirteen", 
        "Four": "Fourteen",
        "Five": "Fifteen", 
        "Six": "Sixteen", 
        "Seven": "Seventeen", 
        "Eight": "Eighteen",
        "Nine": "Nineteen",}

    decimalNames := map[string]string{
        "Two": "Twenty", 
        "Three": "Thirty", 
        "Four": "Forty",
        "Five": "Fifty", 
        "Six": "Sixty", 
        "Seven": "Seventy", 
        "Eight": "Eighty",
        "Nine": "Ninety"}

    if (len(group) > 2 && group[2] != "Zero") {
        result = append(result, group[2])
        result = append(result, "Hundred")
    }
    if (len(group) > 1 && group[1] == "One") {
        result = append(result, twoDigitsNames[group[0]])
    } else {
        if (len(group) > 1 && group[1] != "Zero") {
            result = append(result, decimalNames[group[1]])
        }
        if (group[0] != "Zero") {
            result = append(result, group[0])
        }
    }

    return strings.Join(result, " ")
}

func splitGroups(num int) [][]string {
    names := []string{"Zero", "One", "Two", "Three", "Four", "Five",
                        "Six", "Seven", "Eight", "Nine"}

    result := [][]string{}
    for num > 0 {
        curTriple := []string{}
        for i := 0; i < 3; i++ {
            if num > 0 {
                val := num % 10
                num /= 10
                curTriple = append(curTriple, names[val])
            }
        }
        result = append(result, curTriple)
    }
    return result
}
