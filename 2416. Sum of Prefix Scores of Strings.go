type Node struct {
    Val int
    Children []*Node
}

func newNode() *Node {
    result := &Node{
        0,
        make([]*Node, 26),
    }
    return result
}

func sumPrefixScores(words []string) []int {
    root := newNode()

    for _, word := range words {
        insert(root, word)
    }

    result := make([]int, len(words))
    for i := 0; i < len(words); i++ {
        result[i] = query(root, words[i])
    }
    return result
}

func insert(root *Node, word string) {
    for _, c := range word {
        pos := int(c - 'a')
        if root.Children[pos] == nil {
            root.Children[pos] = newNode()
        }
        root = root.Children[pos]
        root.Val++
    }
}

func query(root *Node, word string) int {
    result := 0
    for _, c := range word {
        pos := int(c - 'a')
        if root.Children[pos] == nil {
            break
        }
        root = root.Children[pos]
        result += root.Val
    }
    return result
}
