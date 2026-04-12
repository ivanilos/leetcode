var keyboard = [][]byte {
    []byte{'A', 'B', 'C', 'D', 'E', 'F'},
    []byte{'G', 'H', 'I', 'J', 'K', 'L'},
    []byte{'M', 'N', 'O', 'P', 'Q', 'R'},
    []byte{'S', 'T', 'U', 'V', 'W', 'X'},
    []byte{'Y', 'Z', ' ', ' ', ' ', ' '},
}

const INF = int(1e9)

type Pos struct {
    x int
    y int
}

type Node struct {
    pos int
    finger1 Pos
    finger2 Pos
}

func minimumDistance(word string) int {
    keyboardRows := len(keyboard)
    keyboardCols := len(keyboard[0])

    keyToPosMap := buildKeyToPosMap()

    dp := map[Node]int{}

    result := INF
    for i := 0; i < keyboardRows; i++ {
        for j := 0; j < keyboardCols; j++ {
            for k := 0; k < keyboardRows; k++ {
                for p := 0; p < keyboardRows; p++ {
                    node := Node{
                        0,
                        Pos{i, j},
                        Pos{k, p},
                    }

                    result = min(result, solve(node, word, keyToPosMap, dp))
                }
            }
        }
    }

    return result
}

func buildKeyToPosMap() map[byte]Pos {
    result := map[byte]Pos{}

    for i := 0; i < len(keyboard); i++ {
        for j := 0; j < len(keyboard[i]); j++ {
            result[keyboard[i][j]] = Pos{i, j}
        }
    }

    return result
}

func solve(node Node, word string, keyToPosMap map[byte]Pos, dp map[Node]int) int {
    if node.pos >= len(word) {
        return 0
    }

    if _, ok := dp[node]; ok {
        return dp[node]
    }

    newPos := keyToPosMap[word[node.pos]]
    // move right finger
    newNode := Node{node.pos + 1, newPos, node.finger2}
    result := getCost(node.finger1, newPos) + solve(newNode, word, keyToPosMap, dp)

    // move left finger
    newNode = Node{node.pos + 1, node.finger1, newPos}
    result = min(result, getCost(node.finger2, newPos) + solve(newNode, word, keyToPosMap, dp))

    dp[node] = result

    return result
}

func getCost(p1 Pos, p2 Pos) int {
    return abs(p1.x - p2.x) + abs(p1.y - p2.y)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func isIn(f Pos) bool {
    return 0 <= f.x && f.x < len(keyboard) && 0 <= f.y && f.y < len(keyboard[f.x])
}
