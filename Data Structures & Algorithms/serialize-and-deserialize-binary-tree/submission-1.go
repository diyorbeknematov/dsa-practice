/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}

    datas := make([]string, 0)
	queue := []*TreeNode{ root }

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			datas = append(datas, "null")
			continue
		}

		str := strconv.Itoa(node.Val)
		datas = append(datas, str)

		// left child
		queue = append(queue, node.Left)
		
		// right child 
		queue = append(queue, node.Right)
	}

	for len(datas) > 0 && datas[len(datas)-1] == "null" {
    	datas = datas[:len(datas)-1]
	}

	return strings.Join(datas, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	
	if len(nodes) == 1 && nodes[0] == "null" {
		return nil
	}

	val, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{ Val: val }
	queue := []*TreeNode{ root }
	i := 1

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != "null" {
			val, _ := strconv.Atoi(nodes[i])

			node := &TreeNode{ Val: val }
			parent.Left = node
			queue = append(queue, node)
		}

		if i + 1 < len(nodes) && nodes[i+1] != "null" {
			val, _ := strconv.Atoi(nodes[i+1])

			node := &TreeNode{ Val: val }
			parent.Right = node
			queue = append(queue, node)
		}

		i += 2
	}

	return root 
}
