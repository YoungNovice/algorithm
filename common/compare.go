package common

type Comparable interface {
	CompareTo(c Comparable) int
}
