func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    adjList := map[string][]string{}
    indegree := map[string]int{}

    recipesMap := map[string]bool{}

    for i, recipe := range recipes {
        recipesMap[recipe] = true
        for j := 0; j < len(ingredients[i]); j++ {
            adjList[ingredients[i][j]] = append(adjList[ingredients[i][j]], recipe)
            indegree[recipe]++
        }
    }

    for _, supply := range supplies {
        indegree[supply] = 0
    }

    return getPossibleRecipes(adjList, indegree, recipesMap)
}

func getPossibleRecipes(adjList map[string][]string, indegree map[string]int, recipesMap map[string]bool) []string {
    result := []string{}

    st := []string{}
    for key, val := range indegree {
        if val == 0 {
            st = append(st, key)
        }
    }

    for len(st) > 0 {
        ingredient := st[len(st) - 1]
        st = st[: len(st) - 1]

        if _, ok := recipesMap[ingredient]; ok {
            result = append(result, ingredient)
        }

        for _, neighbour := range adjList[ingredient] {
            indegree[neighbour]--

            if indegree[neighbour] == 0 {
                st = append(st, neighbour)
            }
        }
    }

    return result
}
