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

	var setCount int
	_, _ = fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
	}
}
