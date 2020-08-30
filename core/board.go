package core

type tileType int

const (
	outOfBoard = tileType(iota)
	empty
	reserved
	city
	greenary
)

type tileInfo struct {
	ttype tileType
	name  string

	placable     func(*tile) bool
	adjacentable func(*tile) bool
	place        []effect
}

type tile struct {
	info   *tileInfo
	x, y   int
	player *player
}

func (t *tile) getPlayer() *player {
	return t.player
}

func (t *tile) test() bool {
	return t.player.game.board.placable(t)
}

func (t *tile) run() {
	t.player.game.board.place(t)
}

type board struct {
	tile [][]*tile
}

var dx = []int{1, 0}
var dy = []int{0, 1}

func (b *board) placable(t *tile) bool {
	if !b.tile[t.x][t.y].info.placable(t) {
		return false
	}

	for i := 0; i < 6; i++ {
		x, y := t.x+dx[i], t.y+dy[i]
		if !b.tile[x][y].info.placable(t) {
			return false
		}
	}

	return true
}

func (b *board) place(t *tile) {
	runEffects(b.tile[t.x][t.y].info.place, t)
	b.tile[t.x][t.y] = t
}
