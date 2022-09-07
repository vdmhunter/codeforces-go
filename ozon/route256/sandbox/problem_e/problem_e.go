package main

import (
	"bufio"
	"fmt"
	"os"
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
		timeSpans = nil

		_, _ = fmt.Fscan(in, &n)
		_, _ = in.ReadString('\n')

		for j := 0; j < n; j++ {
			str, _ := in.ReadString('\n')
			str = strings.ReplaceAll(str, "\r", "")
			str = strings.ReplaceAll(str, "\n", "")
			subStrings := strings.SplitN(str, "-", 2)

			t1, err := Parse("2006-01-02 15:04:05", "2000-01-01 "+subStrings[0])
			if err != nil {
				_, _ = fmt.Fprintln(out, "NO")
				break
			}

			t2, err := Parse("2006-01-02 15:04:05", "2000-01-01 "+subStrings[1])
			if err != nil {
				_, _ = fmt.Fprintln(out, "NO")
				break
			}

			if t1.Unix() > t2.Unix() {
				_, _ = fmt.Fprintln(out, "NO")
				break
			}

			timeSpans = append(timeSpans, timeSpan{t1, t2})
		}

		for x := 0; x < len(timeSpans)-1; x++ {
			for y := x + 1; y < len(timeSpans); y++ {
				if isOverlapped(timeSpans[x].begin, timeSpans[x].end, timeSpans[y].begin, timeSpans[y].end) {
					_, _ = fmt.Fprintln(out, "NO")
					return
				}
			}
		}

		_, _ = fmt.Fprintln(out, "YES")
	}
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func isOverlapped(begin1 Time, end1 Time, begin2 Time, end2 Time) bool {
	return max(begin1.Unix(), begin2.Unix()) <= min(end1.Unix(), end2.Unix())
	//return begin1.Unix() < end2.Unix() && end1.Unix() > begin2.Unix()
}
