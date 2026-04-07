const INF = int(1e9)

func ladderLength(beginWord string, endWord string, wordList []string) int {
    if !inList(endWord, wordList) {
        return 0
    }

    wordList = append(wordList, beginWord)
    wordList = append(wordList, endWord)
    adjList := buildGraph(wordList)

    return BFS(adjList, len(wordList), len(wordList) - 2, len(wordList) - 1)
}

func inList(word string, wordList []string) bool {
    for _, candidate := range wordList {
        if word == candidate {
            return true
        }
    }
    return false
}

func buildGraph(wordList []string) map[int][]int {
    adjList := map[int][]int{}

    for i := 0; i < len(wordList); i++ {
        for j := i + 1; j < len(wordList); j++ {
            if differBySingleLetter(wordList[i], wordList[j]) {
                adjList[i] = append(adjList[i], j)
                adjList[j] = append(adjList[j], i)
            }
        }
    }

    return adjList
}

func differBySingleLetter(a, b string) bool {
    diffs := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            diffs++
        }
    }
    return diffs == 1
}

func BFS(adjList map[int][]int, n, source, target int) int {
    dist := make([]int, n)

    for i := 0; i < n; i++ {
        dist[i] = INF
    }
    dist[source] = 0

    queue := []int{source}
    for len(queue) > 0 {
        cur := queue[0]
        queue = queue[1:]

        for _, neighbor := range adjList[cur] {
            if dist[neighbor] > 1 + dist[cur] {
                dist[neighbor] = 1 + dist[cur]
                queue = append(queue, neighbor)
            }
        }
    }

    if dist[target] == INF {
        return 0
    }
    return dist[target] + 1
}
