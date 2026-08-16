package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/minimum-distance-between-bst-nodes/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func minDiffInBST(root *TreeNode) int {
	// inorder - даст в BST сортированную последовательность
	// 0 <= Node.val <= 10^5

	minDiff := math.MaxInt
	prev := math.MaxInt
	inOrderTree(root, &prev, &minDiff)
	return minDiff
}

func inOrderTree(root *TreeNode, prev *int, minDiff *int) {
	if root == nil {
		return
	}

	inOrderTree(root.Left, prev, minDiff)

	curDiff := abs(*prev - root.Val)
	if *minDiff > curDiff {
		*minDiff = curDiff
	}

	*prev = root.Val
	inOrderTree(root.Right, prev, minDiff)
}

func main() {
	{
		// root = [4,2,6,1,3]
		root := &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			Right: &TreeNode{Val: 6},
		}
		fmt.Println(minDiffInBST(root))
	}
	{
		// root = [1,0,48,null,null,12,49]
		root := &TreeNode{
			Val:  1,
			Left: &TreeNode{Val: 0},
			Right: &TreeNode{
				Val:   48,
				Left:  &TreeNode{Val: 12},
				Right: &TreeNode{Val: 49},
			},
		}
		fmt.Println(minDiffInBST(root))
	}
}
