package main

type LRDirection int

const (
	None LRDirection = iota
	Left
	Right
)

type Ship struct {
	Sprite    [8]byte
	Direction LRDirection
	PosX      int
	PosY      int
	Width     int
	Height    int
}

var shipNormalSprite = [8]byte{
	0b11100111,
	0b11000011,
	0b10000001,
	0b00100100,
	0b01100110,
	0b11100111,
	0b11000011,
	0b10011001,
}
var shipTurningRightSprite = [8]byte{
	0b11001111,
	0b10000111,
	0b00000011,
	0b00001101,
	0b01001111,
	0b10001111,
	0b10000011,
	0b01011011,
}

func NewShip(boardwidth, boardheight int) Ship {
	s := Ship{}
	s.Sprite = shipNormalSprite
	s.Direction = None
	s.Width = 8
	s.Height = 8
	s.PosX = boardwidth/2 - s.Width/2
	s.PosY = boardheight - 20
	return s
}
func (ship *Ship) moveLeft(boardwidth, boardheight int, land *Land) {
	var violation = land.landViolationLeft(ship.PosX - 1)
	if ship.PosX-1 > 0 && !violation {
		ship.PosX--
		ship.Sprite = shipTurningRightSprite
		ship.Direction = Left
	}
}
func (ship *Ship) moveRight(boardwidth, boardheight int, land *Land) {
	var violation = land.landViolationRight(ship.PosX + ship.Width)
	if ship.PosX+1 < boardwidth-ship.Width && !violation {
		ship.PosX++
		ship.Sprite = shipTurningRightSprite
		ship.Direction = Right
	}
}
func (ship *Ship) moveUp(boardwidth, boardheight int, land *Land) {
	var violation = land.landViolationUp(ship.PosY)
	if ship.PosY-1 > 120 && !violation {
		ship.PosY--
	}
}
func (ship *Ship) moveDown(boardwidth, boardheight int) {
	if ship.PosY+1 < boardheight-ship.Height {
		ship.PosY++
	}
}
func (ship *Ship) getRenderFlags() uint {
	if ship.Direction == Left {
		return BLIT_FLIP_X
	}
	return 0
}

func (ship *Ship) setupFrame() {
	ship.Direction = None
	ship.Sprite = shipNormalSprite
}
func (ship *Ship) render() {
	*DRAW_COLORS = 2
	blit(&ship.Sprite[0], ship.PosX, ship.PosY, uint(ship.Width), uint(ship.Height), BLIT_1BPP|ship.getRenderFlags())
}
