package core

type bpID int

const (
	sellProject = bpID(iota)
	bpsize
)

type basicProject struct {
	id     bpID
	player *player
	cost   int
}

type bpEffect struct {
	cost    int
	effects []effect
}

var bpData = [bpsize]bpEffect{
	bpEffect{0, nil}}

func buildBasicProjects(p *player) []basicProject {
	ret := make([]basicProject, bpsize, bpsize)

	for i := range ret {
		ret[i].id = bpID(i)
		ret[i].player = p
		ret[i].cost = bpData[i].cost
	}

	return ret
}

func (bp *basicProject) getPlayer() *player {
	return bp.player
}

func (bp *basicProject) test() bool {
	return testEffects(bpData[int(bp.id)].effects, bp)
}

func (bp *basicProject) run() {
	runEffects(bpData[int(bp.id)].effects, bp)
}
