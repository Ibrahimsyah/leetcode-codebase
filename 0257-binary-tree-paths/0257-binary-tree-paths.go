/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    paths := []string{}
    var dfs func(n *TreeNode, path string)
    dfs = func(n *TreeNode, path string){
        temp := fmt.Sprintf("%d", n.Val)
        if path != "" {
            temp = fmt.Sprintf("%s->%d", path, n.Val)
        }

        if n.Left == nil && n.Right == nil {
            paths = append(paths, temp)
        }

        if n.Left != nil {
            dfs(n.Left, temp)
        }

        if n.Right != nil {
            dfs(n.Right, temp)
        }
    }

    dfs(root, "")
    return paths
}