package main

import (
	"fmt"
	"github.com/lofidv/jsutils"
	"strconv"
)

func main() {
	fmt.Println(jsutils.Contains([]string{"a", "b", "c"}, "b"))
	s := []int{2, 4, 8, 11}
	jsutils.ReverseSlice(s)
	fmt.Println(s)
	ss := []string{"joe", "mike", "hello"}
	jsutils.ReverseSlice(ss)
	fmt.Println(ss)
	sss := []int{2, 4, 8, 11}
	ds := jsutils.Map(sss, func(i int) string {return strconv.Itoa(2*i)})
	fmt.Println(ds)

	ssss := []int{2, 4, 8, 11}
	evens := jsutils.Filter(ssss, func(i int) bool {return i % 2 == 0})
	fmt.Println(evens)

	product := jsutils.Reduce(s, 1, func(a, b int) int {return a*b})
	fmt.Println(product)
}
