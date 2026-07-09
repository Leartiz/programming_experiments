package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sampleTree() *TreeNode {
	return &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
}

// -----------------------------------------------------------------------

func preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := []int{root.Val}                     // обработать корень дерева
	out = append(out, preorder(root.Left)...)  // потом лево
	out = append(out, preorder(root.Right)...) // потом право

	return out
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := inorder(root.Left)

	out = append(out, root.Val)

	out = append(out, inorder(root.Right)...)

	return out
}

func postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := postorder(root.Left)
	out = append(out, postorder(root.Right)...)

	out = append(out, root.Val) // корень обработать последним

	return out
}

// -----------------------------------------------------------------------

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	out := make([]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]  // top
		queue = queue[1:] // pop

		out = append(out, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return out
}

// -----------------------------------------------------------------------

func main() {
	root := sampleTree()

	fmt.Println("preorder:  ", preorder(root))
	fmt.Println("inorder:   ", inorder(root))
	fmt.Println("postorder: ", postorder(root))
	fmt.Println("level:     ", levelOrder(root))
}
