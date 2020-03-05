package paging

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	total := len(list)
	//每页显示三条
	ps := 3
	for pi := 1; pi < 1000; pi++ {
		head, end, next, err := Paging(pi, ps, total)
		info:=fmt.Sprintf("head = %d, end = %d, next = %v. ",head,end,next)
		fmt.Println(info)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Printf("第 %d 次取出的值\n", pi)
		r := list[head:end]
		for _, v := range r {
			fmt.Printf("%d ",v)
		}
		fmt.Println()
	}

}
func TestPagingErr(t *testing.T) {
	head, end, next, err := Paging(11, 10, 0)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(head,end,next)
}