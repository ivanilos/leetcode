type Node struct{
    isWordEnd bool
    next []int
}

func NewTrieNode() Node {
    node := Node{}
    node.isWordEnd = false
    node.next = make([]int, 26)
    for i := 0; i < 26; i++ {
        node.next[i] = -1
    }
    return node
}

func insert(trie *[]Node, word string) {
    curNode := 0

    for _, c := range(word) {
        idx := int(c) - int('a')
        if (*trie)[curNode].next[idx] == -1 {
            (*trie)[curNode].next[idx] = len(*trie)
            (*trie) = append((*trie), NewTrieNode())
        }
        curNode = (*trie)[curNode].next[idx]
    }
    (*trie)[curNode].isWordEnd = true
}

func search(trie *[]Node, word string) string {
    curNode := 0

    curPrefix := []string{}
    for _, c := range(word) {
        idx := int(c) - int('a')
        if (*trie)[curNode].next[idx] == -1 {
            return word
        }
        curPrefix = append(curPrefix, string(c))
        curNode = (*trie)[curNode].next[idx]

        if (*trie)[curNode].isWordEnd {
            return strings.Join(curPrefix, "")
        }
    }
    return word
}

func replaceWords(dictionary []string, sentence string) string {
    trie := []Node{NewTrieNode()}

    for _, word := range(dictionary) {
        insert(&trie, word)
    }

    result := []string{}
    for _, word := range(strings.Split(sentence, " ")) {
        toAdd := search(&trie, word)
        result = append(result, toAdd)
    }
    return strings.Join(result, " ")
}
