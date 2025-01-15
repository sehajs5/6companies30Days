package main

func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)
	buildGraph(root, nil, graph)
	return bfs(graph, start)
}
func buildGraph(node *TreeNode, parent *TreeNode, graph map[int][]int) {
	if node == nil {
		return
	}
	if parent != nil {
		graph[node.Val] = append(graph[node.Val], parent.Val)
		graph[parent.Val] = append(graph[parent.Val], node.Val)
	}
	buildGraph(node.Left, node, graph)
	buildGraph(node.Right, node, graph)
}
func bfs(graph map[int][]int, start int) int {
	visited := make(map[int]bool)
	queue := []int{start}
	visited[start] = true
	minutes := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			popped := queue[0]
			queue = queue[1:]

			for _, nei := range graph[popped] {
				if visited[nei] == false {
					visited[nei] = true
					queue = append(queue, nei)
				}
			}
		}
		if len(queue) > 0 {
			minutes++
		}
	}
	return minutes
}
