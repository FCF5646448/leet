/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }

	Given binary tree [3,9,20,null,null,15,7],

	  3
	/   \
   9    20
		/ \
	   15  7
 */
 func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0 
    }
    a := maxDepth(root.Left) +1
    b := maxDepth(root.Right) +1
    if a > b{
        return a
    }
    return b
}