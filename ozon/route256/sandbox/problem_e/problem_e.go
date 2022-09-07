package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	. "time"
)

type timeSpan struct {
	begin Time
	end   Time
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	defer func(out *bufio.Writer) {
		_ = out.Flush()
	}(out)

	var timeSpans []timeSpan
	var t int
	_, _ = fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		var flag = false
		timeSpans = nil

		_, _ = fmt.Fscan(in, &n)
		_, _ = in.ReadString('\n')

		for j := 0; j < n; j++ {
			str, _ := in.ReadString('\n')
			if flag {
				continue
			}
			str = strings.ReplaceAll(str, "\r", "")
			str = strings.ReplaceAll(str, "\n", "")
			subStrings := strings.SplitN(str, "-", 2)

			t1, err := Parse("2006-01-02 15:04:05", "2000-01-01 "+subStrings[0])
			if err != nil {
				flag = true
				continue
			}

			t2, err := Parse("2006-01-02 15:04:05", "2000-01-01 "+subStrings[1])
			if err != nil {
				flag = true
				continue
			}

			if t1.Unix() > t2.Unix() {
				flag = true
				continue
			}

			timeSpans = append(timeSpans, timeSpan{t1, t2})
		}

		if flag || isOverlapped(timeSpans) {
			_, _ = fmt.Fprintln(out, "NO")
			continue
		}

		_, _ = fmt.Fprintln(out, "YES")
	}
}

func isOverlapped(timeSpans []timeSpan) bool {
	if timeSpans == nil || len(timeSpans) == 0 {
		return false
	}

	sort.Slice(timeSpans, func(i, j int) bool {
		return timeSpans[i].begin.Unix() < timeSpans[j].begin.Unix()
	})

	for i := 1; i < len(timeSpans); i++ {
		var i1 = timeSpans[i-1]
		var i2 = timeSpans[i]
		if i1.end.Unix() >= i2.begin.Unix() {
			return true
		}
	}

	return false
}
