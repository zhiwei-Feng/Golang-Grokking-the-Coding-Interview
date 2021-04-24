package Challenge1_Tree_Diameter

/*
Given a binary tree, find the length of its diameter.
The diameter of a tree is the number of nodes on the longest path between any two leaf nodes.
The diameter of a tree may or may not pass through the root.

Note: You can always assume that there are at least two leaf nodes in the given tree.

类似题，但是题意有小区别
https://leetcode-cn.com/problems/diameter-of-binary-tree
*/

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter = 0
	calculateHeight(root, &maxDiameter)
	return maxDiameter
}

func calculateHeight(node *TreeNode, maxDiameterPtr *int) int {
	if node == nil {
		return 0
	}

	leftTreeH := calculateHeight(node.Left, maxDiameterPtr)
	rightTreeH := calculateHeight(node.Right, maxDiameterPtr)

	// 对于没有左子树或者右子树的节点，不去计算其diameter，因为这种情况不存在左子树的叶子到右子树的叶子这样一条路径
	if leftTreeH != 0 && rightTreeH != 0 {
		diameter := leftTreeH + rightTreeH + 1
		if diameter > *maxDiameterPtr {
			*maxDiameterPtr = diameter
		}
	}

	return max(leftTreeH, rightTreeH) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
