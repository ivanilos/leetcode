var OFFSET = int(1e4 + 1)
var MAX_VALUE_RANGE = int(2e4 + 1)

type BitTree struct {
    v []int
    offset int
    maxVal int
}

func newBitTree(maxValRange, offset int) *BitTree {
    return &BitTree{
        make([]int, maxValRange),
        offset,
        maxValRange,
    }
}

func (b *BitTree) add(idx int) {
    idx += b.offset
    for idx < b.maxVal {
        b.v[idx]++
        idx += idx & -idx
    }
}

func (b *BitTree) query(idx int) int {
    idx += b.offset

    result := 0
    for idx > 0 {
        result += b.v[idx]
        idx -= idx & -idx
    }

    return result
}

func countSmaller(nums []int) []int {
    bitTree := newBitTree(MAX_VALUE_RANGE, OFFSET)

    result := make([]int, len(nums))
    for i := len(nums) - 1; i >= 0; i-- {
        result[i] = bitTree.query(nums[i] - 1)
        bitTree.add(nums[i])
    }

    return result
}
