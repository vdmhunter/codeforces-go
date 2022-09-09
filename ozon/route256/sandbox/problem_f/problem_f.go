package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		_ = out.Flush()
	}(out)

	var n int
	var m int

	_, _ = fmt.Fscan(in, &n, &m)

	adj := make([][]int, n)

	for i := 0; i < m; i++ {
		var u int
		var v int

		_, _ = fmt.Fscan(in, &u, &v)

		addEdge(adj, u-1, v-1)
	}

	for i := 0; i < n; i++ {
		sf := bfs(adj, i)
		for j := 0; j < len(sf); j++ {
			_, _ = fmt.Fprint(out, sf[j])
			if j < len(sf)-1 {
				_, _ = fmt.Fprint(out, " ")
			}
		}
		_, _ = fmt.Fprintln(out)
	}
}

func addEdge(adj [][]int, u int, v int) {
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
}

func bfs(adj [][]int, s int) []int {
	result := make([]int, 0)

	v2 := make(map[int]int)

	lengths := make(map[int]int)
	visited := make(map[int]bool)

	q := make([]int, 0)

	visited[s] = true
	q = append(q, s)

	for len(q) != 0 {
		u := q[0]
		q = q[1:]

		for _, v := range adj[u] {
			if !visited[v] {
				lengths[v] = lengths[u] + 1
				if lengths[v] == 2 {
					v2[v]++
				} else {
					visited[v] = true
					q = append(q, v)
				}
			}
		}
	}

	if len(v2) == 0 {
		return []int{0}
	} else {
		var max = -1
		for p := range v2 {
			if max < v2[p] {
				max = v2[p]
			}
		}
		for p := range v2 {
			if v2[p] == max {
				result = append(result, p+1)
			}
		}
	}

	sort.Ints(result)

	return result
}
