type node struct {
    Children []*node
}

func NewNode() *node {
    result := &node{
        make([]*node, 10),
    }
    for i := 0; i < 10; i++ {
        result.Children[i] = nil
    }
    return result
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
    root := NewNode()

    for _, val := range arr1 {
        valString := strconv.Itoa(val)
        insert(root, valString)
    }

    result := 0
    for _,val := range arr2 {
        valString := strconv.Itoa(val)
        result = max(result, query(root, valString))
    }

    return result
}

func insert(root *node, val string) {
    for _, c := range val {
        pos := int(c - '0')
        if root.Children[pos] == nil {
            root.Children[pos] = NewNode()
        }
        root = root.Children[pos]
    }
}

func query(root *node, val string) int {
    result := 0

    for _, c := range val {
        pos := int(c - '0')
        if root.Children[pos] != nil {
            root = root.Children[pos]
            result++
        } else {
            break
        }
    }
    return result
}
