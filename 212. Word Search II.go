var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

type _Node struct {
    children map[byte]int
    isWordEnd bool
}

type _Trie struct {
    nodes []_Node
}

func newNode() _Node {
    return _Node {
        map[byte]int{},
        false,
    }
}

func NewTrie() _Trie {
    return _Trie {
        []_Node {newNode()},
    }
}

func findWords(board [][]byte, words []string) []string {
    trie := NewTrie()

    for _, word := range words {
        trie.insert(word)
    }

    searchOutput := map[string]bool{}
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            search(i, j, board, trie, 0, []byte{}, searchOutput)
        }
    }

    result := []string{}
    for key, _ := range searchOutput {
        result = append(result, key)
    }
    return result
}

func search(x, y int, board [][]byte, trie _Trie, trieNodeIdx int, curWord []byte, searchOutput map[string]bool) {
    if trie.nodes[trieNodeIdx].isWordEnd {
        searchOutput[string(curWord)] = true
    }

    if !isIn(x, y, board) {
        return
    }

    c := board[x][y]
    if _, ok := trie.nodes[trieNodeIdx].children[c]; ok {
        curWord = append(curWord, c)
        board[x][y] = '.'

        nextTrieNodeIdx := trie.nodes[trieNodeIdx].children[c]

        for i := 0; i < 4; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            search(nx, ny, board, trie, nextTrieNodeIdx, curWord, searchOutput)
        }

        board[x][y] = c
        curWord = curWord[0 : len(curWord) - 1]
    }
}

func isIn(x, y int, board [][]byte) bool {
    return 0 <= x && x < len(board) && 0 <= y && y < len(board[x])
}

func (t *_Trie) insert(word string) {
    curNodeIdx := 0

    for i := 0; i < len(word); i++ {
        c := word[i]
        if _, ok := t.nodes[curNodeIdx].children[c]; !ok {
            t.nodes[curNodeIdx].children[c] = len(t.nodes)
            t.nodes = append(t.nodes, newNode())
        }

        curNodeIdx = t.nodes[curNodeIdx].children[c]
    }

    t.nodes[curNodeIdx].isWordEnd = true
}
