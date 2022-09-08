package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://leetcode.com/discuss/interview-question/1408251/4-single-source-shortest-path-in-undirected-graph
// https://www.geeksforgeeks.org/shortest-path-unweighted-graph/
// https://stackoverflow.com/questions/2818852/is-there-a-queue-implementation
// https://stackoverflow.com/questions/3601180/calculate-distance-between-2-nodes-in-a-graph

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

	_, _ = fmt.Fprintln(out, "!")
}

func addEdge(adj [][]int, u int, v int) {
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
}
