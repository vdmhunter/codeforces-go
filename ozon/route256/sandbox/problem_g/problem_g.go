package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.geeksforgeeks.org/number-of-connected-components-of-a-graph-using-disjoint-set-union/
// https://dou.ua/lenta/articles/union-find/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		out.Flush()
	}(out)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
	}
}
