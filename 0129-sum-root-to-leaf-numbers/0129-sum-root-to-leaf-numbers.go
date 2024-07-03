/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    result := 0
    var dfs func(n *TreeNode, sum int)
    dfs = func(n *TreeNode, sum int){
        s := sum*10 + n.Val
        if n.Left == nil && n.Right == nil {
            result += s
            return
        }

        if n.Left != nil {
            dfs(n.Left, s)
        }

        if n.Right != nil {
            dfs(n.Right, s)
        }
    }

    dfs(root, 0)
    return result 
}