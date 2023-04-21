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
		out.Flush()
	}(out)

	var setCount int
	fmt.Fscan(in, &setCount)

	for i := 0; i < setCount; i++ {
		var developerCount int
		var developerSkill int

		skills := make(map[int]int)
		positions := make([]int, 0, len(skills))

		fmt.Fscan(in, &developerCount)

		for j := 1; j <= developerCount; j++ {
			fmt.Fscan(in, &developerSkill)
			skills[j] = developerSkill
			positions = append(positions, j)
		}

		for _, position := range positions {
			skill, exist := skills[position]
			if !exist {
				continue
			}
			delete(skills, position)
			closestPosition := getClosestIndexByAbsDiff(skills, skill)
			delete(skills, closestPosition)

			fmt.Fprintln(out, position, closestPosition)
		}

		if i != setCount-1 {
			fmt.Fprintln(out)
		}
	}
}

func getClosestIndexByAbsDiff(values map[int]int, value int) (index int) {
	type kv struct {
		Key   int
		Value int
	}

	var arr []kv

	for k, v := range values {
		arr = append(arr, kv{k, v})
	}

	sort.Slice(arr, func(i, j int) bool {
		if abs(arr[i].Value-value) != abs(arr[j].Value-value) {
			return abs(arr[i].Value-value) < abs(arr[j].Value-value)
		}
		return arr[i].Key < arr[j].Key
	})

	return arr[0].Key
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
