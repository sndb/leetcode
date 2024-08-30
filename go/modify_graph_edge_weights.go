package main

const inf = 1_000_000_000

func modifiedGraphEdges(n int, edges [][]int, from, to, target int) [][]int {
	d := shortestDistance(n, edges, from, to)
	if d < target {
		return nil
	}
	ok := d == target

	for _, e := range edges {
		if e[2] == -1 {
			if ok {
				e[2] = inf
			} else {
				e[2] = 1
				d := shortestDistance(n, edges, from, to)
				if d <= target {
					ok = true
					e[2] += target - d
				}
			}
		}
	}
	if ok {
		return edges
	}
	return nil
}

func shortestDistance(n int, edges [][]int, from, to int) int {
	adj := make([][]int, n)
	for u := range n {
		adj[u] = make([]int, n)
		for v := range n {
			adj[u][v] = inf
		}
	}
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		if w != -1 {
			adj[u][v] = w
			adj[v][u] = w
		}
	}

	dist := make([]int, n)
	for u := range n {
		dist[u] = inf
	}
	dist[from] = 0

	seen := make([]bool, n)
	for range n {
		u := -1
		for v := range n {
			if !seen[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}
		seen[u] = true

		for v := range n {
			dist[v] = min(dist[v], dist[u]+adj[u][v])
		}
	}
	return dist[to]
}
