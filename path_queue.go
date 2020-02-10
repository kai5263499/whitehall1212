package whitehall1212

import "github.com/kai5263499/whitehall1212/types"

type pathEntry struct {
	path     []types.Edge
	distance int
}

type pathQueue []pathEntry

func (q pathQueue) Len() int {
	return len(q)
}
func (q pathQueue) Less(i, j int) bool {
	return q[i].distance < q[j].distance
}
func (q pathQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *pathQueue) Push(x interface{}) {
	item := x.(pathEntry)
	*q = append(*q, item)
}

func (q *pathQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
