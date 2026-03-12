const MAX_STR = int(1e6)

type Edge struct {
    u int
    v int
    str int
    must int
}

type DSU struct {
    n int
    s []int
    size []int
}

func newDSU(n int) *DSU {
    s := make([]int, n)
    size := make([]int, n)

    for i := 0; i < n; i++ {
        s[i] = i
        size[i] = 1
    }

    return &DSU {
        n,
        s,
        size,
    }
}

func (d *DSU) find(x int) int {
    if d.s[x] == x {
        return x
    }
    d.s[x] = d.find(d.s[x])
    return d.s[x]
}

func (d *DSU) unite(x, y int) bool {
    x = d.s[x]
    y = d.s[y]

    if x == y {
        return false
    }

    if d.size[x] < d.size[y] {
        x, y = y, x
    }

    d.s[y] = x
    if d.size[y] == d.size[x] {
        d.size[x]++
    }

    return true
}



func maxStability(n int, edges [][]int, k int) int {
    edgeList := getEdges(edges)

    slices.SortFunc(edgeList, func(a, b Edge) int {
        if a.must > b.must {
            return -1
        } else if b.must > a.must {
            return 1
        } else {
            return b.str - a.str
        }
    })

    return solve(n, edgeList, k)
}

func solve(n int, sortedEdgeList []Edge, k int) int {
    low := 0
    high := MAX_STR

    best := -1

    for low <= high {
        mid := (low + high) / 2

        if canBuild(n, sortedEdgeList, k, mid) {
            best = mid
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return best
}

func canBuild(n int, sortedEdgeList []Edge, k int, strNeeded int) bool {
    dsu := newDSU(n)

    got := 0
    for _, edge := range sortedEdgeList {
        u := edge.u
        v := edge.v

        if edge.must == 1 {
            if dsu.find(u) == dsu.find(v) {
                return false
            }
            
            if edge.str < strNeeded {
                return false
            }

            dsu.unite(u, v)
            got++

        } else {
            if dsu.find(u) != dsu.find(v) {
                if edge.str >= strNeeded {
                    dsu.unite(u, v)
                    got++
                } else if edge.str * 2 >= strNeeded && k > 0 {
                    dsu.unite(u, v)
                    got++
                    k--
                }
            }
        }
    }

    return got == n - 1
}

func getEdges(edges [][]int) []Edge {
    result := make([]Edge, len(edges))
    for i, edge := range edges {
        result[i] = Edge{edge[0], edge[1], edge[2], edge[3]}
    }

    return result
}
