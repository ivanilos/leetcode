type Coupon struct {
    code string
    businessLine string
}

var validBusinessLines = map[string]bool {
    "electronics" : true,
    "grocery" : true,
    "pharmacy" : true,
    "restaurant" : true,
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
    coupons := []Coupon{}
    for i := 0; i < len(code); i++ {
        if isValidCode(code[i]) && isValidBusinessLine(businessLine[i]) && isActive[i] == true {
            coupons = append(coupons, Coupon{code[i], businessLine[i]})
        }
    }

    slices.SortFunc(coupons, func(a, b Coupon) int {
        if a.businessLine == b.businessLine {
            if a.code < b.code {
                return -1
            } else {
                return 1
            }
           
        }
        if a.businessLine < b.businessLine {
            return -1
        } else {
            return 1
        }
    })

    result := make([]string, len(coupons))
    for i := 0; i < len(coupons); i++ {
        result[i] = coupons[i].code
    }
    return result
}

func isValidCode(code string) bool {
    if len(code) == 0 {
        return false
    }
    for _, c := range code {
        if !unicode.IsLetter(c) && !unicode.IsDigit(c) && c != '_' {
            return false
        }
    }
    return true
}

func isValidBusinessLine(businessLine string) bool {
    _, ok := validBusinessLines[businessLine]

    return ok
}
