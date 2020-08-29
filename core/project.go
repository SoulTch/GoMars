package core

type projectType int

const (
	eventProject = projectType(iota)
	normalProject
	persistProject
	corporation
)

type effect struct {
	condition func(belongsToPlayer) bool
	effect    func(belongsToPlayer)
}

func testEffects(e []effect, bp belongsToPlayer) bool {
	for _, i := range e {
		if !i.condition(bp) {
			return false
		}
	}
	return true
}

func runEffects(e []effect, bp belongsToPlayer) {
	for _, i := range e {
		i.effect(bp)
	}
}

type projectInfo struct {
	ptype  projectType
	code   string
	cost   int
	score  exInt
	tag    map[tag]int
	reqTag map[tag]int

	effects []effect
}

type project struct {
	info *projectInfo

	id       int
	game     *Game
	player   *player
	enchants map[string]int
}

func (project *project) getEnchant() map[string]int {
	return project.enchants
}

func (project *project) getPlayer() *player {
	return project.player
}

func (project *project) test() bool {
	return testEffects(project.info.effects, project)
}

func (project *project) run() {
	runEffects(project.info.effects, project)
}

type actionInfo struct {
	derived *projectInfo
	effects []effect
}

type action struct {
	info *actionInfo

	id      int
	derived *project
}

func (a *action) test() bool {
	return testEffects(a.info.effects, a.derived)
}

func (a *action) run() {
	runEffects(a.info.effects, a.derived)
}
