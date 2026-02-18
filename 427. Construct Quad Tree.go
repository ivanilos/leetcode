/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

func construct(grid [][]int) *Node {
    return build(0, 0, len(grid) - 1, len(grid[0]) - 1, grid)
}

func build(topRow, leftCol, bottomRow, rightCol int, grid [][]int) *Node {
    curNode := &Node{}

    for i := topRow; i <= bottomRow; i++ {
        for j := leftCol; j <= rightCol; j++ {
            if grid[i][j] != grid[topRow][leftCol] {
                // split

                midRow := (topRow + bottomRow) / 2
                midCol := (leftCol + rightCol) / 2

                curNode.TopLeft = build(topRow, leftCol, midRow, midCol, grid)
                curNode.TopRight = build(topRow, midCol + 1, midRow, rightCol, grid)
                
                curNode.BottomLeft = build(midRow + 1, leftCol, bottomRow, midCol, grid)
                curNode.BottomRight = build(midRow + 1, midCol + 1, bottomRow, rightCol, grid)

                return curNode
            }
        }
    }

    curNode.IsLeaf = true
    
    if grid[topRow][leftCol] == 0 {
        curNode.Val = false
    } else {
        curNode.Val = true
    }

    return curNode
}
