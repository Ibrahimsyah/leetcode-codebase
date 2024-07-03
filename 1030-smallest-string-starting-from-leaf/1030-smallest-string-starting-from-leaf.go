/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
    smallest := ""
    var dfs func(n *TreeNode, curr string)
    dfs = func(n *TreeNode, curr string) {
        str := string(rune(n.Val + 97)) + curr
        if n.Left == nil && n.Right == nil {
            if smallest == "" || str < smallest{
                smallest = str
            }
            return
        }

        if n.Left != nil {
            dfs(n.Left, str)
        }

        if n.Right != nil {
            dfs(n.Right, str)
        }
    }

    dfs(root, "")
    return smallest
}