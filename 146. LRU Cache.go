type Node struct {
    prev *Node
    next *Node
    key int
    val int
}

type LRUCache struct {
    head *Node
    tail *Node
    nodesMap map[int]*Node
    capacity int
}

func Constructor(capacity int) LRUCache {
    headDummy := &Node{nil, nil, -1, -1}
    tailDummy := &Node{nil, nil, -1, -1}

    headDummy.next = tailDummy
    tailDummy.prev = headDummy

    return LRUCache {
        headDummy, 
        tailDummy,
        map[int]*Node{},
        capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    if _, ok := this.nodesMap[key]; ok {
        this.makeMostRecentAccess(this.nodesMap[key])
        return (this.nodesMap[key]).val
    } else {
        return -1
    }
}

func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.nodesMap[key]; ok {
        (this.nodesMap[key]).val = value
        this.makeMostRecentAccess(this.nodesMap[key])
    } else {
        this.addNode(key, value)
    }
}

func (this *LRUCache) addNode(key, val int) {
    nodeBefore := this.head
    nodeAfter := this.head.next

    newNode := &Node{nodeBefore, nodeAfter, key, val}
    nodeBefore.next = newNode
    nodeAfter.prev = newNode

    this.nodesMap[key] = newNode

    if this.capacity < len(this.nodesMap) {
        this.removeLast()
    }
}

func (this *LRUCache) removeLast() {
    toRemove := this.tail.prev

    delete(this.nodesMap, toRemove.key)

    toRemove.prev.next = toRemove.next
    toRemove.next.prev = toRemove.prev
}

func (this *LRUCache) makeMostRecentAccess(node *Node) {
    nodeBefore := this.head
    nodeAfter := this.head.next

    // already has the most priority
    if node.prev == nodeBefore {
        return
    }

    node.prev.next = node.next
    node.next.prev = node.prev

    node.prev = nodeBefore
    node.next = nodeAfter

    nodeBefore.next = node
    nodeAfter.prev = node

    return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
