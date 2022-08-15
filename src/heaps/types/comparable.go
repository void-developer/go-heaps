package types

type Comparable interface {
	Equals(C Comparable) bool
	Compare(C Comparable) int
	IsNull() bool
}
