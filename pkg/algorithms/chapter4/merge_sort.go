package chapter4

func merge[T any](lxs, rxs []T, lt func(l, r T) bool) []T {
	//TODO Make ToList in Golangz iterate backwards for efficiency
	//TODO Make a stack in Golangz with push pop and peek.
	var rs = FromArray(lxs) //This is really stack once I have it in Golangz.  Just needs improvements mentioned previously.
	var ls = FromArray(rxs) //This is really stack once I have it in Golangz
	var x, y T
	r := make([]T, len(lxs)+len(rxs))
	var lerr error
	var rerr error
	for i := range r {
		lerr, y = Peek(ls)
		rerr, x = Peek(rs)
		if lerr == nil && rerr == nil {
			if lt(x, y) {
				_, x, rs = Pop(rs)
				r[i] = x
			} else {
				_, y, ls = Pop(ls)
				r[i] = y
			}
		} else if lerr == nil {
			_, y, ls = Pop(ls)
			r[i] = y
		} else if rerr == nil {
			_, x, rs = Pop(rs)
			r[i] = x
		}
	}
	return r
}

func MergeSort[T any](xs []T, lt func(l, r T) bool) []T {
	//TODO move your stack to Golangz

	a := len(xs)
	if a == 0 {
		return xs
	} else if a == 1 {
		return xs
	} else if a == 2 {
		if lt(xs[0], xs[1]) {
			return xs
		} else {
			return []T{xs[1], xs[0]}
		}
	} else {
		i := a / 2
		left := xs[0:i]
		right := xs[i:]
		l := MergeSort(left, lt)
		r := MergeSort(right, lt)
		merged := merge(l, r, lt)
		return merged
	}
}
