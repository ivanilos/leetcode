type DSU struct {
    n int
    s []int
    rank []int
}

func NewDSU (n int) *DSU {
    s := make([]int, n + 1)
    rank := make([]int, n + 1)
    for i := 1; i <= n; i++ {
        s[i] = i
        rank[i] = 1
    }

    return &DSU {
        n,
        s,
        rank,
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
    x = d.find(x)
    y = d.find(y)

    if x == y {
        return false
    }

    if d.rank[x] < d.rank[y] {
        x, y = y, x
    }

    d.s[y] = d.s[x]
    if d.rank[y] == d.rank[x] {
        d.rank[x]++
    }
    
    return true
}

func findRedundantConnection(edges [][]int) []int {
    n := len(edges) // 

    dsu := NewDSU(n)
    result := []int{}
    for _, e := range edges {
        u := e[0]
        v := e[1]

        if (!dsu.unite(u, v)) {
            result = []int{u, v}
            break
        }
    }

    return result
}
