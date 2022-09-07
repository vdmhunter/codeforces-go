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
		var daysCount int
		var taskId int
		var previousTaskId = 0

		tasks := make(map[int]bool)

		_, _ = fmt.Fscan(in, &daysCount)

		var flag = false

		for j := 0; j < daysCount; j++ {
			_, _ = fmt.Fscan(in, &taskId)
			if previousTaskId == 0 {
				previousTaskId = taskId
			}
			_, exist := tasks[taskId]
			if !exist {
				tasks[taskId] = true
				previousTaskId = taskId
				continue
			} else {
				if previousTaskId == taskId {
					continue
				} else {
					flag = true
					continue
				}
			}
		}
		if flag {
			_, _ = fmt.Fprintln(out, "NO")
		} else {
			_, _ = fmt.Fprintln(out, "YES")
		}
	}
}
