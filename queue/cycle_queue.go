package queue

type CycleQueue struct {
	front, tail int
	inner []interface{}
	size int
}
