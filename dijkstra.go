package main

import "fmt"

func main() {
	graph := [][]int{
		{0, 2, 0, 1, 0},
		{2, 0, 4, 0, 0},
		{0, 4, 0, 0, 3},
		{1, 0, 0, 0, 5},
		{0, 0, 3, 5, 0},
	}

	n := len(graph)
	INF := 999999

	dist := make([]int, n)
	visited := make([]bool, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0

	for i := 0; i < n; i++ {
		u := -1
		for v := 0; v < n; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}

		visited[u] = true

		for v := 0; v < n; v++ {
			if graph[u][v] != 0 && !visited[v] && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	fmt.Println("Dijkstra from node 0:", dist)
}
