package paging

import (
	"errors"
)

/**
	逻辑分页
 */
func Paging(pi, ps, total int) (head int, end int, next bool, err error) {
	if pi == 0 || ps == 0 {
		return -1, -1, false, errors.New("分页参数不能为0")
	}

	head = (pi - 1) * ps

	//判断当前传入参数是否合法
	if head > total-1 {
		return -1, -1, false, errors.New("参数不合法,页码超过了最大范围")
	}

	if head+ps <= total {
		end = head + ps
	} else {
		end = total
	}

	//判断下一页
	if end < total {
		next = true
	}
	return head, end, next, nil
}
