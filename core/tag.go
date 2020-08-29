package core

type tag int

const (
	building = iota
	space
	wildcard
	tagsize
)

type tags []int

func buildTags() tags {
	return make([]int, tagsize)
}

func toTag(str string) tag {
	switch str {
	case "building":
		return building
	case "space":
		return space
	case "wildcard":
		return wildcard
	}

	return 0
}
