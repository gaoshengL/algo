package search

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉查找树时间复杂度O(log2n)~O(n)

*/
/*
前序遍历，访问顺序，根节点=>左节点=>右节点
*/
func preorderTraversal(root *TreeNode) []int {
	n := []int{}
	preorderTraver(root, &n)
	return n
}

func preorderTraver(node *TreeNode, n *[]int) {
	if node == nil {
		return
	}
	*n = append(*n, node.Val)
	preorderTraver(node.Left, n)
	preorderTraver(node.Right, n)

}

/*
中序遍历，访问顺序，左节点=>根节点=>右节点
*/
func inorderTraversal(root *TreeNode) []int {
	n := []int{}
	inorderTraver(root, &n)
	return n
}

func inorderTraver(node *TreeNode, n *[]int) {
	if node == nil {
		return
	}
	inorderTraver(node.Left, n)
	*n = append(*n, node.Val)
	inorderTraver(node.Right, n)
}

/*
后序遍历，访问顺序，左节点=>右节点=>根节点
*/
func postorderTraversal(root *TreeNode) []int {
	n := []int{}
	postorderTraver(root, &n)
	return n
}

func postorderTraver(node *TreeNode, n *[]int) {
	if node == nil {
		return
	}
	postorderTraver(node.Left, n)
	postorderTraver(node.Right, n)
	*n = append(*n, node.Val)
}
