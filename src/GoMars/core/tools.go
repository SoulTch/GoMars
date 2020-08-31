package core

type exInt interface {
	val() int
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
