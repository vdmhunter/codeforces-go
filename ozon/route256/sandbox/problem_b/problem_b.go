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
		var goodsCount int
		var goodPrice int

		goods := make(map[int]int)

		_, _ = fmt.Fscan(in, &goodsCount)

		for j := 0; j < goodsCount; j++ {
			_, _ = fmt.Fscan(in, &goodPrice)
			goods[goodPrice]++
		}

		totalPrice := uint64(0)

		for price, count := range goods {
			goodCountForPay := count/3*2 + count%3
			totalPrice = totalPrice + uint64(price*goodCountForPay)
		}

		_, _ = fmt.Fprintln(out, totalPrice)
	}
}
