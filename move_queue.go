package whitehall1212

import "github.com/kai5263499/whitehall1212/types"

type queueEntry struct {
	vertex   types.Vertex
	distance int
}

type moveQueue []queueEntry

func (q moveQueue) Len() int {
	return len(q)
}
func (q moveQueue) Less(i, j int) bool {
	return q[i].distance < q[j].distance
}
func (q moveQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *moveQueue) Push(x interface{}) {
	item := x.(queueEntry)
	*q = append(*q, item)
}

func (q *moveQueue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[0 : n-1]
	return x
}
