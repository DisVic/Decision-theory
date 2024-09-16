package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Item struct {
	id       int
	p        int
	weight   int
	relation float64
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var err error
	var capacity int
	_, err = fmt.Fscan(in, &capacity)
	if err != nil {
		fmt.Fprint(out, "ошибка ввода значения важности или веса (целочисленный тип обязателен)")
	}
	items := make([]Item, 10)
	for i := 0; i < len(items); i++ {
		items[i].id = i + 1
		_, err = fmt.Fscan(in, &items[i].p, &items[i].weight)
		if err != nil {
			fmt.Fprint(out, "ошибка ввода значения важности или веса (целочисленный тип обязателен)")
		}
		items[i].relation = float64(items[i].p) / float64(items[i].weight)
	}
	sort.Slice(items, func(i, j int) bool { return items[i].relation < items[j].relation })

	sum := 0
	things := make([]int, 0)
	for _, v := range items {
		if (sum + v.weight) <= capacity {
			sum += v.weight
			things = append(things, v.id)
		}
	}
	sort.Slice(things, func(i, j int) bool { return things[i] < things[j] })
	for _, v := range things {
		fmt.Fprintf(out, "%d ", v)
	}
}
