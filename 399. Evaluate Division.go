func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    varsMap := map[string]bool{}

    equationToValue := map[[2]string]float64{}

    for i := 0; i < len(equations); i++ {
        a, b := equations[i][0], equations[i][1]

        varsMap[a] = true
        varsMap[b] = true

        equationToValue[[2]string{a, b}] = values[i]
        equationToValue[[2]string{b, a}] = 1.0 / values[i]
    }

    result := make([]float64, len(queries))
    for i, query := range queries {
        a, b := query[0], query[1]

        if _, ok := varsMap[a]; !ok {
            result[i] = -1.0
        } else if _, ok := varsMap[b]; !ok {
            result[i] = -1.0
        } else {
            used := map[[2]string]bool{}
            result[i] = DFS(a, b, equationToValue, used)
        }
    }

    return result
}

func DFS(a, b string, equationToValue map[[2]string]float64, used map[[2]string]bool) float64 {
    node := [2]string{a, b}
    nodeInv := [2]string{b, a}

    if _, ok := equationToValue[node]; ok {
        return equationToValue[node]
    }

    if _, ok := used[node]; ok {
        return -1e9
    }

    used[node] = true

    result := -1.0
    for key, val := range equationToValue {
        if _, ok := used[key]; ok {
            continue
        }

        if key[0] == b {
            curResult := DFS(a, key[1], equationToValue, used)
            if curResult >= 0 {
                result = (1 / val) * curResult
            }
        }
    }

    if result >= 0 {
        equationToValue[node] = result
        equationToValue[nodeInv] = 1.0 / result
    }

    return result
}
