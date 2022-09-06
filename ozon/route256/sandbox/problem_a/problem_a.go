package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		_ = out.Flush()
	}(out)

	var testCount int
	_, _ = fmt.Fscan(in, &testCount)

	for i := 0; i < testCount; i++ {
		var n, m int
		_, _ = fmt.Fscan(in, &n, &m)
		_, _ = fmt.Fprintln(out, n+m)
	}
}
