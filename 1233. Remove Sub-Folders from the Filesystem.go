type node struct {
    end bool
    children map[string]*node
}

func NewNode() *node {
    return &node{
        false,
        map[string]*node{},
    }
}

func removeSubfolders(folder []string) []string {
    root := NewNode()
    for _, path := range folder {
        splitPath := split(path)
        insert(root, splitPath)
    }

    result := query(root)

    return result
}

func split(path string) []string {
    return strings.Split(path, "/")[1 : ]
}

func insert(root *node, splitPath []string) {
    curNode := root
    for _, val := range splitPath {
        if curNode.end == true {
            break
        }

        if _, ok := curNode.children[val]; !ok {
            curNode.children[val] = NewNode()
        }
        curNode = curNode.children[val]
    }
    curNode.end = true
}

func query(root *node) []string {
    result := []string{}

    DFS(root, []string{}, &result)

    return result
}

func DFS(curNode *node, curPath []string, result *[]string) {
    if curNode.end == true {
        *result = append(*result, strings.Join(curPath, ""))
        return
    }

    for key, _ := range curNode.children {
        nextPath := append(curPath, "/", key)
        DFS(curNode.children[key], nextPath, result)
    }
}
