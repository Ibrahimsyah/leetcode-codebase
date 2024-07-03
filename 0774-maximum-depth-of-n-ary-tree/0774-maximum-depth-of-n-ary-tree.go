/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    
    max := 0
    var dfs func(n *Node, d int)
    dfs = func (n *Node, d int) {
        if d > max {
            max = d
        }

        for _, c := range n.Children {
            dfs(c, d + 1)
        }
    }

    dfs(root, 1)
    return max
}