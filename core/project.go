package core

// ProjectType .
type ProjectType int

const (
	eventProject = ProjectType(iota)
	normalProject
	persistProject
)

// Effect struct
type Effect struct {
	condition func(*Game, *Player, interface{}) bool
	effect    func(*Game, *Player, interface{})
}

// ProjectInfo interface
type ProjectInfo struct {
	ptype ProjectType
	code  string
	cost  int
	score exInt
	tag   map[Tag]int

	effect []Effect
}

// Project .
type Project struct {
	info *ProjectInfo

	game     *Game
	player   *Player
	enchants map[string]int
}

func (project *Project) getEnchant() map[string]int {
	return project.enchants
}

func (project *Project) test() bool {
	for _, i := range project.info.effect {
		if !i.condition(project.player.game, project.player, project) {
			return false
		}
	}
	return true
}

func (project *Project) run() {
	for _, i := range project.info.effect {
		i.effect(project.player.game, project.player, project)
	}
}
