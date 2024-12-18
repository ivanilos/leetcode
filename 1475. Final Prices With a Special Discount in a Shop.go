type Elem struct {
    val int
    idx int
}

type myHeap []Elem

func (h myHeap) Len() int { return len(h) }
func (h myHeap) Less(i, j int) bool { return h[i].val > h[j].val}
func (h myHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *myHeap) Push (x interface{}) {
    *h = append(*h, x.(Elem))
}

func (h *myHeap) Pop () interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[0 : n - 1]
    return x
}

func (h *myHeap) Top () (int, int) {
    return (*h)[0].val, (*h)[0].idx
}

func finalPrices(prices []int) []int {
    h := &myHeap{}

    result := make([]int, len(prices))
    for i := 0; i < len(prices); i++ {
        for h.Len() != 0 {
            val, idx := h.Top()

            if val >= prices[i] {
                result[idx] -= prices[i]
                heap.Pop(h)
            } else {
                break
            }
        }
        result[i] = prices[i]
        heap.Push(h, Elem{prices[i], i})
    }

    return result
}
