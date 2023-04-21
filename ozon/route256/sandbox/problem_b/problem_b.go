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
		var goodCount int
		var goodPrice int

		goods := make(map[int]int)

		fmt.Fscan(in, &goodCount)

		for j := 0; j < goodCount; j++ {
			fmt.Fscan(in, &goodPrice)
			goods[goodPrice]++
		}

		totalPrice := uint64(0)

		for price, count := range goods {
			goodCountForPay := count/3*2 + count%3
			totalPrice = totalPrice + uint64(price*goodCountForPay)
		}

		fmt.Fprintln(out, totalPrice)
	}
}
