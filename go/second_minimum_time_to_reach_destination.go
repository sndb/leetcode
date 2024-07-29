package main

import "fmt"

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make([][]int, n+1)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	now := 0
	answer := -1
	queue := []int{1}
	visit := make([][]int, n+1)
	for {
		k := len(queue)
		for _, v := range queue {
			if v == n {
				if answer != -1 {
					return now
				}
				answer = now
			}
			for _, u := range adj[v] {
				n := len(visit[u])
				if n == 0 || n == 1 && visit[u][0] != now {
					queue = append(queue, u)
					visit[u] = append(visit[u], now)
				}
			}
		}
		queue = queue[k:]
		if now/change%2 == 1 {
			now += change - (now % change)
		}
		now += time
	}
}

func main() {
	fmt.Println(secondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5)) // 13
	fmt.Println(secondMinimum(2, [][]int{{1, 2}}, 3, 2))                                 // 11
}
