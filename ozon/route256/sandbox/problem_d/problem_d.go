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
		out.Flush()
	}(out)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		var daysCount int
		var taskId int
		var previousTaskId = 0

		tasks := make(map[int]bool)

		fmt.Fscan(in, &daysCount)

		flag := false

		for j := 0; j < daysCount; j++ {
			fmt.Fscan(in, &taskId)
			if previousTaskId == 0 {
				previousTaskId = taskId
			}
			if !tasks[taskId] {
				tasks[taskId] = true
				previousTaskId = taskId
			} else if previousTaskId != taskId {
				flag = true
			}
		}
		if flag {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
		}
	}
}
