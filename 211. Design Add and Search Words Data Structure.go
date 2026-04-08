type Node struct {
    children map[byte]int
    isWordEnd bool
}

type Trie struct {
    nodes []Node
}

type WordDictionary struct {
    trie Trie
}

func NewNode() Node {
    return Node {
        map[byte]int{},
        false,
    }
}


func Constructor() WordDictionary {
    return WordDictionary {
        Trie {
            []Node{NewNode()},
        },
    }
}


func (this *WordDictionary) AddWord(word string)  {
    curNodeIdx := 0

    for i := 0; i < len(word); i++ {
        c := word[i]
        if _, ok := this.trie.nodes[curNodeIdx].children[c]; !ok {
            this.trie.nodes[curNodeIdx].children[c] = len(this.trie.nodes)
            this.trie.nodes = append(this.trie.nodes, NewNode())
        }
        curNodeIdx = this.trie.nodes[curNodeIdx].children[c]
    }

    this.trie.nodes[curNodeIdx].isWordEnd = true
}


func (this *WordDictionary) Search(word string) bool {
    return DFS(this, 0, word, 0)
}

func DFS(wd *WordDictionary, curNodeIdx int, word string, idx int) bool {
    if idx >= len(word) {
        return wd.trie.nodes[curNodeIdx].isWordEnd
    }

    result := false
    if word[idx] == '.' {
        for _, val := range wd.trie.nodes[curNodeIdx].children {
            result = result || DFS(wd, val, word, idx + 1)
        }
    } else {
        if _, ok := wd.trie.nodes[curNodeIdx].children[word[idx]]; ok {
            nextNodeIdx := wd.trie.nodes[curNodeIdx].children[word[idx]]

            result = result || DFS(wd, nextNodeIdx, word, idx + 1)
        }
    }

    return result
}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
