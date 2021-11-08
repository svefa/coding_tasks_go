package main

func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))
	//v := make([]int, len(graph))
	stack := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		stack[i] = i
	}
	for len(stack) > 0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if color[i] == 0 {
			color[i] = 1
		}
		for _, j := range graph[i] {
			if color[j] != 0 && color[j] == color[i] {
				return false
			}
			if color[j] == 0 {
				color[j] = -color[i]
				stack = append(stack, j)
			}
		}
	}
	return true
}
