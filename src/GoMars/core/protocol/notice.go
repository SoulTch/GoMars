package protocol

type DrawCardNotice struct {
	Player int
	Amount int
}

type PayResourcesNotice struct {
	Player  int
	ResType []int
	Amount  []int
}
